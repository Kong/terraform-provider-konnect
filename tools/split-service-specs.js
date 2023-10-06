const fs = require("fs").promises;
const fg = require("fast-glob");
const yaml = require("js-yaml");
const path = require("path");
const traverse = require("traverse");
const { tryMkdir } = require("./helpers/fs");
const toolkit = require("oas-toolkit");

function isEmptyOperation(node) {
  const allowedKeys = Object.keys(node).filter((k) =>
    ["get", "post", "put", "patch", "delete"].includes(k)
  );
  return allowedKeys.length == 0;
}

function filterOperations(doc, callback, log) {
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (!this.node || !this.node["operationId"]) {
      return;
    }

    const result = callback(this.node);
    if (!result) {
      this.remove();

      if (isEmptyOperation(this.parent.node)) {
        this.parent.remove();
      }
    }
  });
}

async function main() {
  const baseDir = path.resolve(__dirname, "..");
  const projects = ["konnect", "portal"];

  for (let mode of projects) {
    let files = await fg(
      `${baseDir}/${mode}/definitions/**/computed/openapi.yaml`
    );
    for (let f of files){
    let complete = yaml.load(await fs.readFile(f));
    let name = f.replace("computed/openapi.yaml","");

    // If there's a nested v1, v2 etc folder, remove it
    if (path.basename(name).match(/v\d+/)) {
      name = path.dirname(name);
    }

    name = path.basename(name);

    let dev = filterOperations(complete, function (node) {
      return node["x-internal"] && node["x-unstable"];
    });

    let internal = filterOperations(complete, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    });

    let public = filterOperations(complete, function (node) {
      return !node["x-internal"];
    });

    // Remove unused components from each spec
    dev = toolkit.components.removeUnusedComponents(dev);
    internal = toolkit.components.removeUnusedComponents(internal);
    public = toolkit.components.removeUnusedComponents(public);

    // Remove unused tags too
    dev = toolkit.tags.removeUnusedTags(dev);
    internal = toolkit.tags.removeUnusedTags(internal);
    public = toolkit.tags.removeUnusedTags(public);

    // Write multiple files
    const outputDir = `${baseDir}/output/${mode}`;
    await tryMkdir(`${outputDir}/services/${name}`, { recursive: true });
    await fs.writeFile(`${outputDir}/services/${name}/dev.yaml`, yaml.dump(dev));
    await fs.writeFile(`${outputDir}/services/${name}/internal.yaml`, yaml.dump(internal));
    await fs.writeFile(`${outputDir}/services/${name}/public.yaml`, yaml.dump(public));
    }
  }
}

if (require.main === module) {
  main().catch((e) => {
    console.log(e);
    throw e;
  });
}
