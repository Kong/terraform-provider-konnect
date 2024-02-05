const traverse = require('traverse');

function isEmptyOperation(node) {
  const allowedKeys = Object.keys(node).filter((k) =>
    ["get", "post", "put", "patch", "delete"].includes(k)
  );
  return allowedKeys.length == 0;
}

function filterOperations(doc, callback, overrideKey) {
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (!this.node || !this.node["operationId"]) {
      return;
    }

    // If we're explicitly adding to a specific build, don't remove it
    if (this.node['x-override'] == overrideKey){
      this.node['x-override'] = undefined;
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

function filterSchemas(doc, shouldKeepSchema, shouldRemoveSchema, overrideKey) {
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (!this.node || !this.path.join('.').includes("components.schemas")) {
      return;
    }
    if (this.path.length !== 3) {
      return;
    }

    const schemaPath = this.path.join("/");

    // If it's a schema and it matches our selector, include any
    // paths that reference this schema
    if (shouldKeepSchema(this.node)) {
      annotateWithOverride(doc, schemaPath, overrideKey);
    }

    if (shouldRemoveSchema(this.node)) {
      this.remove();
      removeReferencesTo(doc, schemaPath, overrideKey);
    }

  });
}

function removeReferencesTo(doc, schemaPath) {
  return traverse(doc).forEach(function () {
    if (!this.node) {
      return;
    }
    if (this.node['$ref'] && this.node['$ref'] == `#/${schemaPath}`) {
      this.remove();
    }
  });
}

function annotateWithOverride(doc, schemaPath, overrideKey, alreadyVisited = {}) {
  traverse(doc).forEach(function () {
    if (!this.node) {
      return;
    }

    if (alreadyVisited[this.node['$ref']]){
      return;
    }

    if (this.node['$ref'] && this.node['$ref'] == `#/${schemaPath}`) {
      if (this.path[0] == 'components'){
        alreadyVisited[this.node['$ref']] = true;
        annotateWithOverride(doc, this.path.slice(0,3).join("/"), overrideKey, alreadyVisited);
        return;
      }

      // Annotate the operation so that the path isn't stripped
      let parent = this;
      while (parent && !parent.node['operationId']){
        parent = parent.parent;
      };

      if (parent){
        parent.node['x-override'] = overrideKey;
      }
    }

  });
}

module.exports = {
  filterOperations,
  filterSchemas
}