const productFiles = require("./get-product-files");  
const yaml = require('js-yaml')
const fs = require('fs').promises;
const traverse = require('traverse');
const url = require('url');

module.exports = async function (product) {

  const globalApis = [
    "identity/v2",
    "identity/v3",
    "cloud-gateways/v2",
    "ksearch/v1",
    "organization-lifecycle/v0",
    "serverless-cloud-gateways/v0"
  ]

  const toUpdate = productFiles(product).filter((f) => {
    return globalApis.some((g) => f.includes(g));
  }).map((f) => {
    // Convert to computed path
    f = f.replace('src/konnect/definitions', 'computed/konnect');
    return f;
  });

  for (const f of toUpdate) {
    let oas = yaml.load(await fs.readFile(f));
    oas = removePerOperationServerBlock(oas);
    await fs.writeFile(f, yaml.dump(oas, {noRefs: true}));
  }

  return ({
    "action": "remove-global-server-block",
    product,
    files: toUpdate
  });
};

function removePerOperationServerBlock(doc) {
    doc = traverse(doc).clone();
  
    return traverse(doc).forEach(function () {
      if (!this.node || !this.node["operationId"]) {
        return;
      }

      if (this.node["servers"] && this.node["x-servers-added"]) {
        delete this.node["servers"];
        delete this.node["x-servers-added"];
      }
    });
}


(async function(){
  if (require.main == module) {
    const args = process.argv.slice(2);

    const output = await module.exports.apply(null, args);
    console.log(JSON.stringify(output));
  }
})();
