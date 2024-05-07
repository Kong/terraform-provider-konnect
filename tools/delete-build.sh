#!/usr/bin/env bash

# Always clear computed
rm -rf computed

if [[ ! -z "$1" ]]; then
    PRODUCTS=($1)
    for PRODUCT in "${PRODUCTS[@]}"; do
       rm -rf "build/services/$PRODUCT"
       rm -rf "build/complete/$PRODUCT"
    done
else
    rm -rf build
    mkdir build
fi

