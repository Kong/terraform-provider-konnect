#!/usr/bin/env bash
set -e

MESH_SPEC="./src/konnect/definitions/mesh-control-planes/v0/openapi.yaml"
MESH=$(npx oas-toolkit rewrite-path --oldPath '^/mesh/' --newPath '/' $MESH_SPEC)
echo "$MESH" > $MESH_SPEC