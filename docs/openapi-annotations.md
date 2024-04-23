# OpenAPI Annotations

Kong uses custom annotations in OAS files to communicate an API's visibility, stability and where in the development lifecycle it is. These annotations are prefixed with `x-` to make it easy to understand what is Kong specific and what is part of the spec.

## Publishing Specifications

We use a single `openapi.yaml` file for public, internal and in-development work. Visibility of the available operations is controlled using two annotations, `x-internal` and `x-unstable`.

* If neither annotation is applied, the functionality is considered GA and is rendered to `public.yaml`
* If `x-unstable` is applied, the functionality is labelled as beta and is rendered to `public.yaml`
* If `x-internal` is applied, the functionality is considered GA, but not for public consumption and is rendered to `internal.yaml`
* If `x-internal` and `x-unstable` are applied, it's considered in development and is rendered to `dev.yaml`

These annotations can be applied to either endpoints (alongside `operationId`) or to schemas in `components.schemas`. If applied to a schema, the schema and any references to that schema will be removed from the rendered file. This allows you to add new fields to existing endpoints without publishing them publicly whilst the functionality is in development.

If you want to add a new property to an endpoint that uses a shared schema, you can add an `x-property-annotations` entry at the same level as `properties` within the schema:

```yaml
MyResource:
  type: object
  x-property-annotations:
    my_property: [x-unstable, x-internal]
  properties:
    id:
      type: string
    something:
      type: string
    my_property: # This will be removed
      type: string
```