# OpenAPI Annotations

Kong uses custom annotations in OAS files to communicate an API's visibility, stability and where in the development lifecycle it is. These annotations are prefixed with `x-` to make it easy to understand what is Kong specific and what is part of the spec.

We currently support three annotations:

- `x-publish`: Controls if the specification is published to our public API portal. Should be added inside the `info` block. Possible values: `true` (default), `false`
- `x-internal`: Controls which SDK the endpoint is generated in to. May be added in the `info` block or on a per-operation basis. Per-operation takes precedence. Possible values: `true`, `false` (default)
- `x-unstable`: Is the API contract finalised yet? Possible values: `true`, `false` (default)

