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
      oas = await setServersBlock(oas, f);
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

async function setServersBlock(doc, file) {
  console.log(file);
  // We only want to process Konnect files
  if (!file.includes("computed/konnect")) {
    return doc;
  }

  // Make sure global servers block is set only for global endpoints
  if (doc.servers.length == 1 && doc.servers[0].url.includes("global.api.konghq.com")) {
    return doc;
  }

  // Extract the path from the current doc
  const url = doc.servers[0].url;
  const path = url.split(".konghq.com")[1];

  const konnectServers = yaml.load(await fs.readFile(`${baseDir}/src/common/konnect.yaml`));
  const s = konnectServers.servers.map((s) => {
    s.url = s.url + path;
    return s;
  });

  const geoServers = s.filter((s) => !s.url.includes("global.api.konghq.com"));
  doc.servers = geoServers;

  return doc;
}

if (require.main === module) {
  main();
}

