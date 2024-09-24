const fg = require("fast-glob");
const path = require("path");
const yaml = require("js-yaml");
const fs = require("fs").promises;
const getProducts = require("./get-products");
const traverse = require("traverse");
const mergician = require("mergician");

const $RefParser = require("@apidevtools/json-schema-ref-parser");

const baseDir = path.resolve(__dirname, "..");

async function main() {
  let projects = getProducts();
  if (process.argv.length > 2) {
    projects = process.argv.slice(2);
  }
  for (const mode of projects) {
    const files = await fg(`${baseDir}/computed/${mode}/**/openapi.yaml`);
    for (const f of files) {
      console.log(`[expand-allOf] ${f}`);
      let oas = yaml.load(await fs.readFile(f));
      oas = await expandAllOf(oas);
      await fs.writeFile(f, yaml.dump(oas));
    }
  }
}

async function expandAllOf(oas) {
  const nodesToUpdate = [];

  const r = traverse(oas).forEach(function (x) {
    if (this.node && this.node["x-flatten-allOf"]) {
      nodesToUpdate.push(this);
    }
  });

  // We can't call async functions inside `traverse` so we
  // need to keep track of nodes and do it here
  for (const node of nodesToUpdate) {
    await updateNode(oas, node);
  }

  return r;
}

async function updateNode(oas, item) {
  // Generate all interim paths from the root to the current node
  const toDeref = item.path;
  toDeref.push("allOf");
  const pathParts = toDeref.map(
    (_, i) =>
      "#/" +
      item.path
        .slice(0, i + 1)
        .map((s) => s.replace("/", "~1"))
        .join("/")
  );

  pathParts.unshift("#");

  // Reference any refs in the current path
  const x = await $RefParser.dereference(oas, {
    dereference: {
      excludedPathMatcher: function (path) {
        return !pathParts.includes(path) && !path.match(/allOf\/\d+$/);
      },
      circular: "ignore",
    },
  });

  const subSchema = traverse(x).get(toDeref);
  item.update(mergician({}, ...subSchema));
}

main();
