const fg = require("fast-glob");
const redoc = require('@redocly/cli/lib/commands/bundle')
const path = require('path')

const baseDir = path.resolve(__dirname, '..')
const redocConfigurationPath = path.join(baseDir, 'redocly.yaml')

// Loop over all `src` folders in `konnect`/`portal`/`internal` projects definitions to
// generate the `computed` folder to be consumed by the api tooling
async function main() {
    let projects = ["konnect", "portal", "internal"];
    if (process.argv.length > 2) {
        projects = process.argv.slice(2);
    }
    for (const mode of projects) {
        const files = await fg(`${baseDir}/${mode}/definitions/**/src/openapi.yaml`);
        for (const f of files) {
            const srcDir = path.dirname(f)
            const output = path.resolve(srcDir, '..', 'computed', 'openapi.yaml')
            const args = {
                config: redocConfigurationPath,
                output: output,
                apis: [f],
            }
            await redoc.handleBundle(args).catch(e => {
                console.error("Error while bundling API.", args, e)
                throw e
            })
        }
    }
}

if (require.main === module) {
    main().catch(e => { throw e })
}