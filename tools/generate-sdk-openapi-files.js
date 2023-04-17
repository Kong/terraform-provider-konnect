const fs = require("fs").promises;
const fg = require("fast-glob");
const yaml = require("js-yaml");
const path = require("path");
const mergician = require("mergician");
const $RefParser = require("@apidevtools/json-schema-ref-parser");
const traverse = require("traverse");
const { transformFilterQueryParams } = require('./helpers/transformFilterQueryParams');
const { tryMkdir } = require("./helpers/fs");

const headers = {
  konnect: {
    openapi: "3.0.3",
    info: {
      title: "Konnect API",
      description: "The Konnect platform API",
    },
    servers: [
      {
        url: "https://global.api.konghq.com/v2",
        description: "Production",
      },
    ],
  },

  portal: {
    openapi: "3.0.3",
    info: {
      title: "Portal API",
      description: "Portal API",
    },
    servers: [
      {
        url: "https://custom.example.com",
        description: "Portal API",
      },
    ],
  },
};

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

function deepClone(serializable) {
  return JSON.parse(JSON.stringify(serializable));
}

async function copyInternals(basedir) {
  let files = await fg(`${basedir}/internal/definitions/**/computed/openapi.yaml`)

  for (const specPaths of files) {
    // keeping /servicename/<version> from the path
    const specPathDirname = path.dirname(specPaths).split('/').splice(-3, 2)
    const data = await fs.readFile(specPaths)
    const targetPath = path.join(basedir,'output', 'internal', ...specPathDirname)
    await tryMkdir(targetPath, { recursive: true })
    await fs.writeFile(
      path.join(targetPath, 'openapi.yaml'),
      data,
      { encoding: 'utf-8' }
    )
  }

}

async function main() {
  const baseDir = `${__dirname}/..`;
  const projects = ["konnect", "portal"];

  for (let mode of projects) {

    let files = await fg(`${baseDir}/${mode}/definitions/**/computed/openapi.yaml`);

    const openapiFiles = [];

    // Build a complete OAS
    for (let f of files) {
      let oas = yaml.load(await fs.readFile(f));
      oas = await $RefParser.bundle(oas);
      openapiFiles.push(oas);
    }

    const merger = mergician({
      appendArrays: true,
      dedupArrays: true,
    });
    const mergedSpecs = merger.apply(null, openapiFiles);

    // Overwrite known fields
    const complete = mergician(mergedSpecs, headers[mode]);

    if (!complete.paths) {
      console.log(`No paths found for ${mode}`);
      continue;
    }

    // Make entries in `tags` unique
    complete.tags = unique(complete.tags);

    // Parameters might be duplicated too
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
    const dev = filterOperations(complete, function (node) {
      return node["x-internal"] && node["x-unstable"];
    });

    const internal = filterOperations(complete, function (node) {
      return node["x-internal"] && !node["x-unstable"];
    });

    const public = filterOperations(complete, function (node) {
      return !node["x-internal"];
    });

    // Create additional spec set optimized for API Documentation rendering
    const [devApiDocs, internalApiDocs, publicApiDocs] = [
      transformFilterQueryParams(deepClone(dev)),
      transformFilterQueryParams(deepClone(internal)),
      transformFilterQueryParams(deepClone(public))
    ];

    // Write multiple files
    // @TODO: Remove unused components/tags etc
    const outputDir = `${baseDir}/output/${mode}`;
    await tryMkdir(outputDir, { recursive: true })
    await Promise.all([
      fs.writeFile(`${outputDir}/dev.yaml`, yaml.dump(dev)),
      fs.writeFile(`${outputDir}/dev-api-docs.yaml`, yaml.dump(devApiDocs)),
      fs.writeFile(`${outputDir}/internal.yaml`, yaml.dump(internal)),
      fs.writeFile(`${outputDir}/internal-api-docs.yaml`, yaml.dump(internalApiDocs)),
      fs.writeFile(`${outputDir}/public.yaml`, yaml.dump(public)),
      fs.writeFile(`${outputDir}/public-api-docs.yaml`, yaml.dump(publicApiDocs)),
    ]);
  }
  await copyInternals(baseDir)
}

if (require.main === module) {
  main().catch(e => { throw e })
}