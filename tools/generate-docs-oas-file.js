const fs = require('fs').promises
const path = require('path')
const yaml = require("js-yaml");
const { transformFilterQueryParams } = require('./helpers/transformFilterQueryParams');
const { tryMkdir } = require('./helpers/fs')

const fail = (msg, ...ctx) => {
  console.error(msg, ...ctx)
  process.exit(1)
}

(async () => {
  const [oasName, oasPath] = process.argv.slice(2, 4);

  if (!oasName || !oasPath) {
    fail('Missing required arguments! Usage: node generate-docs-oas-file <oas_name> <oas_path>');
  }

  const baseDir = path.resolve(__dirname, '..');
  const absoluteSpecPath = path.resolve(baseDir, oasPath);
  const absoluteOutputPath = path.resolve(baseDir, 'output', 'konnect', `${oasName}-docs.yaml`);

  if (!(await fs.stat(absoluteSpecPath))) {
    fail(`No OAS found at: ${absoluteSpecPath}`);
  }

  console.log(`Loaded ${oasName} OAS from ${absoluteSpecPath}`)
  const spec = yaml.load(await fs.readFile(absoluteSpecPath))
  
  console.log(`Writing transformed docs ${oasName} file to ${absoluteOutputPath}`)
  await tryMkdir(path.dirname(absoluteOutputPath), { recursive: true })
  await fs.writeFile(absoluteOutputPath, yaml.dump(transformFilterQueryParams(spec)))
})()
