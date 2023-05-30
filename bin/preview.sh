#!/usr/bin/env bash
mkdir -p dist


# get script directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
$DIR/../tools/merge-oas.sh

function build_doc {
  
  # get paths from first argument
  OAS_PATHS=$1
  FILE_NAME_SUFFIX=$2
  LIST_TITLE=$3
  
  DOCS=$(find $OAS_PATHS -name "$FILE_NAME_SUFFIX")
  echo "<h2>$LIST_TITLE</h2>" >> dist/index.html
  
  echo "<ul>" >> dist/index.html
  for doc_path in $DOCS; do

    # remove the definitions and computed from the path

    apiPath=$(echo $doc_path | sed 's#definitions/##' | sed 's#computed/##' | sed 's#output/##'| sed 's#/openapi.yaml##')
    targetDir="dist/$apiPath"
    # remove the openapi.yaml from the path and create the directory
    mkdir -p $targetDir

    # use redocly cli to build the html
    npx @redocly/cli build-docs $doc_path -o $targetDir/index.html

    # write a link to the index 
    echo "<li><a href='/$apiPath'>$apiPath</a></li>" >> dist/index.html
    
  
  done
  echo "</ul>" >> dist/index.html
}

DOCS=$1
portal_paths="portal/definitions/**/computed"
konnect_paths="konnect/definitions/**/*/computed"
konnect_output="output/konnect/*.yaml"
portal_output="output/portal/*.yaml"
internal_paths="internal/definitions/**/**/computed"

rm -f dist/index.html

# if DOCS is not empty, then we are building a single doc
if [[ -z "$DOCS" ]]; then
  # Build internal docs
  build_doc "$internal_paths" "*.yaml" "internal computed"

  # Build computed docs
  build_doc "$portal_paths" "openapi.yaml" "portal computed" 
  build_doc "$konnect_paths" "openapi.yaml" "konnect computed" 
  
  # build bundled docs
  build_doc "$portal_output" "*.yaml" "portal bundled" 
  build_doc "$konnect_output" "*.yaml" "konnnect bundled" 

fi
