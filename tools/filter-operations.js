const fs = require("fs").promises;
const yaml = require("js-yaml");
const path = require("path");
const traverse = require("traverse");
const toolkit = require("oas-toolkit");

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

function removeNodeWithKey(doc, key){
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (this.key == key) {
      this.remove();
    }
  });
}

function isEmptyOperation(node) {
  const allowedKeys = Object.keys(node).filter((k) =>
    ["get", "post", "put", "patch", "delete"].includes(k)
  ); 
  return allowedKeys.length == 0;
}
async function main() {
    const sdkName = process.env.SDK_NAME;
    const baseDir = `${__dirname}/..`;
    const tfFile = `${baseDir}/build/complete/${sdkName}/public.yaml`;
    complete = yaml.load(await fs.readFile(tfFile));

    let tf = complete;

    if (process.env.REQUIRE_SPEAKEASY_ENTITY_OPERATION === "1") {
      // Filter down paths to the groups that we need
      tf = filterOperations(complete, function (node) {
        const pluginOperations = [
          "create-plugin",
          "delete-plugin",
          "get-plugin",
          "fetch-plugin-schema",
          "upsert-plugin",
        ];
        return node["x-speakeasy-entity-operation"] || pluginOperations.includes(node.operationId);
      });
    }

    tf = removeNodeWithKey(tf, 'x-examples');
    tf = removeNodeWithKey(tf, 'examples');

    // Remove unused components from each spec
    tf = toolkit.components.removeUnusedComponents(tf);

    await fs.writeFile(tfFile, yaml.dump(tf));
}

if (require.main === module) {
  main().catch((e) => {
    console.log(e);
    throw e;
  });
}
