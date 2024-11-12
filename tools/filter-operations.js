const fs = require("fs").promises;
const yaml = require("js-yaml");
const path = require("path");
const traverse = require("traverse");
const toolkit = require("oas-toolkit");

function removeDefaults(doc) {
  doc = traverse(doc).clone();

  // Remove default values, unless it's a foreign key or an empty array
  return traverse(doc).forEach(function () {
    // Remove default values from all properties
    const isPropertyDefault =
      this.key === "default" && this.parent.parent.key === "properties";
    const isArrayDefault =
      this.key === "default" && this.parent.key === "items";
    if (isPropertyDefault || isArrayDefault) {
      this.remove();
    }
  });
}

function addCustomDefaults(doc) {
  doc = traverse(doc).clone();

  // Remove default values, unless it's a foreign key or an empty array
  return traverse(doc).forEach(function () {
    // Anything marked as a foreign key should take values from the plan only
    // This allows services, routes etc to be unlinked by removing a value in the manifest
    const isForeignKey = this.node && this.node["x-foreign"] === true;
    if (isForeignKey) {
      this.node["x-speakeasy-terraform-plan-only"] = true;
      this.node["nullable"] = true;
      this.node["default"] = null;
      return;
    }

    // Fix OIDC scopes default value
    if (this.path[2] == "OpenidConnectPluginConfig") {
      if (this.path[this.path.length - 1] == "config") {
        if (!this.node["required"]) {
          this.node["required"] = [];
        }
        this.node["required"].push("scopes");
      }

      if (this.path[6] == "scopes") {
        const isArrayItemsDefault =
          this.key === "default" && this.parent.key === "items";

        if (
          isArrayItemsDefault &&
          this.parent.parent &&
          this.parent.parent.key == "scopes"
        ) {
          this.remove();
        }
        if (
          this.parent &&
          this.parent.key == "scopes" &&
          this.key == "default"
        ) {
          this.update([]);
        }
      }
    }
  });
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

function removeNodeWithKey(doc, key) {
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
      return (
        node["x-speakeasy-entity-operation"] ||
        pluginOperations.includes(node.operationId)
      );
    });
  }

  tf = removeDefaults(tf);
  tf = addCustomDefaults(tf);

  tf = removeNodeWithKey(tf, "x-examples");
  tf = removeNodeWithKey(tf, "examples");

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
