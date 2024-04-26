const fs = require("fs").promises;
const fg = require("fast-glob");
const yaml = require("js-yaml");
const path = require("path");
const { tryMkdir } = require("./helpers/fs");
const toolkit = require("oas-toolkit");
const { splitByVisibility } = require("./helpers/filterOas");
const getProducts = require("./get-products");

async function main() {
  const baseDir = path.resolve(__dirname, "..");
  let projects = getProducts();
  if (process.argv.length > 2) {
    projects = process.argv.slice(2);
  }

  for (let mode of projects) {
    const outputDir = `${baseDir}/build/services/${mode}`;

    let files = await fg(
      `${baseDir}/computed/${mode}/**/openapi.yaml`
    );
    for (let f of files){
      if (!f.includes("portal") || !f.includes("applications")){
        //continue;
      }
    let complete = yaml.load(await fs.readFile(f));
    let name = f.replace(`${baseDir}/computed/${mode}/`,"").replace("openapi.yaml","");

    const { dev, internal, public } = splitByVisibility(complete);

    await tryMkdir(`${outputDir}/${name}`, { recursive: true });
    const builds = {
      dev,
      internal,
      public,
    };

    for (let build in builds){
      let oas = builds[build];
      if (containsPaths(oas)){
        oas = toolkit.components.removeUnusedComponents(oas);
        oas = toolkit.tags.removeUnusedTags(oas);
        await fs.writeFile(`${outputDir}/${name}/${build}.yaml`, yaml.dump(oas));
      }
    }
    }
  }
}
function containsPaths(oas) {
  return oas && oas.paths && Object.keys(oas.paths).length > 0;
}

if (require.main === module) {
  main().catch((e) => {
    console.log(e);
    throw e;
  });
}
