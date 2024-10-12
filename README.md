# platform-api

[Platform API specs](https://kong-platform-api.netlify.app/)

This repository contains OpenAPI specifications for Kong's APIs. This includes Konnect Platform, Konnect Portal and on-prem APIs.

Specifications in this repository are the source of truth. Any changes to an API should be raised as a PR for discussion before changes are made to the underlying API. If you need a copy of the API specification in your service repo (e.g. for contract testing), add it as a destination in [raise-pr-on-change.json](https://github.com/Kong/platform-api/blob/main/.github/raise-pr-on-change.json)

GitHub Actions are used to perform the following tasks:

- Lint APIs to ensure they meet our [AIP definitions](https://kong-aip.netlify.app/)
- Generate SDKs for use internally (and externally eventually)

## Useful documentation

* [API design workflow](https://github.com/Kong/platform-api/blob/main/docs/api-design-workflow.md)
* [Understanding x-internal and x-unstable](https://github.com/Kong/platform-api/blob/main/docs/openapi-annotations.md)
* [Linting rules](https://github.com/Kong/platform-api/blob/main/docs/linting-rules.md)

## Contributing

1. Edit your service in `src/<product>/definitions`.
2. Run `npm ci` to install dependencies
3. Run `./tools/run.sh`. You can use `./tools/run.sh PRODUCT_NAME` to build a specific product. Available products: `konnect`, `portal`, `internal`.
4. Commit everything (even the `build` folder) and raise a PR

### Konnect Control Planes

1. Konnect core-entites are generated from https://github.com/Kong/kong-admin-spec-generator. Specifically,
    a. src/konnect/definitions/control-planes-config/v2/generated.yaml
    b. src/konnect/definitions/control-planes-config/v2/generated-with-plugins.yaml
2. If there are changes to the above files, it would be appreciated if corresponding changes are also pushed to the generator.
3. Linter will complain if these files aren't formated. After generating from the admin spec generator, run `make lint-fix` to format these files.

## How this repo works

This repo has two top level folders, `build` and `src`. Contributors work in `src`, using all the same `x-unstable` and `x-internal` annotations that we use today.

- There is a build step that updates `build/konnect/<API_NAME>/{dev,internal,public}.yaml` from the current state of `src` (this is `./tools/run.sh`)
- The build directory is committed to the repo rather than just being a transient artifact
- Any changes within `src` can be approved by the engineering team
- Any changes to `build/konnect/<API_NAME>/dev.yaml` can be approved by the engineering team
- Any changes to `build/konnect/<API_NAME>/internal.yaml` (Internal Release) must be approved by any @Kong/team-aip member
- Any changes to `build/konnect/<API_NAME>/public.yaml` (GA) must be approved by @mheap (or delegated reviewer if he's OOO)
- When a change is detected in `build/konnect/<API_NAME>/public.yaml` a GitHub Action runs and updates the spec in Konnect dev portal. (PENDING)
