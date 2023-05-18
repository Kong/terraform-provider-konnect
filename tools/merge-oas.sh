#!/usr/bin/env bash
mkdir -p output/{konnect,portal}
./node_modules/.bin/oas-toolkit merge konnect/definitions/*/*/src/openapi.yaml common/konnect.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples > output/konnect/complete.yaml
./node_modules/.bin/oas-toolkit merge portal/definitions/*/src/openapi.yaml common/portal.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples > output/portal/complete.yaml

sed -i 's#^\(\s*\).*common#\1../../common#' output/konnect/complete.yaml
sed -i 's#^\(\s*\).*common#\1../../common#' output/portal/complete.yaml