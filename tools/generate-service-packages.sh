#!/usr/bin/env bash

KONNECT_SPECS=$(find src/konnect/definitions -name openapi.yaml)
INTERNAL_SPECS=$(find src/internal/definitions -name openapi.yaml)

SPECS="$KONNECT_SPECS $INTERNAL_SPECS"

for SPEC in $SPECS; do
  SERVICE=$(dirname $SPEC | xargs dirname | xargs basename)

  # Skip admin api (Control Planes Config)
  # as it doesn't generate
  if [ $SERVICE == "control-planes-config" ]; then
    continue
  fi

  # Check if the version has changed since the last commit
  # if it hasn't, skip generating the service
  #CURRENT_VERSION=$(git show HEAD:$SPEC | yq .info.version -r)
  CURRENT_VERSION=$(cat $SPEC | yq .info.version -r)
  PREVIOUS_VERSION=$(git show HEAD^:$SPEC | yq .info.version -r)

  INTERNAL_LABEL=""

  if [[ $SPEC == *"/internal/"* ]]; then
    INTERNAL_LABEL=" (internal)"
  fi

  VERSION_FOLDER=$(dirname $SPEC | xargs basename)
  echo "-------------------------------------";
  echo "$SERVICE/$VERSION_FOLDER$INTERNAL_LABEL";
  echo "-------------------------------------";
  echo "Current: $CURRENT_VERSION"
  echo "Previous: $PREVIOUS_VERSION"

  if [ "x$CURRENT_VERSION" == "x$PREVIOUS_VERSION" ]; then
    echo "⏭️ Skipping"
    echo "";
    continue
  fi

  COMPUTED_SPEC=$(echo $SPEC | sed 's/src/computed/g' | sed 's/definitions\///g');

  MAJOR_VERSION=$(echo $CURRENT_VERSION | cut -d. -f1)

  if [[ $SPEC == *"/internal/"* ]]; then
    PACKAGE_NAME="internal-$SERVICE-v$MAJOR_VERSION-axios"
  else
    PACKAGE_NAME="service-$SERVICE-v$MAJOR_VERSION-axios"
  fi

  PACKAGE_NAME="@kong/$PACKAGE_NAME"

  STATUS=$(npx openapi-generator-cli generate \
  --generator-name typescript-axios \
  -i $COMPUTED_SPEC \
  -o ./packages/$SERVICE \
  --additional-properties modelPropertyNaming=original \
  --additional-properties useSingleRequestParameter=true \
  -t oas-generator-templates)

  if [ $? -eq 0 ]; then
    echo "✅ Generated $SERVICE as $PACKAGE_NAME"

    cd packages/$SERVICE
    cp ../../oas-generator-templates/package.json package.json
    npm pkg set name=$PACKAGE_NAME version=$CURRENT_VERSION description="Axios Client for the $SERVICE API"
    npm pkg set repository.url="https://github.com/Kong/platform-api/tree/main/$SPEC"
    cp ../../oas-generator-templates/tsconfig.json tsconfig.json
    npm install

    if [ "x$PUBLISH_TO_NPM" != "xtrue" ]; then
      echo "⏭️ Skipping publish"
      echo "";
      continue
    fi

    # If NODE_AUTH_TOKEN is not set, skip publishing
    if [[ -z "${NODE_AUTH_TOKEN}" ]]; then
      echo "❌ NODE_AUTH_TOKEN is not set. Skipping publish"
      echo "";
      continue
    fi

    npm publish
    cd -
    echo "✅ Published $SERVICE"
  else
    echo "❌  Failed to generate $SERVICE"
  fi

  echo "";
done
