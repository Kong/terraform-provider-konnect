#!/usr/bin/env bash
set -e
shopt -s globstar

node ./tools/add-prefix-to-paths.js src/konnect/definitions/mesh-control-planes/v0/kong-mesh.yaml src/konnect/definitions/mesh-control-planes/v0/kong-mesh-prefixed.yaml "/mesh/control-planes/{id}" cpId "name: id
in: path
required: true
description: Id of the Konnect resource
example: bf138ba2-c9b1-4229-b268-04d9d8a6410b
schema:
  type: string
  format: uuid # validation happens via KinConfig()"

node ./tools/remove-common-definitions.js src/konnect/definitions/mesh-control-planes/v0/kong-mesh-prefixed.yaml src/konnect/definitions/mesh-control-planes/v0/kong-mesh-deduped.yaml

node ./tools/add-x-speakeasy-entity-operation.js src/konnect/definitions/mesh-control-planes/v0/kong-mesh-deduped.yaml > src/konnect/definitions/mesh-control-planes/v0/kong-mesh-annotated.yaml

node ./tools/remove-temporarily-unsupported-field.js src/konnect/definitions/mesh-control-planes/v0/kong-mesh-annotated.yaml "components.schemas.MeshHealthCheckItem.properties.spec.properties.to.items.properties.default.properties.http.properties.expectedStatuses.items.format" > src/konnect/definitions/mesh-control-planes/v0/kong-mesh-stripped.yaml

./node_modules/.bin/oas-toolkit merge src/konnect/definitions/mesh-control-planes/v0/kong-mesh-stripped.yaml src/konnect/definitions/mesh-control-planes/v0/mink-vcp-manager.yaml > src/konnect/definitions/mesh-control-planes/v0/openapi.yaml

rm src/konnect/definitions/mesh-control-planes/v0/kong-mesh-prefixed.yaml \
   src/konnect/definitions/mesh-control-planes/v0/kong-mesh-deduped.yaml \
   src/konnect/definitions/mesh-control-planes/v0/kong-mesh-annotated.yaml \
   src/konnect/definitions/mesh-control-planes/v0/kong-mesh-stripped.yaml \
