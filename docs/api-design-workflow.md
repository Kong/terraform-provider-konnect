# API Design Workflow

Kong follows a design-first workflow when building new products. We use OpenAPI to design the API, get feedback from stakeholders and then use the specification to generate implementations, documentation and for contract testing.

There are two types of project: brand new APIs and new features in an existing API. Each project has its own needs and workflow, which you can find documented below.

## New API

New APIs are the easiest kind of API to build. When building a new API, follow these steps:

* Decide which API Product it will be exposed through. Possible values are `Konnect` or `Portal`.
* Create a new `openapi.yaml` file in [platform-api](https://github.com/Kong/platform-api/tree/main) for your API. e.g. `konnect/definitions/my-new-api/v1/openapi.yaml`
* Review the [common definitions](https://github.com/Kong/platform-api/tree/main/common/definitions) to see if there are any reusable schemas relevant to your project
* Start building your API spec. Make sure to annotate all operations with `x-internal:true` and `x-unstable: true` ([learn more](https://github.com/Kong/platform-api/blob/main/docs/openapi-annotations.md)). This allows us to generate a development SDK for your service without exposing it to all customers.
* Once your API spec has been reviewed and approved, merge to `main` to trigger downstream SDK builds.

## Existing API

If you're working on an existing API, the workflow is simpler. 

* Find the specification for the API that you're working on and start adding new API endpoints. 
* Remember to tag each endpoint with `x-internal: true` and `x-unstable: true` to make sure that it's only available to internal devs.
* Once your API spec has been reviewed and approved, merge to `main` to trigger downstream SDK builds.

## Migrating from Dev to Beta

> This is only applicable for public facing APIs. Internal APIs go directly from Dev to GA

You'll probably want to try out your new API with customers without marking it as Generally Available. To do this, remove the `x-internal: true` annotation from all endpoints that you want to expose.

The presence of the `x-unstable: true` annotation will produce a warning in the documentation for these endpoints stating that they're in beta and are subject to change.

## Migrating from Beta to GA

Once your API is built and tested, it's time to promote it from Beta to GA. To do this, you can remove the `x-unstable: true` annotation. At this point, your API will be accessible in one of two SDKs:

* If the endpoint has the `x-internal: true` annotation, it will be published in the `internal` SDK. These endpoints are technically publicly accessible, but we make no guarantees for consumers and may change them with zero notice
* If the endpoint does not have the `x-internal` or `x-unstable` annotations it is considered GA and is published in our public SDK. These endpoints must go through our [sunsetting an API](https://kong-aip.netlify.app/aip/3101/#:~:text=api/v2/projects-,Sunsetting%20an%20API,-Any%20GA%20API) process if later removed.