const fs = require("fs");
const path = require("path");

/**
 * A utility function for comparing an actual output against a golden file.
 * If the environment variable `UPDATE_GOLDEN_FILES=true` is set, the golden file is updated.
 * @param {string} actualOutput - The actual output to compare.
 * @param {string} goldenFilePath - The path to the golden file.
 */
function compareWithGoldenFile(actualOutput, goldenFilePath) {
    const resolvedGoldenFilePath = path.resolve(goldenFilePath);

    if (!fs.existsSync(resolvedGoldenFilePath)) {
        throw new Error(`Golden file not found: ${resolvedGoldenFilePath}`);
    }

    const expectedOutput = fs.readFileSync(resolvedGoldenFilePath, "utf8");

    if (actualOutput !== expectedOutput) {
        if (process.env.UPDATE_GOLDEN_FILES === "true") {
            fs.writeFileSync(resolvedGoldenFilePath, actualOutput, "utf8");
            console.log(`Golden file updated: ${resolvedGoldenFilePath}`);
        } else {
            throw new Error(
                `Output does not match the golden file. To update the golden file, set UPDATE_GOLDEN_FILES=true and re-run the tests.\n` +
                `Golden file path: ${resolvedGoldenFilePath}`
            );
        }
    } else {
        console.log(`Output matches the golden file: ${resolvedGoldenFilePath}`);
    }
}

module.exports = { compareWithGoldenFile };
