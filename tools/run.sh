#!/usr/bin/env bash

./tools/merge-oas.sh
npm run generate-sdk-oas
node ./tools/split-service-specs.js