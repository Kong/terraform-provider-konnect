#!/usr/bin/env bash
mkdir -p build/complete/go-sdk

export KEEP_SPEAKEASY_ANNOTATIONS=1
export REQUIRE_SPEAKEASY_ENTITY_OPERATION=0
export SDK_NAME=go-sdk

# Get path to current file
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Process specs
$DIR/run.sh go-sdk; 

# Apply overlays to fix the spec
echo "Applying overlays"

function apply_overlay() {
  set -x
  speakeasy overlay apply -o $DIR/../$1 -s $DIR/../build/complete/go-sdk/public.yaml > $DIR/../build/complete/go-sdk/public2.yaml
  mv $DIR/../build/complete/go-sdk/public2.yaml $DIR/../build/complete/go-sdk/public.yaml
}

# NOTE(pmalek): Remove defaults from SDK types. Those defaults cause the generated
# types to not have "omitempty" struct tags which then causes the generated CRDs
# to require those fields to be set.
apply_overlay "overlays/go-sdk/remove-route-defaults.yaml"
apply_overlay "overlays/go-sdk/remove-service-defaults.yaml"
# NOTE(pmalek): Remove client_certificate reference until that gets its own strongly typed struct.
apply_overlay "overlays/go-sdk/remove-client-certificate.yaml"

# Remove non-annotated fields
node $DIR/filter-operations.js

# Copy to sdk-konnect-go folder
SDK_KONNECT_GO_DIR="${DIR}/../../sdk-konnect-go"
if [[ ! -d "${SDK_KONNECT_GO_DIR}" ]]
then
  echo "${SDK_KONNECT_GO_DIR} directory not found, skipping copy"
else
  cp $DIR/../build/complete/go-sdk/public.yaml "${SDK_KONNECT_GO_DIR}/openapi.yaml"
fi
