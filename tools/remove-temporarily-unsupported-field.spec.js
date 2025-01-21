const { describe, test } = require("node:test");
const fs = require("fs");
const path = require("path");
const yaml = require("js-yaml");
const assert = require("assert");
const { compareWithGoldenFile } = require("./helpers/golden");
const removeTemporarilyUnsupportedField = require("./remove-temporarily-unsupported-field");

// File paths
const INPUT_FILE = path.resolve(__dirname, "fixtures", "remove-temporarily-unsupported-field", "input.yaml");
const GOLDEN_FILE = path.resolve(__dirname, "fixtures", "remove-temporarily-unsupported-field", "output.golden.yaml");

// Field path to be removed
const FIELD_PATH = "components.schemas.MeshHealthCheckItem.properties.spec.properties.to.items.properties.default.properties.http.properties.expectedStatuses.items.format";

describe("removeTemporarilyUnsupportedField", () => {
    test("Removes the specified field from the schema", async () => {
        // Load the input schema
        const inputSchemaYaml = fs.readFileSync(INPUT_FILE, "utf8");

        // Run the function to remove the field
        const updatedSchemaYaml = removeTemporarilyUnsupportedField(inputSchemaYaml, FIELD_PATH);

        // Compare with the golden file
        compareWithGoldenFile(updatedSchemaYaml, GOLDEN_FILE);
    });
});
