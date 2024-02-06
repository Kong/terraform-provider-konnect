const fs = require("fs").promises;
const fg = require("fast-glob");
const yaml = require("js-yaml");
const path = require("path");
const traverse = require("traverse");
const { tryMkdir } = require("./helpers/fs");
const toolkit = require("oas-toolkit");
const { filterOperations, filterSchemas, removeProperty } = require("./helpers/filterOas");

async function main() {
  const baseDir = path.resolve(__dirname, "..");
  let projects = ["konnect", "portal", "internal"];
  if (process.argv.length > 2) {
    projects = process.argv.slice(2);
  }

  for (let mode of projects) {
    const outputDir = `${baseDir}/build/services/${mode}`;

    let files = await fg(
      `${baseDir}/computed/${mode}/**/openapi.yaml`
    );
    for (let f of files){
      if (!f.includes("portal") || !f.includes("applications")){
        //continue;
      }
    let complete = yaml.load(await fs.readFile(f));
    let name = f.replace(`${baseDir}/computed/${mode}/`,"").replace("openapi.yaml","");

    // Remove any annotated schemas (and the parent object)
    let dev = filterSchemas(complete, function (node) {
      return node["x-internal"] && node["x-unstable"];
    }, function(node){ return false; }, 'dev');

    let internal = filterSchemas(complete, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    }, function(node){ return false; }, 'internal');

    let public = filterSchemas(complete, function (node) {
      return !node["x-internal"];
    }, function(node){ return node["x-internal"] });

    // Remove any annotated operations
    dev = filterOperations(dev, function (node) {
      return node["x-internal"] && node["x-unstable"];
    }, 'dev');

    internal = filterOperations(internal, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    }, 'internal');

    public = filterOperations(public, function (node) {
      return !node["x-internal"];
    }, 'public');

    // Remove the `x-override` field
    const fields = ['x-override', 'x-internal', 'x-unstable'];
    dev = removeProperty(dev, fields);
    internal = removeProperty(internal, fields);
    public = removeProperty(public, fields);

    await tryMkdir(`${outputDir}/${name}`, { recursive: true });
    const builds = {
      dev,
      internal,
      public,
    };

    for (let build in builds){
      let oas = builds[build];
      if (containsPaths(oas)){
        oas = toolkit.components.removeUnusedComponents(oas);
        oas = toolkit.tags.removeUnusedTags(oas);
        await fs.writeFile(`${outputDir}/${name}/${build}.yaml`, yaml.dump(oas));
      }
    }
    }
  }
}
function containsPaths(oas) {
  return oas && oas.paths && Object.keys(oas.paths).length > 0;
}

if (require.main === module) {
  main().catch((e) => {
    console.log(e);
    throw e;
  });
}
