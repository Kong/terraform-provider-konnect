const fs = require("fs");
const yaml = require("js-yaml");
const fg = require("fast-glob");
const minimatch = require("minimatch");

module.exports = function (product) {
  const config = yaml.load(
    fs.readFileSync(`${__dirname}/../products.yaml`, "utf8")
  );

  const productConfig = config.products[product];

  // Load all the matching files
  let allFiles = [];
  for (let m of productConfig.match) {
    allFiles = allFiles.concat(fg.sync(m));
  }

  // And remove any items in notMatch
  for (let nm of productConfig.notMatch) {
    allFiles = allFiles.filter((f) => {
      return !minimatch(f, nm);
    });
  }

  return allFiles;
};

if (require.main == module) {
  const args = process.argv.slice(2);

  const output = module.exports.apply(null, args);

  if (args[1] == "--plain") {
    console.log(output.join(" "));
    process.exit(0);
  }

  console.log(JSON.stringify(output));
}
