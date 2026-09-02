#!/usr/bin/env bash
#
# End-to-end check for the v0 -> v1 `workspace` state migration.
#
# Applies a fixture with the provider built from a baseline ref (default:
# main, which predates `workspace`), then swaps in the provider built from the
# working tree and asserts:
#
#   1. `terraform plan` is empty        - the upgrade introduces no drift
#   2. `terraform apply` succeeds
#   3. every workspace-capable resource in state has workspace = "default"
#      and schema_version = 1
#   4. a second plan is still empty     - the upgrade is idempotent
#
# This drives the real Terraform CLI against two real provider binaries, so it
# exercises the same code path a user hits on `terraform init -upgrade`.
# For fast, credential-free coverage of every upgrader, see tests/migration.
#
# Usage:
#   scripts/workspace-migration-test.sh [options]
#
#   --baseline-ref REF   git ref to build the "before" provider from (default: main)
#   --fixtures DIR       directory to take .tf files from
#                        (default: tests/migration/fixtures, which covers ~20
#                        workspace-carrying resource types under one control
#                        plane). The script supplies its own terraform/provider
#                        block, so fixtures must not declare one.
#   --files LIST         comma-separated basenames to apply, or "all"
#                        (default: all)
#   --work-dir DIR       where to run terraform (default: a temp dir)
#   --keep               skip `terraform destroy` and leave the work dir behind
#
# Requires: terraform, go, git, jq, and Konnect credentials in either
# KONNECT_SPAT or KONNECT_TOKEN.

set -euo pipefail

BASELINE_REF="main"
FIXTURES=""
FILES="all"
WORK_DIR=""
KEEP=0

# Distinct versions so `terraform init -upgrade` performs a genuine provider
# upgrade rather than silently reusing a cached binary of the same version.
BASELINE_VERSION="999.99.8"
CANDIDATE_VERSION="999.99.9"
SERVER_URL="${E2E_SERVER_URL:-https://us.api.konghq.com}"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --baseline-ref) BASELINE_REF="$2"; shift 2 ;;
    --fixtures)     FIXTURES="$2"; shift 2 ;;
    --files)        FILES="$2"; shift 2 ;;
    --work-dir)     WORK_DIR="$2"; shift 2 ;;
    --keep)         KEEP=1; shift ;;
    -h|--help)      sed -n '2,30p' "$0" | sed 's/^# \{0,1\}//'; exit 0 ;;
    *) echo "unknown option: $1" >&2; exit 64 ;;
  esac
done

REPO_ROOT="$(git rev-parse --show-toplevel)"
FIXTURES="${FIXTURES:-$REPO_ROOT/tests/migration/fixtures}"
OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/')"

step() { printf '\n\033[1;34m==> %s\033[0m\n' "$*"; }
info() { printf '    %s\n' "$*"; }
fail() { printf '\n\033[1;31mFAIL: %s\033[0m\n' "$*" >&2; exit 1; }
pass() { printf '\033[1;32m  ok\033[0m %s\n' "$*"; }

for tool in terraform go git jq; do
  command -v "$tool" >/dev/null 2>&1 || fail "$tool is not installed"
done
[[ -n "${KONNECT_SPAT:-}${KONNECT_TOKEN:-}" ]] \
  || fail "set KONNECT_SPAT or KONNECT_TOKEN; this test creates real Konnect resources"
[[ -d "$FIXTURES" ]] || fail "fixtures directory not found: $FIXTURES"

SCRATCH="$(mktemp -d)"
WORK_DIR="${WORK_DIR:-$SCRATCH/work}"
MIRROR="$SCRATCH/mirror"
BASELINE_TREE="$SCRATCH/baseline"

# Resolve only kong/konnect from our mirror. `-plugin-dir` would force *every*
# provider through it, which breaks fixtures that also use e.g. hashicorp/http.
mkdir -p "$MIRROR"
cat > "$SCRATCH/terraformrc" <<EOF
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
export TF_CLI_CONFIG_FILE="$SCRATCH/terraformrc"

cleanup() {
  local rc=$?
  # Only worth destroying if something actually got into state - a failure
  # during init or build leaves nothing behind.
  if [[ $KEEP -eq 0 && -s "$WORK_DIR/terraform.tfstate" ]] \
    && jq -e '(.resources // []) | length > 0' "$WORK_DIR/terraform.tfstate" >/dev/null 2>&1; then
    step "Destroying fixture resources"
    terraform -chdir="$WORK_DIR" destroy -auto-approve -input=false >/dev/null 2>&1 \
      || echo "warning: destroy failed; check for leaked Konnect resources" >&2
  fi
  # Always drop the worktree - a stale registration blocks the next run.
  git -C "$REPO_ROOT" worktree remove --force "$BASELINE_TREE" >/dev/null 2>&1 || true
  if [[ $KEEP -eq 1 ]]; then
    echo "work dir kept at $WORK_DIR"
  else
    rm -rf "$SCRATCH"
  fi
  exit $rc
}
trap cleanup EXIT

# build_provider <source-dir> <version>
# Installs a provider binary into the filesystem mirror under the layout
# Terraform's filesystem_mirror expects.
build_provider() {
  local src="$1" version="$2"
  local dest="$MIRROR/registry.terraform.io/kong/konnect/$version/${OS}_${ARCH}"
  mkdir -p "$dest"
  (cd "$src" && go build -o "$dest/terraform-provider-konnect_v$version" .)
}

write_provider_config() {
  local version="$1"
  cat > "$WORK_DIR/_migration_provider.tf" <<EOF
# Generated by scripts/workspace-migration-test.sh - do not edit.
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "$version"
    }
  }
}

provider "konnect" {
  server_url = "$SERVER_URL"
}
EOF
}

# Emits one `<address>\t<type>\t<schema_version>\t<workspace-or-@absent>` line
# per konnect-managed resource in state, walking nested modules too.
state_rows() {
  terraform -chdir="$WORK_DIR" show -json | jq -r '
    (.values.root_module // {})
    | [recurse(.child_modules[]?) | .resources[]?]
    | .[]
    | select(.provider_name | test("kong/konnect$"))
    | [ .address,
        .type,
        (.schema_version | tostring),
        (if (.values | has("workspace")) and (.values.workspace != null)
         then .values.workspace else "@absent" end)
      ]
    | @tsv'
}

step "Preparing work dir: $WORK_DIR"
mkdir -p "$WORK_DIR"
copied=0
for f in "$FIXTURES"/*.tf; do
  [[ -e "$f" ]] || continue
  base="$(basename "$f")"
  if [[ "$FILES" != "all" && ",$FILES," != *",$base,"* ]]; then
    continue
  fi
  cp "$f" "$WORK_DIR/"
  info "using $base"
  copied=$((copied + 1))
done
[[ $copied -gt 0 ]] || fail "no .tf files matched --files '$FILES' in $FIXTURES"

if grep -q "required_providers" "$WORK_DIR"/*.tf 2>/dev/null; then
  fail "fixture declares required_providers, which conflicts with the generated
provider block; exclude that file via --files"
fi

step "Building baseline provider from '$BASELINE_REF' (v$BASELINE_VERSION)"
# A CI checkout often has no local branch, only the remote-tracking ref.
if ! git -C "$REPO_ROOT" rev-parse --verify --quiet "$BASELINE_REF^{commit}" >/dev/null; then
  if git -C "$REPO_ROOT" rev-parse --verify --quiet "origin/$BASELINE_REF^{commit}" >/dev/null; then
    info "'$BASELINE_REF' not found locally, using 'origin/$BASELINE_REF'"
    BASELINE_REF="origin/$BASELINE_REF"
  else
    fail "baseline ref '$BASELINE_REF' does not resolve (in CI, check out with fetch-depth: 0)"
  fi
fi
git -C "$REPO_ROOT" worktree add --detach "$BASELINE_TREE" "$BASELINE_REF" >/dev/null
build_provider "$BASELINE_TREE" "$BASELINE_VERSION"
info "$(git -C "$BASELINE_TREE" log --oneline -1)"

step "Building candidate provider from working tree (v$CANDIDATE_VERSION)"
build_provider "$REPO_ROOT" "$CANDIDATE_VERSION"

step "Phase 1: apply fixture with the baseline provider"
write_provider_config "$BASELINE_VERSION"
terraform -chdir="$WORK_DIR" init -input=false >/dev/null
terraform -chdir="$WORK_DIR" apply -auto-approve -input=false >/dev/null
info "applied"

baseline_rows="$(state_rows)"
[[ -n "$baseline_rows" ]] || fail "baseline apply produced no konnect resources in state"
info "$(wc -l <<<"$baseline_rows" | tr -d ' ') konnect resource(s) in state"

# Guards against the baseline ref having been merged forward: if it already
# knows about `workspace` there is no migration left to test.
if leaked="$(awk -F'\t' '$4 != "@absent"' <<<"$baseline_rows")" && [[ -n "$leaked" ]]; then
  fail "baseline state already contains workspace values - is '$BASELINE_REF' really pre-workspace?
$leaked"
fi
pass "baseline state has no workspace attribute"

step "Phase 2: swap in the candidate provider and plan"
write_provider_config "$CANDIDATE_VERSION"
# The version constraint changed, so the recorded lock entry no longer applies.
rm -f "$WORK_DIR/.terraform.lock.hcl"
terraform -chdir="$WORK_DIR" init -input=false -upgrade >/dev/null

set +e
terraform -chdir="$WORK_DIR" plan -detailed-exitcode -input=false -out="$WORK_DIR/migration.tfplan" >/dev/null
plan_rc=$?
set -e

case $plan_rc in
  0) pass "plan after upgrade is empty (no diff)" ;;
  2) terraform -chdir="$WORK_DIR" show "$WORK_DIR/migration.tfplan"
     fail "plan after upgrade is NOT empty - the state migration leaves drift (see plan above)" ;;
  *) fail "terraform plan errored (exit $plan_rc)" ;;
esac

step "Phase 2b: confirm workspace is absent from the state file before apply"
if grep -q '"workspace": "default"' "$WORK_DIR/terraform.tfstate"; then
  fail "state file already contains \"workspace\": \"default\" before apply - phase ordering is broken"
fi
pass "state file has no \"workspace\": \"default\" before apply"

step "Phase 3: apply with the candidate provider"
terraform -chdir="$WORK_DIR" apply -auto-approve -input=false "$WORK_DIR/migration.tfplan" >/dev/null
info "applied"

step "Phase 3b: confirm workspace was backfilled in the state file after apply"
if ! grep -q '"workspace": "default"' "$WORK_DIR/terraform.tfstate"; then
  fail "state file does not contain \"workspace\": \"default\" after apply"
fi
pass "state file contains \"workspace\": \"default\" after apply"

step "Phase 4: assert workspace was backfilled in state"
# Ask the candidate provider itself which resource types gained `workspace`,
# so this stays correct as more resources adopt the attribute.
ws_type_set="$(terraform -chdir="$WORK_DIR" providers schema -json | jq -r '
  .provider_schemas
  | to_entries[]
  | select(.key | test("kong/konnect$"))
  | .value.resource_schemas
  | to_entries[]
  | select(.value.block.attributes.workspace != null)
  | .key')"
[[ -n "$ws_type_set" ]] || fail "candidate provider exposes no resource type with a workspace attribute"
info "$(wc -l <<<"$ws_type_set" | tr -d ' ') resource type(s) carry workspace"
checked=0
problems=""
while IFS=$'\t' read -r address type schema_version workspace; do
  [[ -n "$address" ]] || continue
  grep -qxF "$type" <<<"$ws_type_set" || continue
  checked=$((checked + 1))
  if [[ "$workspace" != "default" ]]; then
    problems+="  $address: workspace = ${workspace/@absent/<absent>}, want \"default\""$'\n'
  fi
  if [[ "$schema_version" != "1" ]]; then
    problems+="  $address: schema_version = $schema_version, want 1"$'\n'
  fi
done <<<"$(state_rows)"

[[ $checked -gt 0 ]] || fail "no workspace-capable resources in state - the fixture does not exercise the migration"
[[ -z "$problems" ]] || fail "workspace was not backfilled correctly:
$problems"
pass "all $checked workspace-capable resource(s) have workspace=\"default\", schema_version=1"

step "Phase 5: re-plan to confirm idempotency"
set +e
terraform -chdir="$WORK_DIR" plan -detailed-exitcode -input=false >/dev/null
replan_rc=$?
set -e
case $replan_rc in
  0) pass "second plan is still empty" ;;
  2) terraform -chdir="$WORK_DIR" plan
     fail "second plan is not empty - the upgrade is not idempotent" ;;
  *) fail "terraform plan errored (exit $replan_rc)" ;;
esac

printf '\n\033[1;32mMigration test passed\033[0m (%s -> working tree, %d resources checked)\n' \
  "$BASELINE_REF" "$checked"
