# platform-api

This repository contains OpenAPI specifications for Kong's APIs. This includes Konnect Platform, Konnect Portal and on-prem APIs.

Specifications in this repository are the source of truth. Any changes to an API should be raised as a PR for discussion before changes are made to the underlying API. If you need a copy of the API specification in your service repo (e.g. for contract testing) it should be a copy taken from this repository (automation pending, see [API-6](https://konghq.atlassian.net/browse/API-6))

GitHub Actions are used to perform the following tasks:

- Lint APIs to ensure they meet our [AIP definitions](https://kong-aip.netlify.app/)
- Generate SDKs for use internally (and externally eventually)
