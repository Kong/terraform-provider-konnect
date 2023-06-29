#!/usr/bin/env bash
shopt -s globstar

PRODUCTS=(konnect portal)
for PRODUCT in "${PRODUCTS[@]}"; do
echo "############################################"
echo "Processing: $PRODUCT"
echo "############################################"
  mkdir -p output/$PRODUCT
  ./node_modules/.bin/oas-toolkit check-conflicts $PRODUCT/definitions/**/src/openapi.yaml common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples
  # Bundle here to make sure `computed` comes from `src` after we've linted
  node ./tools/bundle.js $PRODUCT
  node ./tools/process-computed.js $PRODUCT
  ./node_modules/.bin/oas-toolkit merge $PRODUCT/definitions/**/computed/openapi.yaml common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples > output/$PRODUCT/openapi.yaml
done