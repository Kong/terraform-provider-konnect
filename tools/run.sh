#!/usr/bin/env bash

# Merge the generated and hand maintained OAS files for Control Planes Config
FILES=(
  "src/konnect/definitions/control-planes-config/v2/konnect.yaml"
)

if [[ $1 == "terraform" ]]; then
  FILES+=("src/konnect/definitions/control-planes-config/v2/generated-with-plugins.yaml")
else
  FILES+=("src/konnect/definitions/control-planes-config/v2/generated.yaml")
fi

./node_modules/.bin/oas-toolkit check-conflicts ${FILES[@]} --ignorePrefix components.securitySchemes --ignorePrefix components.examples
./node_modules/.bin/oas-toolkit merge ${FILES[@]} > src/konnect/definitions/control-planes-config/v2/openapi.yaml

# Build all services
./tools/delete-build.sh $1
./tools/merge-oas.sh $1
node ./tools/generate-sdk-openapi-files.js $1
node ./tools/split-service-specs.js $1

# Make sure all specs are ordered deterministically
for i in $(find build -name '*.yaml'); do
  npx openapi-format --sortFile .openapi-format-sort.json $i -o $i
done