const { describe, test } = require("node:test");
const path = require("path");
const { compareWithGoldenFile } = require("./helpers/golden");
const addXSpeakeasyEntityOperation = require("./add-x-speakeasy-entity-operation");

// Test file paths
const INPUT_FILE = path.resolve(__dirname, "fixtures", "add-x-speakeasy-entity-operation", "input.yaml");
const GOLDEN_FILE = path.resolve(__dirname, "fixtures", "add-x-speakeasy-entity-operation", "output.golden.yaml");

describe("addXSpeakeasyEntityOperation", () => {
    test("Test addXSpeakeasyEntityOperation", async () => {
        // Run the function and get the updated schema as a string
        const updatedSchema = await addXSpeakeasyEntityOperation(INPUT_FILE);

        // Compare the updated schema with the golden file
        compareWithGoldenFile(updatedSchema, GOLDEN_FILE);
    });
});
