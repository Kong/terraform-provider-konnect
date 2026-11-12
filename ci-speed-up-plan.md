# CI speed-up plan: shard `tests/resources` across parallel jobs

## Current state

- `.github/workflows/provider-tests.yaml` has one job, `acceptance`, which runs `make acceptance`.
- `make acceptance` runs `TF_ACC=1 go test -v -timeout 30m ./tests/resources`. That is one `go test` process over about 128 test files and about 1,400 `resource.Test` steps.
- `t.Parallel()` is used in only 37 of the 128 files, so most tests run serially. Parallelism is also capped by `-parallel` (default `GOMAXPROCS`, which is 2-4 on `ubuntu-latest`).
- The test run time is dominated by waiting on the Konnect API. Gateway plugin tests alone are about 70 files.
- The workflow comment says the tests "are not namespaced and will conflict". Only one Terraform version is run, and concurrent PR runs share the same org.
- Most test data (`testdata/<TestName>/...`) creates its own control plane or portal, with hard-coded names such as `tf-test-config-store`. There is no randomization.

## Approach: Go-level shards, one GitHub Actions matrix job per shard

Use a `strategy.matrix` where each entry is a shard name plus a `go test -run` regex. Every shard is a separate runner, but the code and `make` target are the same. This is better than the alternatives below:

| Option | Verdict |
|---|---|
| Matrix with `-run` regexes (**recommended**) | Needs no test code changes. Each shard can be rerun alone. The failing group is visible in the PR checks. |
| Split into Go packages (`tests/resources/gateway`, `/portal`, ...) | Cleaner long-term, and `go test ./...` parallelizes by package. But it means moving about 128 files and sharing `provider_test.go` helpers. Do this later, if at all. |
| Raise `-parallel` / add `t.Parallel()` everywhere | Cheap, but hits Konnect rate limits and the org-level conflicts below. Use it only inside a shard. |
| Auto-balancing by timing (e.g. `gotestsum --partition`) | Over-engineered for now. Revisit once the shards are measured. |

## Proposed shards (first iteration)

Start coarse: fewer shards, grouped by domain. Split further only if the measured timings say a shard is too slow (see "Revisit criteria").

| Shard | Files (prefix) | Why it is grouped |
|---|---|---|
| `gateway` | Everything Kong Gateway: `gateway_*` (control plane, service, route, consumer, key, vault, partial, ...), all `gateway_plugin_*` including the AI plugins, `gateway_custom_plugin*`, `gateway_cloned_plugin`, `centralized_consumers`, `cloud_gateway_addon`, `cloud_gateway_custom_domain`, `platform_available_regions` | Control plane and plugin tests are related, and each test builds its own control plane. This will be the largest shard and the one to watch. |
| `ai-gateway` | `ai_gateway`, `ai_cost_*`, `catalog_ai_model`, `catalog_mcp` | A separate product surface from Kong Gateway. |
| `event-gateway` | `event_gateway`, `event_gateway_1.2`, `event_gateway_1.3` | Independent of the others. |
| `portal-api` | `portal_*`, `api`, `api_*`, `api_product*`, `application_auth_strategy`, `dcr_provider`, `catalog_service`, `integration_instance` | Portal and API products are used together (publications, auth strategies). |
| `identity-org` | `identity*`, `authentication_settings`, `organization_pat_settings`, `system_account*`, `team*`, `user_role`, `basic_auth`, `provider` | **Org-level singletons.** These tests mutate shared org settings, so they must run in one shard and not in parallel with each other. |

Not every test here is verified independent. Step 1 below confirms the grouping before the split ships.

### Revisit criteria

After the first few runs, split a shard only if both hold: it is the slowest shard by a clear margin (for example more than 1.5x the next one), and it is still the bottleneck for wall-clock time. The likely candidate is `gateway`, which would split into core and plugins. If the split shows no meaningful wall-clock gain, merge the shards back.

## Shard self-sufficiency

**Rule:** a shard must pass when run alone, in any order, on a fresh runner, with nothing run before it and nothing running beside it that it depends on. No shard may read, or assume the existence of, a resource that another shard creates.

What the audit of `tests/resources` found:

- **No cross-test helpers or state.** The only shared Go code is `provider_test.go` (provider factory and the `E2E_SERVER_*` env vars). No test reads a variable or output from another test.
- **Every gateway, portal and event-gateway test creates its own parent resource** (control plane, portal, ...) inside its own `testdata/<TestName>/` directory, so those shards are self-sufficient already.
- **Pre-existing org fixtures (not created by any test).** These are the real self-sufficiency gaps, because a shard only works if the org already has them:

  | Test | Assumes in the org | Shard |
  |---|---|---|
  | `TestTeam/data` | a team named `portal-admin` | `identity-org` |
  | `TestSystemAccount/data` | a system account named `Sample System Account` | `identity-org` |
  | `TestOrgPATSettings` | hard-coded organization id `72ab1dcd-...` | `identity-org` |
  | `TestUserRole` | hard-coded user id `df5404e1-...` (a user can't be created through Terraform) | `identity-org` |

  Each one lives in a single shard, so the split doesn't break it. But they depend on the org's state, so list them in the README as org prerequisites, or change the tests to create the fixture themselves (a team and a system account can be created, so those two are easy).
- **Org-wide data sources and singletons.** `*_list` data sources, `authentication_settings`, and the identity-provider group mappings read or change org-wide state. Keep them in `identity-org`, so no other shard sees them half-changed. Check that list assertions don't depend on an exact count of objects that other shards create and delete concurrently.
- **Shared environment.** Every shard needs the same `KONNECT_SPAT`, Go, and the Terraform binary (set up in each matrix job). Nothing is passed between jobs except the result artifacts used for validation.

Enforcement:

1. Run each shard alone, from a clean checkout, at least once before merging (the matrix does this by design, since every job is a fresh runner).
2. Add a script check that fails if a new `testdata` directory uses a `data` block for a resource that its own directory doesn't create, unless it is on a short, reviewed allowlist (the table above).

## Org resource quotas (the real concurrency ceiling)

Konnect limits how many of some resources an org can hold at once. Examples given for this org: **4 portals** and **5 directories**. Every test in every running shard draws from the same pool, so quotas, and not CPU or runner count, set how much can run at once. A test that hits a quota fails with an error unrelated to what it is testing.

What the tests need today (count of tests whose own config creates the resource; each uses 1 per config unless noted):

| Resource | Tests creating it | Known limit | Consequence |
|---|---|---|---|
| Portals | about 15: 11 in `testdata` (`TestPortalV3`, `Snippet`, `Page`, `Logo`, `Favicon`, `Developer`, `Customization`, ...) plus 4 defined inline in Go (`portal_ip_allow_list`, `portal_audit_log_webhook`, `portal_idp_team_group_mapping`, `portal_identity_provider`) | 4 per org | At most 4 portal tests can run at the same time, across **all shards and all concurrent PR runs**. |
| Identity directories | `TestIdentityDirectory` | 5 per org | Fine alone. Becomes a problem if another run holds directories at the same time. |
| Control planes, AI gateways, event gateways, teams, system accounts | many | **no limit** | Not a constraint. Only API rate limits and runner memory apply. |
| Cloud networks | none (not created in CI) | n/a | Too slow and expensive, so `TestCloudGatewayNetwork` is skipped with `t.Skip`. The addon and custom-domain tests use a control plane with `cloud_gateway = true`, not a network. |

Three things follow from this:

1. **Sharding doesn't create more quota.** Moving portal tests into their own job doesn't help if they are still launched together. The portal tests are the tightest case: about 15 tests and a limit of 4. Portals and directories are the only quota-limited resources, so they are the only thing that needs a cap.
2. **Concurrency must be capped per resource type, not just by `-parallel`.** The suite needs a counting semaphore for each limited resource type, shared by the tests in a shard. For example, a test that creates a portal takes one of 4 slots before `resource.Test`, and releases it after. The cap for each type is the quota minus headroom for other runs, so use 2 or 3 of 4 portals, not all 4.
3. **Shards must not share a quota.** Two shards that each use up to 3 portals break a limit of 4. Give each limited type to **one** shard, or give each shard its own fixed slice. This is another reason `portal-api` is a single shard that owns the portal quota, and `identity-org` owns directories.

Options for enforcing the caps, in order of preference:

| Option | Notes |
|---|---|
| In-process semaphore in the test helpers (a small `acquire("portal")` wrapper) | Simple, no new infrastructure. Only limits tests within one `go test` process, which is correct as long as each limited type belongs to exactly one shard. |
| Run the quota-limited tests with `-parallel 1` or `-parallel N` in their own shard (`-run` regex) | No code change. Coarse, but enough for a first iteration: for example, `portal` runs serially (about 15 portals one after another) and everything else runs wide. |
| A job-level `concurrency:` group per resource type | Serializes whole jobs across PRs too, which also protects against two PRs using the quota together. It costs queueing time. |
| Cleanup before and after | A sweeper (see Risks) that deletes leftovers, because one leaked portal from a cancelled run can use up a slot for every later run. |

Cross-PR concurrency still matters: two PRs running at the same moment share the org and its quotas. Either add `concurrency: { group: acceptance-konnect-org, cancel-in-progress: false }` on the workflow, so only one acceptance run uses the org at a time (simple, but PRs queue), or move the tests to separate Konnect orgs per run. The second option removes the quota problem for good but needs more orgs and credentials.

## Impact of `t.Parallel()` in every file

Currently 37 of 128 test files call `t.Parallel()` at the top level. The other 91 run one after another. This is **not measured**, so treat the numbers below as expectations to check in step 1 of the implementation.

How Go behaves here:

- Go runs all the non-parallel top-level tests first, one at a time. Tests that call `t.Parallel()` pause and are released together only after the serial ones finish. So today the run is about 91 files in sequence, then 37 files in a parallel burst.
- Only top-level tests are parallelized. The `t.Run(...)` subtests inside one test, such as `update-nullify-fields`, still run in sequence unless each subtest also calls `t.Parallel()`.
- The number running at once is capped by `-parallel`, which defaults to `GOMAXPROCS` (4 on a standard GitHub runner). These tests mostly wait on the Konnect API, so the cap can be much higher, around 8 to 16, without CPU being a problem. Memory is the limit: each running test starts a `terraform` process and a provider process.

Expected effect, if every file were parallel:

- The best case is a big gain **within a shard**. The serial part, which is most of today's runtime, would run `-parallel` tests at once. For API-bound tests, the speed-up is limited by the longest single test and by the cap, not by the file count.
- It also stacks with sharding. Sharding splits the work across runners, and `t.Parallel()` makes each runner busier. If the slowest shard is dominated by a few long tests, more parallelism won't help it.

Why it is not a free win:

| Risk | Detail |
|---|---|
| **Org-wide name collisions** | The portal tests all create a portal named `"My v3 portal name"` (8 tests). If portal names must be unique in an org (to be confirmed), running those in parallel fails. Names like `HTTPBin` or `tf-test-consumer` are scoped to a control plane or gateway that each test creates itself, so they don't collide. |
| **Org singletons** | `authentication_settings`, `organization_pat_settings`, and the identity-provider mappings change shared state. These must stay serial. |
| **Shared fixtures** | The tests that read the pre-existing `portal-admin` team or `Sample System Account` are safe in parallel as readers, but not if another test modifies them. |
| **Org quotas** | about 15 portal tests against a limit of 4 portals per org, and 5 directories. Blanket `t.Parallel()` would exceed these. See "Org resource quotas". |
| **Rate limits** | The Konnect API may return 429s with many concurrent calls. This gets worse with sharding plus parallelism together. |
| **Memory and flakiness** | Many concurrent Terraform processes use real memory. A timeout on one slow resource can fail the whole `go test` run at the `-timeout` limit. |
| **Harder debugging** | Interleaved `-v` logs from many tests are harder to read. Use `-json` and `gotestsum` for per-test output. |

Recommendation:

1. Do sharding first (it is the safer, structural change), and don't add `t.Parallel()` to everything at the same time. Otherwise a regression can't be attributed to either change.
2. Then opt files in by shard, starting with `gateway`, where every test owns its control plane. Raise `-parallel` for that shard only.
3. Leave `identity-org` serial. Make the portal tests parallel only after fixing the shared portal name, and then cap them at the portal quota (see "Org resource quotas").
4. Compare against the baseline using the validation section: the pass count and the set of tests must stay identical, and the runtime must drop. Run each change a few times to check for flakiness before keeping it.

## Validating the test count before and after

The goal is to prove the sharded run executes exactly the same tests as today's single job, with the same results. Neither the count nor the set of names may drift.

1. **Baseline, before any workflow change.** On the current branch, produce the full inventory and one real run:
   ```sh
   go test -list '.*' ./tests/resources | grep '^Test' | sort > baseline-tests.txt
   TF_ACC=1 go test -json -timeout 30m ./tests/resources > baseline.json
   ```
   `-list` gives the top-level test functions. The JSON gives every pass, fail and skip, including subtests such as `TestAIGateway/update-nullify-fields`.
2. **Summarize into a comparable file.** One line per test with its outcome, sorted:
   ```sh
   jq -r 'select(.Test != null and (.Action=="pass" or .Action=="fail" or .Action=="skip")) | "\(.Test) \(.Action)"' baseline.json | sort > baseline-results.txt
   ```
   Subtests are included on purpose, so the count is the real number of executed cases and not just 128 files.
3. **After.** Each shard uploads its own `-json` file as an artifact. A final aggregate job downloads all of them, concatenates the same `jq` summary into `sharded-results.txt`, and compares.
4. **Checks the aggregate job enforces** (fail the build on any of them):
   - **Coverage of the inventory.** Every name in `go test -list` appears in exactly one shard's results. No test missing (not matched by any regex) and none run twice (matched by two).
   - **Count equality.** The number of `pass` lines in `sharded-results.txt` equals the baseline `pass` count. Same for `skip`, and `fail` must be 0.
   - **Set equality.** `diff baseline-results.txt sharded-results.txt` is empty, apart from known flaky tests listed in an allowlist.
5. **Cheap pre-merge check, no API calls needed.** Run only the coverage part (the first bullet in step 4) locally with `go test -list` against each shard regex. It proves the split is complete and non-overlapping before spending any CI time. Keep it as a permanent guard script, for example `scripts/check-shards.sh`, so a new test file that no shard matches fails CI instead of being silently skipped.
6. **Flaky and skipped tests.** Run the baseline twice, so tests that already flap are known up front and don't look like regressions from the split. Record the skip count too, since a test that starts skipping because an env var is missing in a shard looks like a pass otherwise.

## Risks to handle

1. **Org-level singletons.** `authentication_settings`, `organization_pat_settings` and the identity-provider group mappings change settings for the whole org. Keep them in one shard. Do not use `t.Parallel()` in that shard.
2. **Name collisions between concurrent runs.** Names are hard-coded, so two PRs running at once, or two shards creating a same-named control plane, can conflict. Mitigations:
   - Short term: add a `concurrency:` group on the workflow, as `provider-tests-${{ github.ref }}` with `cancel-in-progress: true`, so one PR doesn't overlap itself.
   - Next: audit that resource names that must be unique org-wide (control planes, portals, teams, system accounts) are unique across shards. Cross-PR safety needs a per-run suffix, e.g. an `E2E_RUN_ID` env var that the HCL reads through a `variable`. That is a larger change, so do it as a follow-up and not as a blocker.
3. **Konnect API rate limits.** More parallel runners means more concurrent calls. Start with 5 shards, watch for 429s, and add retry/backoff if they appear.
4. **Secrets and cost.** Each shard uses `KONNECT_SPAT`. Fork PRs don't get it, which is unchanged from today.
5. **Required status checks.** The branch protection rule currently names `Acceptance Tests (Terraform 1.13.*)`. Matrix job names change. Add a final aggregate job (`needs: [acceptance]`, `if: always()`) with a stable name, and make that the required check.
6. **Leaked resources on cancel.** `cancel-in-progress` can kill a run mid-test and leave orphaned control planes. Check whether a cleanup job already exists, and add a scheduled sweeper if not.

## Implementation steps

1. **Measure and baseline.** Capture the baseline from "Validating the test count". Run `go test -json ./tests/resources` once and record per-test durations (`gotestsum --jsonfile` works). Use the result to balance the shards, since file count is a poor proxy for time. Also confirm which tests create org-level singletons: `grep -l 'authentication_settings\|organization_pat\|identity_provider\|konnect_team\|system_account' tests/resources/testdata -r`.
2. **Add a Makefile target** that takes a regex:
   ```make
   acceptance-shard:
   	@TF_ACC=1 go test -v -timeout 30m -run '$(RUN)' ./tests/resources
   ```
   Keep `make acceptance` as the run-everything target for local use.
3. **Convert the workflow to a matrix**, with `shard` (name) and `run` (regex) entries, and `fail-fast: false`. Upload each shard's `-json` output as an artifact.
4. **Add the shard-coverage guard and the aggregate comparison** described in "Validating the test count before and after".
5. **Add the aggregate job and the `concurrency:` group**, then update branch protection.
6. **Rebalance** after the first few runs, using the real timings. Aim for the slowest shard under about 10 minutes.
7. **Follow-ups (optional):** a per-run name suffix, per-shard `-parallel` tuning, and splitting into Go packages.

## Expected outcome

The wall-clock time becomes roughly that of the slowest shard, and the total runner time stays about the same. The plugin tests are the largest block, which is why they are split in two. Once they are measured, the final shard boundaries may move.

## Open questions

- None on quotas. Portals (4) and directories (5) are the only limited resources. Control planes, AI gateways, event gateways, teams and system accounts have no limit, and cloud networks aren't used in CI.

- Is there an existing sweeper for leaked test resources in the Konnect org?
- Is the org's API rate limit known for the shared `KONNECT_SPAT`?
- Do the `ai_gateway` or `catalog_*` tests depend on a gateway control plane from other shards? The testdata suggests not, but this needs confirming.

## Implementation status: shards are directories

The shards are directories under `tests/resources/` (`gateway`, `aigateway`, `eventgateway`, `portalapi`, `identityorg`), each its own Go package with its own `testdata/`. This replaced an earlier version that assigned tests to shards with name regexes in a script, which hid the rule from developers. Now the folder is the shard, `tests/resources/README.md` explains where a test goes, and `make check-shards` fails if a test file is left in `tests/resources/` itself. The shared provider setup is in `tests/internal/acc`. The shard names in the earlier sections of this plan map as: `ai-gateway` is `aigateway`, `event-gateway` is `eventgateway`, `portal-api` is `portalapi`, `identity-org` is `identityorg`.
