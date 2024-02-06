#!/usr/bin/env bash

./tools/merge-oas.sh $1
node ./tools/generate-sdk-openapi-files.js $1
node ./tools/split-service-specs.js $1