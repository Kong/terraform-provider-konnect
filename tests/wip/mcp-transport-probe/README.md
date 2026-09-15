# MCP `packages[*].transport` anyOf probe

Goal: empirically verify what the API accepts and how Terraform behaves when
switching between the `anyOf` branches of `MCPTransport`
(`mcp_stdio_transport`, `mcp_streamable_http_transport`, `mcp_sse_transport`).

## Files

| File | Purpose |
| --- | --- |
| `_provider.tf` | Provider + auth. Fill in your PAT + server URL. |
| `_base.tf`     | Parent `konnect_catalog_mcp` shared by every scenario. |
| `scenario_1_stdio.tf` | Baseline: only `mcp_stdio_transport`. |
| `scenario_2_streamable_http.tf` | Update: switch to `mcp_streamable_http_transport`. |
| `scenario_3_sse.tf` | Update: switch to `mcp_sse_transport`. |
| `scenario_4_multiple.tf` | Violation: set two branches at once. |
| `scenario_5_empty.tf` | Violation: empty `transport = {}`. |

Only **one** `scenario_*.tf` file should be present at a time. Rename the others
to `.tf.off` (Terraform ignores non-`.tf` files) between runs:

```bash
mv scenario_1_stdio.tf scenario_1_stdio.tf.off
mv scenario_2_streamable_http.tf.off scenario_2_streamable_http.tf
```

## Workflow

```bash
cd tests/wip/mcp-transport-probe

# 1. Baseline
terraform init
terraform apply -auto-approve
terraform plan                # expect: No changes.
terraform state show konnect_catalog_mcp_version.probe | sed -n '/transport/,/^  }/p'

# 2. Switch to streamable-http (the important update test)
mv scenario_1_stdio.tf scenario_1_stdio.tf.off
mv scenario_2_streamable_http.tf.off scenario_2_streamable_http.tf
terraform plan                # inspect update plan
terraform apply -auto-approve
terraform plan                # expect: No changes.  <-- KEY CHECK

# 3. Switch to sse
mv scenario_2_streamable_http.tf scenario_2_streamable_http.tf.off
mv scenario_3_sse.tf.off scenario_3_sse.tf
terraform plan
terraform apply -auto-approve
terraform plan                # expect: No changes.

# 4. Two branches at once (should fail; see how it fails)
mv scenario_3_sse.tf scenario_3_sse.tf.off
mv scenario_4_multiple.tf.off scenario_4_multiple.tf
terraform plan                # inspect
terraform apply               # capture error message

# 5. Empty (should fail)
mv scenario_4_multiple.tf scenario_4_multiple.tf.off
mv scenario_5_empty.tf.off scenario_5_empty.tf
terraform plan
terraform apply

# 6. Import round-trip
# Grab an existing MCP version ID (e.g. from scenario 1) and:
mv scenario_5_empty.tf scenario_5_empty.tf.off
mv scenario_1_stdio.tf.off scenario_1_stdio.tf
terraform state rm konnect_catalog_mcp_version.probe
terraform import konnect_catalog_mcp_version.probe "<version-id>"
terraform plan                # expect: No changes.
```

## What to record for the reviewer

| Scenario | apply result | plan after apply | API GET body (`transport`) |
| --- | --- | --- | --- |
| 1. stdio | | | |
| 2. streamable_http (update) | | | |
| 3. sse (update) | | | |
| 4. two branches | | | |
| 5. empty | | | |
| 6. import round-trip | | | |

Row 2 and row 6 are the ones the reviewer specifically cares about.

## Sibling anyOf fields you can piggyback

The same schema shape (named `anyOf` branches) appears in:

- `packages[*].package_arguments[*]` → `mcp_named_argument` / `mcp_positional_argument`
- `packages[*].runtime.arguments[*]` → same
- `remotes[*]` → `mcp_sse_transport` / `mcp_streamable_http_transport`

Run the same 6-scenario matrix on `remotes` if you want full coverage; the
mechanics are identical.

