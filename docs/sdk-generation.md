# SDK Generation

SDKs are generated using [openapi-generator](https://github.com/OpenAPITools/openapi-generator) via GitHub Actions.

We currently generate JavaScript SDKs and publish them to NPM.

The generation process looks like the following:

- Combine the `openapi.yaml` file from all APIs in to a single `openapi.yaml` file
- Split the combined `openapi.yaml` in to `public.yaml`, `internal.yaml` and `dev.yaml`
- Generate one SDK for each of those files
- Publish all SDKs

## Splitting `openapi.yaml`

The combined file is split based on the following psuedocode:

```js
// Operations go in to the public SDK by default
let sdk = "public";

// If the endpoint is tagged with `x-internal` then it
// goes in to the internal SDK
if (operation["x-internal"]) {
  sdk = "internal";

  // Unless it's also tagged with `x-unstable`, in which case
  // it is added to the dev SDK
  if (operation["x-unstable"]) {
    sdk = "dev";
  }
}
```

You may notice that `x-unstable` endpoints can end up in the public SDK. This is by design, and is used to run public betas. The description field for these items will have a large warning added by a post-processing step
