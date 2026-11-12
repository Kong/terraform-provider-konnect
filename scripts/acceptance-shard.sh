#!/usr/bin/env bash
# Splits the acceptance tests in ./tests/resources into independent shards.
#
#   acceptance-shard.sh shards              print shard names
#   acceptance-shard.sh list <shard>        print the test functions in a shard
#   acceptance-shard.sh regex <shard>       print the `go test -run` regex for a shard
#   acceptance-shard.sh check               fail unless every test is in exactly one shard
#   acceptance-shard.sh run <shard>         run one shard (needs TF_ACC=1 and credentials)
#   acceptance-shard.sh verify <dir> [base] check shard results in <dir> (*.json from `run`)
#                                           against the test inventory, optionally diff
#                                           against a baseline results file
#
# A test belongs to the FIRST shard whose pattern matches its function name.
# See ci-speed-up-plan.md for why the shards are drawn this way.
set -euo pipefail

PKG=./tests/resources

# shard:extended-regex on the Go test function name (first match wins)
RULES=(
  "ai-gateway:^Test(AIGateway|AICost|CatalogAiModel|CatalogMCP)"
  "event-gateway:^TestEventGateway"
  "identity-org:^Test(Identity|AuthenticationSettings|OrgPATSettings|SystemAccount|Team|UserRole|BasicAuth|Provider)"
  "portal-api:^Test(Portal|API|ApplicationAuthStrategy|DcrProvider|CatalogService|IntegrationInstance)"
  "gateway:^Test(Gateway|CloudGateway|CentralizedConsumers|PlatformAvailableRegions|Vault)"
)

shard_names() { for r in "${RULES[@]}"; do echo "${r%%:*}"; done; }

all_tests() { go test -list '.*' "$PKG" | grep '^Test' | sort; }

# print "<shard> <test>" for every test; shard is "NONE" when nothing matches
# and "a+b" when several do.
assign() {
  local t r hits
  while read -r t; do
    hits=""
    for r in "${RULES[@]}"; do
      if [[ "$t" =~ ${r#*:} ]]; then hits="${hits:+$hits+}${r%%:*}"; fi
    done
    echo "${hits:-NONE} $t"
  done < <(all_tests)
}

list_shard() { assign | awk -v s="$1" '$1==s {print $2}'; }

# The rules are first-match-wins, so a test matching several patterns is
# reported by `check`; `list` only reports the exact single-shard matches.
regex_for() {
  local names
  names=$(list_shard "$1" | paste -sd'|' -)
  [[ -n "$names" ]] || { echo "shard '$1' has no tests" >&2; exit 1; }
  echo "^(${names})\$"
}

cmd_check() {
  local bad total
  bad=$(assign | awk '$1=="NONE" || $1 ~ /\+/')
  total=$(all_tests | wc -l | tr -d ' ')
  if [[ -n "$bad" ]]; then
    echo "Tests not in exactly one shard:" >&2
    echo "$bad" >&2
    exit 1
  fi
  echo "OK: $total tests, each in exactly one shard"
  assign | awk '{c[$1]++} END{for(s in c) printf "  %-14s %d\n", s, c[s]}' | sort
}

cmd_run() {
  local shard=$1 out="${OUT_DIR:-.}/shard-$1.json"
  local re; re=$(regex_for "$shard")
  mkdir -p "$(dirname "$out")"
  echo "shard=$shard run=$re" >&2
  TF_ACC=1 go test -json -timeout "${TEST_TIMEOUT:-30m}" ${PARALLEL:+-parallel "$PARALLEL"} \
    -run "$re" "$PKG" | tee "$out" | jq -j '.Output // empty'
}

# One "<test> <outcome>" line per executed test and subtest.
summarize() {
  cat "$@" | jq -r 'select(.Test != null and (.Action=="pass" or .Action=="fail" or .Action=="skip")) | "\(.Test) \(.Action)"' | sort
}

cmd_verify() {
  local dir=$1 base=${2:-} res; res=$(mktemp)
  summarize "$dir"/*.json > "$res"
  local pass fail skip
  pass=$(grep -c ' pass$' "$res" || true); fail=$(grep -c ' fail$' "$res" || true); skip=$(grep -c ' skip$' "$res" || true)
  echo "results: pass=$pass fail=$fail skip=$skip"
  local rc=0
  # every top-level test must have run exactly once
  local missing dup
  missing=$(comm -23 <(all_tests) <(awk '$1 !~ /\// {print $1}' "$res" | sort -u))
  dup=$(awk '$1 !~ /\// {print $1}' "$res" | sort | uniq -d)
  [[ -z "$missing" ]] || { echo "NOT RUN:" >&2; echo "$missing" >&2; rc=1; }
  [[ -z "$dup" ]] || { echo "RAN MORE THAN ONCE:" >&2; echo "$dup" >&2; rc=1; }
  [[ "$fail" -eq 0 ]] || { echo "FAILED:" >&2; grep ' fail$' "$res" >&2; rc=1; }
  if [[ -n "$base" ]]; then
    if diff <(sort "$base") "$res" > /tmp/results.diff; then
      echo "results identical to baseline $base"
    else
      echo "results differ from baseline:" >&2; cat /tmp/results.diff >&2; rc=1
    fi
  fi
  cp "$res" "${RESULTS_OUT:-sharded-results.txt}"
  exit $rc
}

case "${1:-}" in
  shards) shard_names ;;
  list)   list_shard "${2:?shard}" ;;
  regex)  regex_for "${2:?shard}" ;;
  check)  cmd_check ;;
  run)    cmd_run "${2:?shard}" ;;
  verify) cmd_verify "${2:?dir}" "${3:-}" ;;
  summarize) shift; summarize "$@" ;;
  *) sed -n '2,15p' "$0"; exit 2 ;;
esac
