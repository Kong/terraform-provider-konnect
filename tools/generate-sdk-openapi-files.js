const fs = require("fs").promises;
const yaml = require("js-yaml");
const path = require("path");
const traverse = require("traverse");
const { tryMkdir } = require("./helpers/fs");
const toolkit = require("oas-toolkit");
const { filterSchemas, removeProperty } = require("./helpers/filterOas");

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
  let projects = ["konnect", "portal"];
  if (process.argv.length > 2) {
    projects = process.argv.slice(2);
  }

  for (let mode of projects) {
    if (mode == "internal") {
      console.log("Skipping combined file generation: Internal APIs do not produce a combined complete.yaml file");
      continue;
    }

    const f = `${baseDir}/computed/${mode}/complete.yaml`;

    try {
      await fs.access(f, fs.constants.F_OK);
    } catch {
      console.log(`No file found for ${mode}`);
      continue;
    }

    let complete = yaml.load(await fs.readFile(f));

    if (!complete?.paths) {
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

    let dev = filterSchemas(complete, function (node) {
      return node["x-internal"] && node["x-unstable"];
    }, function (node) { return false; }, 'dev');

    let internal = filterSchemas(complete, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    }, function (node) { return false; }, 'internal');

    let public = filterSchemas(complete, function (node) {
      return !node["x-internal"];
    }, function (node) { return node["x-internal"] });


    // Filter down paths to the groups that we need
    dev = filterOperations(dev, function (node) {
      return node["x-internal"] && node["x-unstable"];
    });

    internal = filterOperations(internal, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    });

    public = filterOperations(public, function (node) {
      return !node["x-internal"];
    });

    // Remove the `x-override`, `x-internal` and `x-unstable` fields
    const fields = ['x-override', 'x-internal', 'x-unstable'];
    dev = removeProperty(dev, fields);
    internal = removeProperty(internal, fields);
    public = removeProperty(public, fields);

    // Write multiple files
    const outputDir = `${baseDir}/build/complete/${mode}`;
    await tryMkdir(outputDir, { recursive: true });

    const builds = {
      dev,
      internal,
      public,
    };

    for (let build in builds) {
      let oas = builds[build];
      if (containsPaths(oas)) {
        oas = toolkit.components.removeUnusedComponents(oas);
        oas = toolkit.tags.removeUnusedTags(oas);
        await fs.writeFile(`${outputDir}/${build}.yaml`, yaml.dump(oas));
      }
    }
  }
}

function containsPaths(oas) {
  return oas && oas.paths && Object.keys(oas.paths).length > 0;
}

if (require.main === module) {
  main().then(() => {
    console.log("Done");
  }).catch((e) => {
    console.error(e);
    throw e;
  });
}
