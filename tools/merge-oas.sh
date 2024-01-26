#!/usr/bin/env bash
shopt -s globstar
set -e

PRODUCTS=(konnect portal)
for PRODUCT in "${PRODUCTS[@]}"; do
echo "############################################"
echo "Processing: $PRODUCT"
echo "############################################"
  mkdir -p computed/$PRODUCT

  FILES=$(node ./tools/get-product-files.js $PRODUCT --plain);
  COMPUTED_FILES=$(echo $FILES | sed 's/src/computed/g' | sed 's/definitions\///g');

  ./node_modules/.bin/oas-toolkit check-conflicts $FILES src/common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples
  # Bundle here to make sure `computed` comes from `src` after we've linted
  node ./tools/bundle.js $PRODUCT
  node ./tools/process-computed.js $PRODUCT
  echo "Creating merged $PRODUCT/complete.yaml";
  ./node_modules/.bin/oas-toolkit merge $COMPUTED_FILES src/common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples > computed/$PRODUCT/complete.yaml
done

# Finally, bundle all internal APIs. They do not need merging
# so we do it outside the loop
echo "############################################"
echo "Processing: internal"
echo "############################################"
node ./tools/bundle.js internal
node ./tools/process-computed.js internal