#!/usr/bin/env bash

# Merge the generated and hand maintained OAS files for Control Planes Config
FILES=(
  "src/konnect/definitions/control-planes-config/v2/generated.yaml"
  "src/konnect/definitions/control-planes-config/v2/konnect.yaml"
)

./node_modules/.bin/oas-toolkit check-conflicts ${FILES[@]} --ignorePrefix components.securitySchemes --ignorePrefix components.examples
./node_modules/.bin/oas-toolkit merge ${FILES[@]} > src/konnect/definitions/control-planes-config/v2/openapi.yaml

# Build all services
./tools/delete-build.sh $1
./tools/merge-oas.sh $1
node ./tools/generate-sdk-openapi-files.js $1
node ./tools/split-service-specs.js $1