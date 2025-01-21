const fs = require("fs");
const yaml = require("js-yaml");

/**
 * Remove a deeply nested field from an OpenAPI schema that is temporarily unsupported.
 * @param {string} inputSchemaYaml - The input OpenAPI schema in YAML format.
 * @param {string} fieldPath - The dot-separated path of the field to remove.
 * @returns {string} - The updated OpenAPI schema as a YAML string.
 */
function removeTemporarilyUnsupportedField(inputSchemaYaml, fieldPath) {
    const schema = yaml.load(inputSchemaYaml);
    const pathSegments = fieldPath.split(".");

    // Navigate to the parent object of the field
    let current = schema;
    for (let i = 0; i < pathSegments.length - 1; i++) {
        if (current[pathSegments[i]] === undefined) {
            return yaml.dump(schema); // Field path does not exist, no change needed
        }
        current = current[pathSegments[i]];
    }

    // Remove the target field
    const targetField = pathSegments[pathSegments.length - 1];
    if (current && current.hasOwnProperty(targetField)) {
        delete current[targetField];
    }

    return yaml.dump(schema);
}

// Main execution flow
(async function () {
    if (require.main === module) {
        const args = process.argv.slice(2);

        if (args.length < 2) {
            console.error("Usage: node remove-temporarily-unsupported-field.js <inputFile> <fieldPath>");
            process.exit(1);
        }

        const [inputFile, fieldPath] = args;

        try {
            const inputSchemaYaml = fs.readFileSync(inputFile, "utf8");
            const updatedSchema = removeTemporarilyUnsupportedField(inputSchemaYaml, fieldPath);
            console.log(updatedSchema); // Print the updated schema to stdout
        } catch (error) {
            console.error(`Error: ${error.message}`);
            process.exit(1);
        }
    }
})();

module.exports = removeTemporarilyUnsupportedField;
