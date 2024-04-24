const traverse = require('traverse');
const get = require('lodash.get');

function isEmptyOperation(node) {
  const allowedKeys = Object.keys(node).filter((k) =>
    ["get", "post", "put", "patch", "delete"].includes(k)
  );
  return allowedKeys.length == 0;
}

function removeProperty(doc, fields) {
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (!this.node) {
      return;
    }
    if (fields.includes(this.key)){
      this.remove();
    }
  });
}

function annotateWithTarget(doc, targetIsValid, callback, targetKey) {
  return traverse(doc).forEach(function (node) {
    if (!this.node) {
      return;
    }

    if (!targetIsValid(this)){
      return;
    }

    const result = callback(this.node);
    if (result) {
      if (!this.node['x-target-specification']) {
        this.node['x-target-specification'] = [];
      }
      this.node['x-target-specification'].push(targetKey);
    }
  });

}

function removeSchemasNotIn(doc, target) {
  doc = traverse(doc).clone();
  return forEachSchema(doc, function (t) {
    let shouldRemove = false;

    if (t.node['x-target-specification']?.length > 0) {
      shouldRemove = !t.node['x-target-specification'].includes(target);
    }

    if (shouldRemove) {
      removeReferencesTo(doc, t.path.join('/'));
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

function includePathsUsedBySchema(doc, schemaPath, targetSpec, alreadyVisited = {}) {

  if (!targetSpec){
    const val = get(doc, schemaPath.split('/').join('.'));
    targetSpec = val['x-target-specification'];
    if (!targetSpec){
      return;
    }
  }

  traverse(doc).forEach(function () {
    if (!this.node) {
      return;
    }

    if (alreadyVisited[this.node['$ref']]){
      return;
    }

    // Find the target specification for this schema

    if (this.node['$ref'] && this.node['$ref'] == `#/${schemaPath}`) {
      if (this.path[0] == 'components'){
        alreadyVisited[this.node['$ref']] = true;
        includePathsUsedBySchema(doc, this.path.slice(0,3).join("/"), targetSpec, alreadyVisited);
        return;
      }

      // Annotate the operation so that the path isn't stripped
      let parent = this;
      while (parent && !parent.node['operationId']){
        parent = parent.parent;
      };

      if (parent){
        parent.node['x-target-specification'] = Array.from(new Set(parent.node['x-target-specification'].concat(targetSpec)));
      }
    }

  });
}

function forEachSchema(doc, callback) {
  return traverse(doc).forEach(function () {
    if (!this.node || !this.path.join('.').includes("components.schemas")) {
      return;
    }
    if (this.path.length !== 3) {
      return;
    }
    callback(this);
  });
}

function filterPathsToTarget(doc, target) {
  doc = traverse(doc).clone();
  return traverse(doc).forEach(function () {
    if (!this.node || !this.node['operationId']) {
      return;
    }

    // If we're targeting public paths, then x-target-specification is missing
    if (target == 'public' && !this.node['x-target-specification']) {
      return;
    }

    if (!this.node['x-target-specification'] || !this.node['x-target-specification'].includes(target)){
      this.remove();

      if (isEmptyOperation(this.parent.node)){
        this.parent.remove();
      }
    }


  });

}

function splitByVisibility(doc) {

  // Remove speakeasy annotations
  if (!process.env.KEEP_SPEAKEASY_ANNOTATIONS){
    traverse(doc).forEach(function () {
      if (!this.node) {
        return;
      }
      if (this.key && this.key.startsWith("x-speakeasy")){
        this.remove();
      }
    });
  }

  function isDev(node){
    return node["x-internal"] && node["x-unstable"];
  }

  function isInternal(node){
    return node["x-internal"] && !node["x-unstable"];
  }

  function isPublic(node){
    return !node["x-internal"] && !node["x-unstable"];
  }

  // Annotate any paths with a target specification
  function isPath(t){
    if (!t.node || !t.node["operationId"]) {
      return false;
    }
    return true;
  }

  annotateWithTarget(doc, isPath, isDev, 'dev');
  annotateWithTarget(doc, isPath, isInternal, 'internal');
  annotateWithTarget(doc, isPath, isPublic, 'public');

  // Annotate any schemas with a target specification
  function isSchema(t){
    if (!t.node || !t.path.join('.').includes("components.schemas")) {
      return false;
    }
    if (t.path.length !== 3) {
      return false;
    }
    return true;
  }

  annotateWithTarget(doc, isSchema, isDev, 'dev');
  annotateWithTarget(doc, isSchema, isInternal, 'internal');

  // For any schema that are annotated with a target specification,
  // annotate the paths that use them
  forEachSchema(doc, function (t) {
    includePathsUsedBySchema(doc, t.path.join('/'))
  });

  // Split in to three separate OAS files
  let dev = filterPathsToTarget(doc, 'dev');
  let internal = filterPathsToTarget(doc, 'internal');
  let public = filterPathsToTarget(doc, 'public');

  // For each OAS file, if any of the used schemas are annotated with 'internal' etc
  // Then remove any properties from paths that use them unless the target matches
  dev = removeSchemasNotIn(dev, 'dev');
  internal = removeSchemasNotIn(internal, 'internal');
  public = removeSchemasNotIn(public, 'public');

  // Finally, remove any properties that are explicitly marked as 'internal' etc
  function propertyIsDev(v){
    return v.includes('x-internal') && v.includes('x-unstable');
  }
  function propertyIsInternal(v){
    return v.includes('x-internal') && !v.includes('x-unstable');
  }

  function propertyIsPublic(v){
    return !v.includes('x-internal') && !v.includes('x-unstable');
  }

  forEachSchema(dev, function (t) {
    const a = t.node['x-property-annotations'];
    if (!a){ return; }
    for (let k in a) {
      if (propertyIsInternal(a[k]) ){
        delete t.node.properties[k];
      }
    }
  });

  forEachSchema(internal, function (t) {
    const a = t.node['x-property-annotations'];
    if (!a){ return; }
    for (let k in a) {
      if (propertyIsDev(a[k])){
        delete t.node.properties[k];
      }
    }
  });

  forEachSchema(public, function (t) {
    const a = t.node['x-property-annotations'];
    if (!a){ return; }
    for (let k in a) {
      if (propertyIsPublic(a[k])){
        delete t.node.properties[k];
      }
    }
  });

  // Remove the `x-property-annotations` field
  const fields = [
    'x-property-annotations',
    'x-target-specification',
    'x-internal',
    'x-unstable'
  ];

  dev = removeProperty(dev, fields);
  internal = removeProperty(internal, fields);
  public = removeProperty(public, fields);

  return {dev, internal, public};
}

module.exports = {
  annotateWithTarget,
  removeProperty,
  includePathsUsedBySchema,
  forEachSchema,
  filterPathsToTarget,
  removeSchemasNotIn,
  splitByVisibility
}