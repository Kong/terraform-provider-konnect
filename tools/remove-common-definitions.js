const fs = require("fs");
const yaml = require("js-yaml");

/**
 * Recursively updates $refs in any nested structure provided by a path,
 * but only if the referenced component exists in the common definitions.
 * @param {Object} obj - The object in which to perform substitutions.
 * @param {string} refPath - The path of keys pointing to the nested $ref location, e.g. 'components.schemas.Error.properties.invalid_parameters.items'.
 * @param {Object} commonDefinitions - The object containing common definitions (schemas, responses, etc.).
 * @param {string} refPrefix - The prefix to prepend to the updated $refs, e.g. "../../../../common/definitions/errors.yaml#".
 */
function updateRefInPath(obj, refPath, commonDefinitions, refPrefix) {
    const pathParts = refPath.split('.'); // Convert path into a list of keys.
    let current = obj;

    // Traverse down to the location of the object
    for (const part of pathParts) {
        if (current && current[part]) {
            current = current[part];
        } else {
            return; // Path does not exist, so exit
        }
    }

    // Perform substitution for $refs in the current location (array or object with nested refs)
    const updateRefs = (node) => {
        if (!node || typeof node !== "object") return;

        for (const key of Object.keys(node)) {
            const value = node[key];
            if (typeof value === "object") {
                // Handle items array case, or recurse deeper into the object
                if (key === "items" && value["$ref"]) {
                    // If there's a $ref in items
                    if (value["$ref"].startsWith("#/components/")) {
                        const [, refCategory, refType, refName] = value["$ref"].split("/");
                        if (commonDefinitions[refCategory] && commonDefinitions[refCategory][refType] && commonDefinitions[refCategory][refType][refName]) {
                            node[key]["$ref"] = `${refPrefix}/components/${refType}/${refName}`;
                            console.log(`Updated $ref in array item: ${value["$ref"]} -> ${node[key]["$ref"]}`);
                        }
                    }
                } else {
                    updateRefs(value); // Recursively update refs
                }
            } else if (key === "$ref" && typeof value === "string" && value.startsWith("#/components/")) {
                // If there's a $ref field
                const [, refCategory, refType, refName] = value.split("/");
                if (refCategory && refType && refName) {
                    if (commonDefinitions[refCategory] && commonDefinitions[refCategory][refType] && commonDefinitions[refCategory][refType][refName]) {
                        node[key] = `${refPrefix}/components/${refType}/${refName}`;
                        console.log(`Updated $ref: ${value} -> ${node[key]}`);
                    }
                }
            }
        }
    };

    // Execute updateRefs at the target level.
    updateRefs(current);
}

/**
 * Removes common definitions and updates $refs to common definitions in an OpenAPI file.
 * @param {string} commonDefinitionsFile - Path to the common definitions YAML file (default: src/common/definitions/errors.yaml).
 * @param {string} targetFile - Path to the target OpenAPI YAML file to process.
 * @param {string} outputFile - Path to save the updated OpenAPI YAML file.
 * @param {string} refPrefix - Prefix to update $refs to common definitions (default: "../../../../common/definitions/errors.yaml#").
 */
async function removeCommonDefinitionsAndUpdateRefs(
    commonDefinitionsFile,
    targetFile,
    refPrefix = "../../../../common/definitions/errors.yaml#"
) {
    try {
        // Load common definitions
        const commonDefinitions = yaml.load(fs.readFileSync(commonDefinitionsFile, "utf8"));

        // Load target schema
        const targetSchema = yaml.load(fs.readFileSync(targetFile, "utf8"));

        // Components to process for removal
        const componentsToProcess = ["schemas", "responses", "examples", "parameters"];

        if (commonDefinitions.components) {
            if (!targetSchema.components) {
                console.log("No components to remove in target file.");
            } else {
                // Loop through components to process
                for (const componentKey of componentsToProcess) {
                    const commonItems = commonDefinitions.components[componentKey];
                    const targetItems = targetSchema.components[componentKey];

                    if (commonItems && targetItems) {
                        const commonKeys = Object.keys(commonItems);
                        for (const key of commonKeys) {
                            // Remove common definitions
                            if (targetItems[key]) {
                                delete targetItems[key];
                                console.log(`Removed common ${componentKey.slice(0, -1)}: ${key}`);
                            }
                        }
                    }
                }
            }
        }

        // Update $refs in `paths.{path}.{method}.responses` or other locations
        const paths = targetSchema.paths;
        if (paths) {
            for (const [pathKey, pathObj] of Object.entries(paths)) {
                for (const [methodKey, methodObj] of Object.entries(pathObj)) {
                    if (["get", "post", "put", "delete", "patch", "options", "head"].includes(methodKey)) {
                        // Update $refs in responses
                        updateRefInPath(methodObj, 'responses', commonDefinitions, refPrefix);
                        // Optionally update other nested locations (parameters, etc.)
                        updateRefInPath(methodObj, 'parameters', commonDefinitions, refPrefix);
                    }
                }
            }
        }

        // Update $refs throughout components.{component}
        const components = targetSchema.components;
        if (components) {
            updateRefInPath(components, 'schemas', commonDefinitions, refPrefix); // Apply to schemas component
            updateRefInPath(components, 'responses', commonDefinitions, refPrefix); // Apply to responses component
        }

        // Save the updated schema
        return yaml.dump(targetSchema, { noRefs: true });
    } catch (err) {
        console.error(`Error processing OpenAPI schema: ${err.message}`);
    }
}

// Main execution flow
(async function () {
    if (require.main === module) {
        const args = process.argv.slice(2);

        if (args.length < 2) {
            console.error(
                "Usage: node remove-common-definitions.js <targetFile> <outputFile> [commonDefinitionsFile] [refPrefix]"
            );
            console.error(
                "Default commonDefinitionsFile: src/common/definitions/errors.yaml"
            );
            console.error(
                'Default refPrefix: "../../../../common/definitions/errors.yaml#"'
            );
            process.exit(1);
        }

        const targetFile = args[0];
        const outputFile = args[1];
        const commonDefinitionsFile = args[2] || "src/common/definitions/errors.yaml";
        const refPrefix = args[3] || "../../../../common/definitions/errors.yaml#";

        try {
            const updatedSchema = await removeCommonDefinitionsAndUpdateRefs(
                commonDefinitionsFile,
                targetFile,
                refPrefix
            );

            fs.writeFileSync(outputFile, updatedSchema, "utf8");
            console.log(`Updated OpenAPI schema saved to ${outputFile}`);
        } catch (error) {
            console.error(error.message);
            process.exit(1);
        }
    }
})();

module.exports = removeCommonDefinitionsAndUpdateRefs;
