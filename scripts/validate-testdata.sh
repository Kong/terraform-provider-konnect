#!/usr/bin/env bash
#
# Type-checks every acceptance-test fixture against the provider built from
# this tree, without touching Konnect.
#
# The acceptance suite needs credentials, creates real resources and takes
# tens of minutes, so a typo or a stale attribute in a testdata .tf file is
# otherwise only discovered deep inside that run - or, on a fork PR where the
# credentialed jobs cannot run at all, not discovered until merge. This does
# the same schema check in seconds with no secrets.
#
# Usage:
#   scripts/validate-testdata.sh [dir ...]     # default: tests/resources/testdata
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

for cmd in terraform go; do
  command -v "$cmd" >/dev/null 2>&1 || { echo "error: $cmd is required" >&2; exit 1; }
done

ROOTS=("$@")
if [ ${#ROOTS[@]} -eq 0 ]; then
  ROOTS=("tests/resources/testdata")
fi

VERSION="999.99.9"
OS="$(go env GOOS)"
ARCH="$(go env GOARCH)"

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

MIRROR="$WORK/mirror"
PLUGIN_DIR="$MIRROR/registry.terraform.io/kong/konnect/$VERSION/${OS}_${ARCH}"
mkdir -p "$PLUGIN_DIR"

echo "==> building provider"
go build -o "$PLUGIN_DIR/terraform-provider-konnect_v$VERSION" .

# Resolve kong/konnect from the local build only; never reach the registry.
cat > "$WORK/terraformrc" <<EOF
provider_installation {
  filesystem_mirror {
    path    = "$MIRROR"
    include = ["registry.terraform.io/kong/konnect"]
  }
  direct {
    exclude = ["registry.terraform.io/kong/konnect"]
  }
}
EOF
export TF_CLI_CONFIG_FILE="$WORK/terraformrc"
export TF_IN_AUTOMATION=1

# Fixtures carry no terraform{} or provider{} block - the acceptance harness
# supplies those at runtime - so supply equivalents here.
#
# required_providers is split out because a child module does not inherit the
# root's source mapping: without its own block a module using konnect_*
# resources resolves to hashicorp/konnect and fails to install. Provider
# *configuration*, by contrast, is inherited from the root.
required_providers() {
  cat <<EOF
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "$VERSION"
    }
  }
}
EOF
}

harness() {
  required_providers
  cat <<EOF

# Never used: validate does not contact the API.
provider "konnect" {
  personal_access_token = "spat_validate_only"
  server_url            = "https://us.api.konghq.com"
}
EOF
}

# Every directory that directly contains .tf files is a standalone config:
# ConfigDirectory copies exactly one directory per test step.
#
# Not mapfile: that is bash 4+, and macOS still ships bash 3.2.
DIRS=()
while IFS= read -r d; do
  DIRS+=("$d")
done < <(find "${ROOTS[@]}" -name '*.tf' -exec dirname {} \; | sort -u)

if [ ${#DIRS[@]} -eq 0 ]; then
  echo "error: no .tf fixtures found under ${ROOTS[*]}" >&2
  exit 1
fi

# Fixtures under a fail_* directory back an ExpectError test: they are invalid
# on purpose. They are checked one at a time below, and must stay invalid - a
# schema change that quietly makes one valid should break here rather than
# silently turning the test that depends on it into a no-op.
BATCH=()
EXPECT_INVALID=()
for dir in "${DIRS[@]}"; do
  case "$(basename "$dir")" in
    fail_*) EXPECT_INVALID+=("$dir") ;;
    *)      BATCH+=("$dir") ;;
  esac
done

# validate_one stages a single fixture as a root module and validates it.
# Prints terraform's output on failure. Returns non-zero if invalid.
validate_one() {
  local dir="$1" stage
  stage="$WORK/one"
  rm -rf "$stage"
  mkdir -p "$stage"
  cp "$dir"/*.tf "$stage/"
  harness > "$stage/zz_validate_harness.tf"
  ( cd "$stage" && terraform init -input=false -backend=false 2>&1 && terraform validate 2>&1 )
}

failed=0

# Spawning the provider costs ~200ms - it is a 78MB binary that ships a 1.8MB
# schema over gRPC - and that dwarfs the actual type-checking. Validating each
# fixture separately pays it once per directory. Nesting them all as modules
# under one root pays it once in total.
#
# Module nesting namespaces resource addresses, so fixtures that reuse names do
# not collide. On failure we fall back to per-directory validation, because a
# batched error is reported against a module path rather than a file.
if [ ${#BATCH[@]} -gt 0 ]; then
  echo "==> validating ${#BATCH[@]} fixture directories"
  BATCH_ROOT="$WORK/batch"
  mkdir -p "$BATCH_ROOT/modules"
  harness > "$BATCH_ROOT/main.tf"

  for dir in "${BATCH[@]}"; do
    # Module labels allow letters, digits, underscore and dash only.
    name="$(printf '%s' "$dir" | tr -c 'A-Za-z0-9_-' '_')"
    mkdir -p "$BATCH_ROOT/modules/$name"
    cp "$dir"/*.tf "$BATCH_ROOT/modules/$name/"
    required_providers > "$BATCH_ROOT/modules/$name/zz_validate_harness.tf"
    cat >> "$BATCH_ROOT/main.tf" <<EOF

module "$name" {
  source = "./modules/$name"
}
EOF
  done

  if ! ( cd "$BATCH_ROOT" && terraform init -input=false -backend=false >/dev/null 2>&1 && terraform validate >/dev/null 2>&1 ); then
    echo "    batch validation failed; re-checking each directory for precise errors"
    for dir in "${BATCH[@]}"; do
      if out=$(validate_one "$dir"); then
        continue
      fi
      failed=$((failed + 1))
      echo
      echo "FAIL $dir"
      echo "$out" | sed 's/^/    /'
    done
  fi
fi

for dir in "${EXPECT_INVALID[@]}"; do
  if validate_one "$dir" >/dev/null 2>&1; then
    failed=$((failed + 1))
    echo
    echo "FAIL $dir"
    echo "    expected this fixture to be rejected (it backs an ExpectError test), but it validated cleanly"
  fi
done

total=$(( ${#BATCH[@]} + ${#EXPECT_INVALID[@]} ))
echo
if [ "$failed" -ne 0 ]; then
  echo "==> $failed of $total fixture directories failed validation"
  exit 1
fi
echo "==> all $total fixture directories are valid"
