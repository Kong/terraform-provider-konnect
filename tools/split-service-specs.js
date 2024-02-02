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
  const projects = ["konnect", "portal", "internal"];

  for (let mode of projects) {
    const outputDir = `${baseDir}/build/services/${mode}`;

    let files = await fg(
      `${baseDir}/computed/${mode}/**/openapi.yaml`
    );
    for (let f of files){
    let complete = yaml.load(await fs.readFile(f));
    let name = f.replace(`${baseDir}/computed/${mode}/`,"").replace("openapi.yaml","");

    let dev = filterOperations(complete, function (node) {
      return node["x-internal"] && node["x-unstable"];
    });

    let internal = filterOperations(complete, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    });

    let public = filterOperations(complete, function (node) {
      return !node["x-internal"];
    });

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
