.PHONY: *
all: speakeasy

speakeasy: check-speakeasy
	speakeasy generate sdk --lang terraform -o . -s ./openapi.yaml
	@git clean -fd examples > /dev/null
	@git checkout -- README.md
	@rm USAGE.md

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }

