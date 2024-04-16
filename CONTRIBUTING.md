# Contributing

This provider is generated from `openapi.yaml` using Speakeasy. To test a new version, use the following steps:

- Run `make speakeasy`
- Run `go mod tidy` to download all dependencies
- Run `go run main.go -debug`
- Take the `TF_REATTACH_PROVIDERS` variable and run `export TF_REATTACH_PROVIDERS=...` in your terminal
- Run `terraform apply` in the same directory as your manifests

## openapi.yaml

The `openapi.yaml` file in this repository is generated from Kong's platform-api repository (this repository is not OSS). Any changes made directly to `openapi.yaml` will be lost.

Please raise an issue if you have suggested changes to `openapi.yaml`.
