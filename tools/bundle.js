const fg = require("fast-glob");
const { handleBundle: redocBundle } = require('@redocly/cli/lib/commands/bundle')
const { loadConfigAndHandleErrors: redocConfig } = require('@redocly/cli/lib/utils');
const path = require('path')
const getProductFiles = require("./get-product-files");
const getProducts = require("./get-products");

const baseDir = path.resolve(__dirname, '..')
const redocConfigurationPath = path.join(baseDir, 'redocly.yaml')

// Loop over all `src` folders in project definitions to
// generate the `computed` folder to be consumed by the api tooling
async function main() {
    let allFiles = false
    let projects = getProducts();
    if (process.argv.length > 2) {
        const arg = process.argv.slice(2)
        if (arg[0] === 'all-files') {
            allFiles = true
        } else {
            projects = arg
        }
    } else {
        // If we're not building a specific product, bundle _everything_
        // as we need service OAS files in computed to generate per-service SDKs
        allFiles = true;
    }

    for (const mode of projects) {
        let files = []
        if (allFiles) {
            files = await fg(`${baseDir}/src/${mode}/definitions/**/openapi.yaml`);
        } else {
            files = getProductFiles(mode).map(f => `${baseDir}/${f}`)
        }

        for (const f of files) {
            const srcDir = path.dirname(f)
            const output = path.join(
                srcDir.replace(`${baseDir}/src`, `${baseDir}/computed`).replace('definitions/', ''),
                "openapi.yaml"
            );
            const args = {
                output: output,
                apis: [f],
            };

            const config = await redocConfig({
                configPath: redocConfigurationPath
            });

            await redocBundle(args, config);
        }
    }
}

if (require.main === module) {
    main()
}
