# Contributing

## Local Generation

This provider is generated from `openapi.yaml` using Speakeasy. To test a new version, use the following steps:

- Run `make speakeasy`
- Run `go mod tidy` to download all dependencies
- Run `go run main.go -debug`
- Take the `TF_REATTACH_PROVIDERS` variable and run `export TF_REATTACH_PROVIDERS=...` in your terminal
- Run `terraform apply` in the same directory as your manifests

Assuming that your changes are working, raise a PR to `platform-api` for review then see [contributing features](#contributing-features)

## Contributing features

- Create a new release branch (if needed) named `release/x.y.z`, where `x.y.z` is the new version
- Run `./tools/gen_terraform.sh` on the `main` branch of `platform-api` to update the `openapi.yaml` file in this repo
- Run `make speakeasy` to generate the provider (see [local generation](#local-generation))
- Test your changes to ensure the provider still works
- `git add openapi.yaml` and commit the changes as a single commit
- `git add` and commit the provider's `.go` files, plus the `.speakeasy` and `docs` directory
- Add examples for your new resources to `examples/resources` and commit them as a third commit
- Finally, update `CHANGELOG.md` and add it as a fourth commit
- Raise a PR against `release/x.y.z` for review

Submitting your changes as four commits is not a hard requirement, but it makes it easier to review the PR as four distinct sets of work:

1. OpenAPI changes
1. Provider changes
1. New examples
1. Changelog

Once the `release/x.y.z` branch is merged to `main`, a new release will automatically be created and published to the Terraform registry.

## Releasing the provider

This repository is automatically released when the `.speakeasy/gen.lock` file on `main` changes.

This usually happens when a `release/x.y.z` branch (see above) is merged in to `main`.

Before merging a release branch, update the release branch directly to set a release date in `CHANGELOG.md`. Once the release is complete, copy the changelog entry in to GitHub releases.

## openapi.yaml

The `openapi.yaml` file in this repository is generated from Kong's platform-api repository (this repository is not OSS). Any changes made directly to `openapi.yaml` will be lost.

Please raise an issue if you have suggested changes to `openapi.yaml`.

## Automated Testing

### Acceptance tests  
Acceptance tests are found in `/tests/resources` directory, and the terraform config files can be found in `/tests/resources/testdata` directory.
#### Running the tests
Export your access token as an environment variable, and then run the tests.  
```
export KONNECT_SPAT=
make acceptance
```  
#### Directory and naming conventions
- Each resource should have one test file, named `{resource}_test.go`, placed in `/tests/resources`, and corresponding config files in `/tests/resources/testdata/{testName}/` directory.   
- If there are subtests to cover different scenarios, create one sub directory per scenario under the directory for the test, and name both after the scenario. We follow the go convention of subtest names being the combination of the name of the top-level test and the sequence of names passed to Run, separated by slashes - pointing to the config files used in that test.  

#### Example
If we are adding tests for `gateway_control_plane resource` with different scenarios for each type, we would:  
1. Add `gateway_control_plane_test.go` in `/tests/resources`
2. Create a test `TestGatewayControlPlane` in `gateway_control_plane_test.go`
3. Create a directory `TestGatewayControlPlane` in `/tests/resources/testdata` to add the config.  
4. Add subtests in `gateway_control_plane_test.go` for each scenario
5. Add config files for each scenario in `/tests/resources/testdata/TestGatewayControlPlane/{scenarioName}/`.

For more information, refer https://developer.hashicorp.com/terraform/plugin/testing/acceptance-tests

### End to end tests  
The `make test` target runs some end to end tests against an org that @mheap owns. The tests need some refactoring (system account team IDs need to be dynamic) before anyone else can run them.
