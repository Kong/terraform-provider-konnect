const fs = require("fs");
const yaml = require("js-yaml");

/**
 * Add a prefix and a parameter to all paths in the OpenAPI schema.
 * @param {string} inputFile - Path to the input OpenAPI YAML file.
 * @param {string} prefix - Prefix to add to all paths.
 * @param {string} parameterName - Name of the parameter to add.
 * @param {string} parameterSchemaYaml - YAML string defining the parameter schema.
 */
async function addPrefixAndParameterToOpenApiPaths(inputFile, prefix, parameterName, parameterSchemaYaml) {
    try {
        const schema = yaml.load(fs.readFileSync(inputFile, "utf8"));

        if (!schema.paths) {
            throw new Error("No paths found in the OpenAPI schema");
        }

        const parameterRef = `#/components/parameters/${parameterName}`;

        // Add prefix and parameter to each path
        const modifiedPaths = {};
        for (const [path, methods] of Object.entries(schema.paths)) {
            const newPath = path === "/" ? `${prefix}` : `${prefix}${path}`;
            modifiedPaths[newPath] = { ...methods };

            // Add parameter to each method in the path
            for (const [method, operation] of Object.entries(modifiedPaths[newPath])) {
                if (typeof operation === "object") {
                    if (!operation.parameters) {
                        operation.parameters = [];
                    }

                    // Avoid duplicate references
                    const alreadyExists = operation.parameters.some(
                        (param) => param.$ref === parameterRef
                    );
                    if (!alreadyExists) {
                        operation.parameters.push({ $ref: parameterRef });
                    }
                }
            }
        }

        // Update schema paths
        schema.paths = modifiedPaths;

        // Ensure the parameter is in components/parameters
        if (!schema.components) {
            schema.components = {};
        }
        if (!schema.components.parameters) {
            schema.components.parameters = {};
        }

        // Parse parameter schema from YAML, including custom fields
        schema.components.parameters[parameterName] = yaml.load(parameterSchemaYaml);

        // Save modified schema
        return yaml.dump(schema);
    } catch (err) {
        throw new Error(`Error processing OpenAPI schema: ${err.message}`);
    }
}

// Export function for use as a module
module.exports = addPrefixAndParameterToOpenApiPaths;

// Main execution flow
(async function () {
    if (require.main === module) {
        const args = process.argv.slice(2);

        if (args.length < 5) {
            console.error(
                "Usage: node add-prefix-to-paths.js <inputFile> <outputFile> <prefix> <parameterName> <parameterSchemaYaml>"
            );
            console.error("Provide parameterSchemaYaml as a YAML-formatted string, enclosed in quotes.");
            process.exit(1);
        }

        try {
            const [inputFile, outputFile, prefix, parameterName, parameterSchemaYaml] = args;
            const output = await addPrefixAndParameterToOpenApiPaths(inputFile, prefix, parameterName, parameterSchemaYaml);
            console.log(`Schema modified and saved to ${outputFile}`)
            fs.writeFileSync(outputFile, output, "utf8");
        } catch (error) {
            console.error(error.message);
            process.exit(1);
        }
    }
})();
