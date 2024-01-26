const fs = require("fs").promises;
const yaml = require("js-yaml");
const path = require("path");
const traverse = require("traverse");
const { tryMkdir } = require("./helpers/fs");
const toolkit = require("oas-toolkit");

function filterOperations(doc, callback, log) {
  doc = traverse(doc).clone();

  return traverse(doc).forEach(function () {
    if (!this.node || !this.node["operationId"]) {
      return;
    }

    const result = callback(this.node);
    if (!result) {
      this.remove();

      if (isEmptyOperation(this.parent.node)) {
        this.parent.remove();
      }
    }
  });
}

function isEmptyOperation(node) {
  const allowedKeys = Object.keys(node).filter((k) =>
    ["get", "post", "put", "patch", "delete"].includes(k)
  );
  return allowedKeys.length == 0;
}

function unique(s) {
  if (!s) {
    return s;
  }
  return [...new Set(s.map((n) => JSON.stringify(n)))].map((n) =>
    JSON.parse(n)
  );
}

async function main() {
  const baseDir = path.join(__dirname, "..");
  const projects = ["konnect", "portal"];

  for (let mode of projects) {
    const f = `${baseDir}/computed/${mode}/complete.yaml`;
    let complete = yaml.load(await fs.readFile(f));

    if (!complete.paths) {
      console.log(`No paths found for ${mode}`);
      continue;
    }

    // Parameters might be duplicated
    for (let k in complete.paths) {
      const p = complete.paths[k];
      if (p.parameters) {
        p.parameters = unique(p.parameters);
      }
      for (verb in p) {
        const op = p[verb];
        if (!op.parameters) {
          continue;
        }
        op.parameters = unique(op.parameters);
      }
    }

    // Sort by paths alphabetically to make diffs deterministic
    // This makes the OAS oddly ordered, but any UI will group by tags anyway
    complete.paths = Object.fromEntries(
      Object.entries(complete.paths).sort(function (a, b) {
        return a[0].localeCompare(b[0]);
      })
    );

    complete.security = unique(complete.security);

    // Filter down paths to the groups that we need
    let dev = filterOperations(complete, function (node) {
      return node["x-internal"] && node["x-unstable"];
    });

    let internal = filterOperations(complete, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    });

    let public = filterOperations(complete, function (node) {
      return !node["x-internal"];
    });

    // Remove unused components from each spec
    dev = toolkit.components.removeUnusedComponents(dev);
    internal = toolkit.components.removeUnusedComponents(internal);
    public = toolkit.components.removeUnusedComponents(public);

    // Remove unused tags too
    dev = toolkit.tags.removeUnusedTags(dev);
    internal = toolkit.tags.removeUnusedTags(internal);
    public = toolkit.tags.removeUnusedTags(public);

    // Write multiple files
    const outputDir = `${baseDir}/build/complete/${mode}`;
    await tryMkdir(outputDir, { recursive: true });
    await Promise.all([
      fs.writeFile(`${outputDir}/dev.yaml`, yaml.dump(dev)),
      fs.writeFile(`${outputDir}/internal.yaml`, yaml.dump(internal)),
      fs.writeFile(`${outputDir}/public.yaml`, yaml.dump(public)),
    ]);
  }
}

if (require.main === module) {
  main().catch((e) => {
    console.log(e);
    throw e;
  });
}
