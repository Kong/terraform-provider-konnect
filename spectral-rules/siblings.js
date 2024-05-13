function isObject(value) {
  return value !== null && typeof value === 'object' && !Array.isArray(value);
}

function getParentValue(document, path) {
  if (path.length === 0) {
    return null;
  }

  let piece = document;

  for (let i = 0; i < path.length - 1; i += 1) {
    if (!isObject(piece)) {
      return null;
    }

    piece = piece[path[i]];
  }

  return piece;
}

export default function (targetVal, opts, { documentInventory, path }) {
    const value = getParentValue(documentInventory.resolved, path);
    if (!isObject(value)) return;

    const results = [];
    const actualObjPath = path.slice(0, -1);

    const keys = Object.keys(value);
    if (keys.length === 1) {
      return;
    }

    if (keys.includes("properties")) {
      results.push({
        message: `allOf, oneOf and anyOf must not have a 'properties' sibling`,
        path: actualObjPath
      });
    }

    return results;
  };