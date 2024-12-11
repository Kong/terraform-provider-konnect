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

    // List of fields that are not allowed to be siblings
    const disallowedSiblingFields = ['properties'];

    // If allOf is present and gets flattened, description & required are not allowed to be a sibling. It should be present under allOf
    if (keys.includes('allOf') && keys.includes('x-flatten-allOf')) {
      disallowedSiblingFields.push('required', 'description');
    }

    const foundDisallowedSiblings = disallowedSiblingFields.filter(field => keys.includes(field));

    if (foundDisallowedSiblings.length > 0) {
      results.push({
        message: `allOf, oneOf and anyOf must not have the following sibling(s): ${foundDisallowedSiblings.join(', ')}`,
        path: actualObjPath
      });
    }

    return results;
  };