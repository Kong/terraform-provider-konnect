const fs = require("fs").promises;
const fg = require("fast-glob");
const yaml = require("js-yaml");
const mergician = require("mergician");
const $RefParser = require("@apidevtools/json-schema-ref-parser");
const traverse = require("traverse");
const { transformFilterQueryParams } = require('./helpers/transformFilterQueryParams')

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
    security: [{ accessToken: [] }],
  },

  portal: {
    openapi: "3.0.3",
    info: {
      title: "Portal API",
      description: "Portal API",
    },
    servers: [
      {
        url: "https://custom.example.com/api",
        description: "Portal API",
      },
    ],
    security: [{ accessToken: [] }],
  },
};

(async function () {
  const projects = ["konnect", "portal"];

  for (let mode of projects) {
    const baseDir = `${__dirname}/..`;

    let files = await fg(`${baseDir}/${mode}/definitions/**/openapi.yaml`);

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
    try {
      await fs.mkdir(outputDir, { recursive: true });
    } catch (e) {
      if (e.code !== "EEXIST") {
        throw e;
      }
    }
    await Promise.all([
      fs.writeFile(`${outputDir}/dev.yaml`, yaml.dump(dev)),
      fs.writeFile(`${outputDir}/dev-api-docs.yaml`, yaml.dump(devApiDocs)),
      fs.writeFile(`${outputDir}/internal.yaml`, yaml.dump(internal)),
      fs.writeFile(`${outputDir}/internal-api-docs.yaml`, yaml.dump(internalApiDocs)),
      fs.writeFile(`${outputDir}/public.yaml`, yaml.dump(public)),
      fs.writeFile(`${outputDir}/public-api-docs.yaml`, yaml.dump(publicApiDocs)),
    ]);
  }
})();

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
