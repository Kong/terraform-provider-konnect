#!/usr/bin/env bash

set -euo pipefail
mkdir -p dist

# get script directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
$DIR/../tools/run.sh

function add_h1 {
  echo "<h1>$1</h1>" >> dist/index.html
}

function add_h2 {
  echo "<h2>$1</h2>" >> dist/index.html
}

function build_doc {
  
  # get paths from first argument
  OAS_PATHS=$1
  FILE_NAME_SUFFIX=$2
  ADD_SEPARATOR=$3

  LAST_SERVICE_NAME="UNSET_VALUE"
  
  DOCS=$(ls --color=never $OAS_PATHS/$FILE_NAME_SUFFIX)

  echo "<ul>" >> dist/index.html

  for doc_path in $DOCS; do

    # continue if doc_path contains api-docs
    if [[ $doc_path == *"api-docs"* ]]; then
      continue
    fi

    # Header
    if [[ $ADD_SEPARATOR == "true" ]]; then
      # Strip the filename from the path
      serviceName=$(dirname $doc_path) 

      # If $serviceName contains /v2/ or /v3/ at the end, remove it
      serviceName=$(echo $serviceName | sed -E 's#/v[0-9]+$##')

      # Extract the service name (last segment)
      serviceName=$(basename $serviceName)

      if [[ $serviceName != $LAST_SERVICE_NAME ]]; then
        echo "</ul>" >> dist/index.html
        echo "<h3>$serviceName</h3>" >> dist/index.html
        echo "<ul>" >> dist/index.html
        LAST_SERVICE_NAME=$serviceName
      fi
    fi
    # remove the definitions and computed from the path

    apiPath=$(echo $doc_path | sed 's#computed/##' | sed 's#build/##')
    targetDir="dist/$apiPath"
    # remove the openapi.yaml from the path and create the directory
    mkdir -p $targetDir

    # use redocly cli to build the html
    npx @redocly/cli build-docs $doc_path -o $targetDir/index.html

    # write a link to the index 
    echo "<li><a href='/$apiPath/index.html'>$apiPath</a></li>" >> dist/index.html
    
  
  done
  echo "</ul>" >> dist/index.html
}

function start_details {
  echo "<details><summary><h2 style='display: inline;'>$1</h2></summary>" >> dist/index.html
}

function end_details {
  echo "</details>" >> dist/index.html
}

rm -f dist/index.html

add_h1 "Konnect"
start_details "Service Combined"
build_doc "computed/konnect/*/*" "openapi.yaml" "false"
end_details

start_details "Service By Visibility"
build_doc "build/services/konnect/*/*" "*.yaml" "true"
end_details

start_details "Public Facing"
build_doc "build/complete/konnect" "*.yaml" "false"
end_details

add_h1 "Portal"
start_details "Service Combined"
build_doc "computed/portal/*/*" "openapi.yaml" "false"
end_details

start_details "Service By Visibility"
build_doc "build/services/portal/*/*" "*.yaml" "true"
end_details

start_details "Public Facing"
build_doc "build/complete/portal" "*.yaml" "false"
end_details

add_h1 "Internal"
start_details "Service Combined"
build_doc "computed/internal/*/*" "openapi.yaml" "true"
end_details