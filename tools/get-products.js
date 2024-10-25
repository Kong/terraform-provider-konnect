const fs = require("fs");
const yaml = require("js-yaml");

module.exports = function (product) {
  const config = yaml.load(
    fs.readFileSync(`${__dirname}/../products.yaml`, "utf8")
  );

  let products = Object.keys(config.products);

  // These products are only built when requested
  // e.g. ./tools/run.sh terraform
  const onDemandProducts = ['terraform', 'go-sdk'];
  return products.filter(p => {
    return !onDemandProducts.includes(p)
  });

};

if (require.main == module) {
  const args = process.argv.slice(2);

  const output = module.exports.apply(null, args);

  if (args[0] == "--plain") {
    console.log(output.join(" "));
    process.exit(0);
  }

  console.log(JSON.stringify(output));
}