const fg = require("fast-glob");
const path = require("path");
const yaml = require("js-yaml");
const fs = require("fs").promises;
const traverse = require("traverse");
const getProducts = require("./get-products");

const baseDir = path.resolve(__dirname, "..");

async function main() {
  let projects = getProducts();
  if (process.argv.length > 2) {
    projects = process.argv.slice(2);
  }
  for (const mode of projects) {
    const files = await fg(
      `${baseDir}/computed/${mode}/**/openapi.yaml`
    );
    for (const f of files) {
      console.log(`Processing ${f}`);
      let oas = yaml.load(await fs.readFile(f));
      oas = addUnstablePublicDescription(oas);
      await fs.writeFile(f, yaml.dump(oas));
    }
  }
}

function addUnstablePublicDescription(doc) {
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (!this.node || !this.node["x-unstable"]) {
      return;
    }

    if (this.node["x-internal"]) {
      return;
    }

    let description = this.node["description"] || "";
    this.node["description"] = `**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

${description}`.trim();
  });
}

if (require.main === module) {
  main();
}
