#!/usr/bin/env bash
mkdir -p dist

DOCS=$1

if [[ -z "$DOCS" ]]; then
  DOCS=$(find konnect/definitions -name 'openapi.yaml')
fi

rm -f dist/index.html
echo "<ul>" >> dist/index.html
for i in $DOCS; do
  API=$(echo $i | sed 's#definitions/##')
  mkdir -p dist/$(echo $API | sed 's#/openapi.yaml##')
  ./node_modules/.bin/redoc-cli build $i -o dist/$API.html
  echo "<li><a href='/$API.html'>$API</a></li>" >> dist/index.html
done
echo "</ul>" >> dist/index.html