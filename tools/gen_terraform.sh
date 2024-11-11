#!/usr/bin/env bash
mkdir -p build/complete/terraform

export KEEP_SPEAKEASY_ANNOTATIONS=1
export REQUIRE_SPEAKEASY_ENTITY_OPERATION=1
export SDK_NAME=terraform

# Get path to current file
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Process specs
$DIR/run.sh terraform; 

# Apply overlays to fix the spec
echo "Applying overlays"

function apply_overlay() {
  speakeasy overlay apply -o $DIR/../$1 -s $DIR/../build/complete/terraform/public.yaml > $DIR/../build/complete/terraform/public2.yaml
  mv $DIR/../build/complete/terraform/public2.yaml $DIR/../build/complete/terraform/public.yaml
}

apply_overlay overlays/konnect/control-planes-config.yaml
apply_overlay overlays/konnect/application-auth-strategies.yaml
apply_overlay overlays/konnect/api-products.yaml
apply_overlay overlays/konnect/identity.yaml
apply_overlay overlays/konnect/cloud-gateways.yaml
apply_overlay overlays/konnect/complex-filters.yaml

# Remove non-annotated fields
node $DIR/filter-operations.js

# Reformat the spec
npx openapi-format --sortFile .openapi-format-sort.json build/complete/terraform/public.yaml -o build/complete/terraform/public.yaml;

# Copy to TF provider folder
cp $DIR/../build/complete/terraform/public.yaml $DIR/../../terraform-provider-konnect/openapi.yaml
