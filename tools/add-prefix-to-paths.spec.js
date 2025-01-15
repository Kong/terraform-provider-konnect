const { describe, test } = require("node:test");
const path = require("path");
const fs = require("fs");
const yaml = require("js-yaml");
const { compareWithGoldenFile } = require("./helpers/golden");
const addPrefixAndParameterToOpenApiPaths = require("./add-prefix-to-paths");

// Test file paths
const INPUT_FILE = path.resolve(__dirname, "fixtures", "add-prefix-to-paths", "input.yaml");
const GOLDEN_FILE = path.resolve(__dirname, "fixtures", "add-prefix-to-paths", "output.golden.yaml");
const PARAMETER_FILE = path.resolve(__dirname, "fixtures", "add-prefix-to-paths", "parameter.yaml");

// Test parameters
const PREFIX = "/v2";

describe("addPrefixAndParameterToOpenApiPaths", () => {
    test("Test addPrefixAndParameterToOpenApiPaths", async () => {
        // Load parameter from file
        const parameter = fs.readFileSync(PARAMETER_FILE, "utf8");

        // Run the function and get the updated schema as a string
        const updatedSchema = await addPrefixAndParameterToOpenApiPaths(INPUT_FILE, PREFIX, "testParameter", parameter);

        // Compare the updated schema with the golden file
        compareWithGoldenFile(updatedSchema, GOLDEN_FILE);
    });
});
