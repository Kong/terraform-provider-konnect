#!/usr/bin/env bash
# Helpers for the sharded acceptance tests.
#
# A shard is a directory under tests/resources/. Put a new test in the
# directory for its product area; CI runs one job per directory.
#
#   acceptance-shard.sh shards              print shard (directory) names
#   acceptance-shard.sh check               fail on tests outside a shard directory
#   acceptance-shard.sh run <shard>         run one shard (needs TF_ACC=1 and credentials)
#   acceptance-shard.sh verify <dir> [base] check shard results (*.json from `run`) in <dir>
#                                           against the test inventory, optionally diff
#                                           against a baseline results file
#   acceptance-shard.sh summarize <json>... print "<test> <outcome>" lines
#   acceptance-shard.sh report <dir> [base-ref]  markdown table of results per shard (for
#                                           $GITHUB_STEP_SUMMARY); base-ref adds a before/after count
set -euo pipefail

ROOT=tests/resources

shard_names() { find "$ROOT" -mindepth 1 -maxdepth 1 -type d | sed "s|^$ROOT/||" | sort; }

all_tests() { go test -list '.*' "./$ROOT/..." | grep '^Test' | sort; }

cmd_check() {
  local stray empty s
  stray=$(find "$ROOT" -maxdepth 1 -name '*_test.go')
  if [[ -n "$stray" ]]; then
    echo "Test files must live in a shard directory under $ROOT/, not in $ROOT/ itself:" >&2
    echo "$stray" >&2
    echo "Shards: $(shard_names | paste -sd' ' -). See $ROOT/README.md." >&2
    exit 1
  fi
  for s in $(shard_names); do
    ls "$ROOT/$s"/*_test.go > /dev/null 2>&1 || empty="${empty:-} $s"
  done
  if [[ -n "${empty:-}" ]]; then echo "Shards with no tests:$empty" >&2; exit 1; fi
  local dup; dup=$(all_tests | uniq -d)
  if [[ -n "$dup" ]]; then echo "Test names must be unique across shards:" >&2; echo "$dup" >&2; exit 1; fi
  echo "OK: $(all_tests | wc -l | tr -d ' ') tests"
  for s in $(shard_names); do printf '  %-14s %s\n' "$s" "$(go test -list '.*' "./$ROOT/$s" | grep -c '^Test')"; done
}

cmd_run() {
  [[ -d "$ROOT/$1" ]] || { echo "unknown shard '$1'; shards: $(shard_names | paste -sd' ' -)" >&2; exit 1; }
  local out="${OUT_DIR:-.}/shard-$1.json"
  mkdir -p "$(dirname "$out")"
  TF_ACC=1 go test -json -timeout "${TEST_TIMEOUT:-30m}" ${PARALLEL:+-parallel "$PARALLEL"} \
    "./$ROOT/$1" | tee "$out" | jq -j '.Output // empty'
}

# One "<test> <outcome>" line per executed test and subtest.
summarize() {
  cat "$@" | jq -r 'select(.Test != null and (.Action=="pass" or .Action=="fail" or .Action=="skip")) | "\(.Test) \(.Action)"' | sort
}

# Markdown report: one row per shard plus totals, and an optional comparison
# with the number of tests defined on <base-ref> (the "before" layout).
cmd_report() {
  local dir=$1 base_ref=${2:-} f name t p fl sk sub tp tf ts tt=0 inv
  tp=0; tf=0; ts=0
  echo "## Acceptance test results"
  echo
  echo "| Shard | Tests | Passed | Failed | Skipped | Subtests run |"
  echo "|---|---:|---:|---:|---:|---:|"
  for f in "$dir"/shard-*.json; do
    [[ -e "$f" ]] || continue
    name=$(basename "$f" .json); name=${name#shard-}
    read -r p fl sk sub < <(jq -rs '
      [.[] | select(.Test != null and (.Action=="pass" or .Action=="fail" or .Action=="skip"))] as $r
      | [($r|map(select((.Test|contains("/")|not) and .Action=="pass"))|length),
         ($r|map(select((.Test|contains("/")|not) and .Action=="fail"))|length),
         ($r|map(select((.Test|contains("/")|not) and .Action=="skip"))|length),
         ($r|map(select(.Test|contains("/")))|length)] | @tsv' "$f")
    t=$((p+fl+sk)); tt=$((tt+t)); tp=$((tp+p)); tf=$((tf+fl)); ts=$((ts+sk))
    echo "| $name | $t | $p | $fl | $sk | $sub |"
  done
  echo "| **Total** | **$tt** | **$tp** | **$tf** | **$ts** | |"
  echo
  inv=$(all_tests | wc -l | tr -d ' ')
  echo "- Tests defined in this branch: **$inv**; tests that ran: **$tt**."
  if [[ -n "$base_ref" ]] && git rev-parse -q --verify "$base_ref" > /dev/null; then
    local before; before=$(git grep -h '^func Test' "$base_ref" -- "$ROOT" | wc -l | tr -d ' ')
    echo "- Tests defined on \`$base_ref\` (single job, before sharding): **$before**."
    if [[ "$before" -eq "$inv" ]]; then echo "- Same number of tests before and after sharding."
    else echo "- Number of tests differs from \`$base_ref\` by $((inv-before)) (expected only if this branch adds or removes tests)."; fi
  fi
}

cmd_verify() {
  local dir=$1 base=${2:-} res; res=$(mktemp)
  summarize "$dir"/*.json > "$res"
  local pass fail skip rc=0 missing dup
  pass=$(grep -c ' pass$' "$res" || true); fail=$(grep -c ' fail$' "$res" || true); skip=$(grep -c ' skip$' "$res" || true)
  echo "results: pass=$pass fail=$fail skip=$skip"
  missing=$(comm -23 <(all_tests) <(awk '$1 !~ /\// {print $1}' "$res" | sort -u))
  dup=$(awk '$1 !~ /\// {print $1}' "$res" | sort | uniq -d)
  [[ -z "$missing" ]] || { echo "NOT RUN:" >&2; echo "$missing" >&2; rc=1; }
  [[ -z "$dup" ]] || { echo "RAN MORE THAN ONCE:" >&2; echo "$dup" >&2; rc=1; }
  [[ "$fail" -eq 0 ]] || { echo "FAILED:" >&2; grep ' fail$' "$res" >&2; rc=1; }
  if [[ -n "$base" ]]; then
    if diff <(sort "$base") "$res" > "$res.diff"; then
      echo "results identical to baseline $base"
    else
      echo "results differ from baseline:" >&2; cat "$res.diff" >&2; rc=1
    fi
  fi
  cp "$res" "${RESULTS_OUT:-sharded-results.txt}"
  exit $rc
}

case "${1:-}" in
  shards)    shard_names ;;
  check)     cmd_check ;;
  run)       cmd_run "${2:?shard}" ;;
  report)    cmd_report "${2:?dir}" "${3:-}" ;;
  verify)    cmd_verify "${2:?dir}" "${3:-}" ;;
  summarize) shift; summarize "$@" ;;
  *) sed -n '2,17p' "$0"; exit 2 ;;
esac
