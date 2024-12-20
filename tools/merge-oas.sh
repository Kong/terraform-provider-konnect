#!/usr/bin/env bash
shopt -s globstar
set -e

BUNDLE_HAS_RUN=0
if [[ ! -z "$1" ]]; then
  PRODUCTS=($1)
else
  PRODUCTS=(konnect portal internal)

  # If we're not building a single product build all computed files
  # so that we can build per-service SDKs
  node ./tools/bundle.js
  BUNDLE_HAS_RUN=1
fi

for PRODUCT in "${PRODUCTS[@]}"; do
  if [ "$PRODUCT" == "internal" ]; then
    continue
  fi

  echo "############################################"
  echo "Processing: $PRODUCT"
  echo "############################################"
  mkdir -p computed/$PRODUCT

  FILES=$(node ./tools/get-product-files.js $PRODUCT --plain);
  COMPUTED_FILES=$(echo $FILES | sed 's/src/computed/g' | sed 's/definitions\///g');

  ./node_modules/.bin/oas-toolkit check-conflicts $FILES src/common/$PRODUCT.yaml --ignorePrefix components.securitySchemes --ignorePrefix components.examples

  if [[ $BUNDLE_HAS_RUN == 0 ]]; then
    node ./tools/bundle.js $PRODUCT
  fi

  # Move any global.api.konghq.com server blocks to be per-operation
  node ./tools/move-global-server-block.js $PRODUCT

  # Flatten allOf references where needed
  node ./tools/flatten-allof.js $PRODUCT

  # Add beta warnings
  node ./tools/process-computed.js $PRODUCT

  # Merge everything together
  echo "Creating merged $PRODUCT/complete.yaml";
  ./node_modules/.bin/oas-toolkit merge --move-path-to-operation $COMPUTED_FILES src/common/$PRODUCT.yaml > computed/$PRODUCT/complete.yaml

  # Remove per-endpoint servers from per-service specs
  node ./tools/remove-global-server-block.js $PRODUCT
done

# Finally, bundle all internal APIs. They do not need merging
# so we do it outside the loop
if [[ " ${PRODUCTS[@]} " =~ " internal " ]]; then
echo "############################################"
echo "Processing: internal"
echo "############################################"
node ./tools/bundle.js internal
node ./tools/process-computed.js internal
fi