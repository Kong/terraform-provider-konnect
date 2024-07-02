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
    const baseDir = `${__dirname}/..`;
    const tfFile = `${baseDir}/build/complete/terraform/public.yaml`;
    complete = yaml.load(await fs.readFile(tfFile));

    // Filter down paths to the groups that we need
    let tf = filterOperations(complete, function (node) {
      return node["x-speakeasy-entity-operation"];
    });

    tf = removeNodeWithKey(tf, 'x-examples');
    tf = removeNodeWithKey(tf, 'examples');

    // Remove unused components from each spec
    tf = toolkit.components.removeUnusedComponents(tf);

    // Make the order deterministic
    // Top level keys
    const topOrder = {
      "openapi": 1000,
      "info": 900,
      "servers": 800,
      "paths": 700,
      "components": 600,
      "security": 500,
      "tags": 400,
      "externalDocs": 300,
    }

    tf = Object.fromEntries(
      Object.entries(tf).sort(function (a, b) {
        return topOrder[b[0]] - topOrder[a[0]];
      })
    );

    // Paths should be alphabetical
    tf.paths = Object.fromEntries(
      Object.entries(tf.paths).sort(function (a, b) {
        return a[0].localeCompare(b[0]);
      })
    );

    // Items in /paths should be sorted
    const pathOrder = {
      "parameters": 1000,
      "get": 900,
      "post": 800,
      "put": 700,
      "patch": 600,
      "delete": 500,
    }

    for (let path in tf.paths) {
      tf.paths[path] = Object.fromEntries(
        Object.entries(tf.paths[path]).sort(function (a, b) {
          return pathOrder[b[0]] - pathOrder[a[0]];
        })
      );
    }

    // As should schemas
    tf.components.schemas = Object.fromEntries(
      Object.entries(tf.components.schemas).sort(function (a, b) {
        return a[0].localeCompare(b[0]);
      })
    );

    // And items within schemas
    for (let type in tf.components) {
      if (tf.components[type]) {
        tf.components[type] = Object.fromEntries(
          Object.entries(tf.components[type]).sort(function (a, b) {
            return a[0].localeCompare(b[0]);
          })
        );
      }
    }

    await fs.writeFile(tfFile, yaml.dump(tf));
}

if (require.main === module) {
  main().catch((e) => {
    console.log(e);
    throw e;
  });
}
