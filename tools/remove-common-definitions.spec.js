const { describe, test, beforeEach } = require("node:test");
const assert = require("assert");
const fs = require("fs");
const yaml = require("js-yaml");
const path = require("path");
const { compareWithGoldenFile } = require("./helpers/golden");
const removeCommonDefinitionsAndUpdateRefs = require("./remove-common-definitions");

// Test file paths
const INPUT_FILE = path.resolve(__dirname, "fixtures", "remove-common-definitions", "input.yaml");
const GOLDEN_FILE = path.resolve(__dirname, "fixtures", "remove-common-definitions", "output.golden.yaml");
const COMMON_DEFINITIONS_FILE = path.resolve(__dirname, "fixtures", "remove-common-definitions", "errors.yaml");

describe("removeCommonDefinitionsAndUpdateRefs", () => {
    test("should remove common definitions and update refs", async (t) => {
        // Run the main script to process the OpenAPI YAML
        const actualOutput = await removeCommonDefinitionsAndUpdateRefs(
            COMMON_DEFINITIONS_FILE,
            INPUT_FILE,
        );

        // Compare with the golden file
        try {
            compareWithGoldenFile(actualOutput, GOLDEN_FILE);
        } catch (error) {
            assert.fail(error.message);
        }
    });

})
