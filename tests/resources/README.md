# Acceptance tests

Each directory here is a **shard**: CI runs one job per directory, in parallel.
Put a new test in the directory for its product area. Each directory has its own
`testdata/` folder for its Terraform configs.

| Directory | What goes here | Notes |
|---|---|---|
| `gateway/` | Kong Gateway: control planes, services, routes, consumers, keys, vaults, all `gateway_plugin_*`, cloud gateway add-ons and custom domains | Each test creates its own control plane. |
| `aigateway/` | AI Gateway, AI cost, and the AI model / MCP catalog | |
| `eventgateway/` | Event Gateway | |
| `portalapi/` | Portals, APIs, API products, application auth strategies, DCR providers, service catalog, integrations | **Orgs allow only 4 portals.** Portal tests are the quota-limited ones. |
| `identityorg/` | Identity, authentication settings, PAT settings, teams, system accounts, user roles | Changes org-wide settings, so keep tests serial. |

The shared provider setup is in `tests/internal/acc`; each directory exposes it through a small `provider_test.go`.

## Rules

- A test belongs to exactly one directory. Test function names (`TestFoo`) must be unique across all of them.
- A shard must pass on its own. Don't read a resource another shard creates.
- Org quotas are shared by every shard and every concurrent run, for example 4 portals and 5 identity directories. Keep tests that need a limited resource in the shard that already owns it.
- A new area that fits none of these: add a directory with a `provider_test.go` (copy one) and add its name to the matrix in `.github/workflows/provider-tests.yaml`.
- Never put `*_test.go` files directly in `tests/resources/`. `make check-shards` fails if you do.

## Running

```sh
make acceptance-shard SHARD=gateway      # one shard (needs KONNECT_SPAT)
make acceptance                          # all shards
make check-shards                        # layout check, no credentials needed
```
