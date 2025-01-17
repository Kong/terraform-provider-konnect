const fs = require("fs");
const yaml = require("js-yaml");

/**
 * Modifies OpenAPI paths to add `x-speakeasy-entity-operation` derived from `operationId`.
 * @param {string} inputFile - The path to the input OpenAPI YAML file.
 * @returns {string} The updated OpenAPI schema as a YAML string.
 */
async function addXSpeakeasyEntityOperation(inputFile) {
    const schema = yaml.load(fs.readFileSync(inputFile, "utf8"));

    if (!schema.paths) {
        throw new Error("No paths found in the OpenAPI schema.");
    }

    const methods = ["get", "post", "put", "delete", "patch", "options", "head"];

    for (const [pathKey, pathObj] of Object.entries(schema.paths)) {
        for (const [methodKey, methodObj] of Object.entries(pathObj)) {
            if (methods.includes(methodKey) && methodObj.operationId) {
                const operationId = methodObj.operationId;

                // Extract the resource name and HTTP method from the operationId
                const resourceName = operationId.replace(new RegExp(`^(${methods.join("|")})`), "");
                const httpMethod = operationId.match(new RegExp(`^(${methods.join("|")})`))?.[0];

                // Skip if the resourceName does not start with an uppercase letter
                if (!/^[A-Z]/.test(resourceName)) {
                    continue;
                }

                if (httpMethod && resourceName) {
                    const methodMapping = { put: "update" };
                    const entityOperation = `${resourceName}#${methodMapping[httpMethod] || httpMethod}`;
                    methodObj["x-speakeasy-entity-operation"] = entityOperation;
                }
            }
        }
    }

    return yaml.dump(schema, { noRefs: true });
}

module.exports = addXSpeakeasyEntityOperation;

// For standalone execution
if (require.main === module) {
    (async () => {
        const args = process.argv.slice(2);

        if (args.length < 1) {
            console.error("Usage: node add-x-speakeasy-entity-operation.js <inputFile>");
            process.exit(1);
        }

        const inputFile = args[0];
        try {
            const updatedSchema = await addXSpeakeasyEntityOperation(inputFile);
            console.log(updatedSchema); // Outputs the updated schema as YAML
        } catch (error) {
            console.error(`Error: ${error.message}`);
        }
    })();
}
