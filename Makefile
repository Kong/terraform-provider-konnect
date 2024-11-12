.PHONY: *

all: speakeasy

speakeasy: check-speakeasy
	speakeasy run --skip-versioning --output console
	@go mod tidy
	@go generate .
	@git clean -fd examples docs > /dev/null
	@git checkout -- README.md examples/README.md
	@rm USAGE.md

FILES=$(shell find internal/provider -type f | grep data_source | grep -v portallist | grep -v cloudgatewayprovideraccountlist)
remove-data-sources:
	@if [ -n "$(FILES)" ]; then rm $(FILES); fi
	@rm -r examples/data-sources

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }

OS=$(shell uname | tr "[:upper:]" "[:lower:]")
ARCH=$(shell uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/')
test:
	@cd tests/e2e; rm -rf .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup .terraform local-plugins
	mkdir -p tests/e2e/local-plugins/registry.terraform.io/kong/konnect/999.99.9/$(OS)_$(ARCH)
	go mod tidy
	go build -o tests/e2e/local-plugins/registry.terraform.io/kong/konnect/999.99.9/$(OS)_$(ARCH)/terraform-provider-konnect_v999.99.9
	@cd tests/e2e; terraform providers mirror ./local-plugins || true
	@cd tests/e2e; ls -R local-plugins; terraform init -plugin-dir ./local-plugins; terraform apply -auto-approve; terraform destroy -auto-approve


test-cleanup:
	@cd tests/e2e; rm -rf local-plugins .terraform .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup
