#!/usr/bin/env bash
shopt -s globstar
set -e

PRODUCTS=(konnect portal)
for PRODUCT in "${PRODUCTS[@]}"; do
echo "############################################"
echo "Processing: $PRODUCT"
echo "############################################"
  mkdir -p output/$PRODUCT

  FILES=$(node ./tools/get-product-files.js $PRODUCT --plain);
  COMPUTED_FILES=$(echo $FILES | sed 's/src/computed/g');

  ./node_modules/.bin/oas-toolkit check-conflicts $FILES common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples
  # Bundle here to make sure `computed` comes from `src` after we've linted
  node ./tools/bundle.js $PRODUCT
  node ./tools/process-computed.js $PRODUCT
  ./node_modules/.bin/oas-toolkit merge $COMPUTED_FILES common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples > output/$PRODUCT/openapi.yaml
done

# Finally, bundle all internal APIs. They do not need merging
# so we do it outside the loop
echo "############################################"
echo "Processing: internal"
echo "############################################"
node ./tools/bundle.js internal
node ./tools/process-computed.js internal