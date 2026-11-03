# A limit budget caps spend for the usage its scope selects, per window. `scope`
# and `window` change in place, and a change applies to the whole current budget
# window.
resource "konnect_ai_cost_budget" "team_models" {
  limit = {
    name     = "team-models-monthly"
    priority = 10
    posture  = "enforce"

    currency = "USD"
    value    = "5000"

    window = {
      type     = "fixed"
      interval = "P1M"
      anchor   = "2026-01-01T00:00:00+01:00[Europe/Budapest]"
    }

    scope = {
      # One derived budget per model in the openai provider.
      each = "model"

      filters = jsonencode([
        { field = "provider", op = "eq", value = "openai" },
        { field = "tags.cost_center", op = "in", value = ["platform", "ml"] },
      ])
    }

    labels = {
      owner = "finops"
    }
  }
}

# An exemption declares that usage in its scope is exempt from limits for a
# period. It sets no limit, so it takes `validity` instead of `limit`, `window`
# and `posture`.
resource "konnect_ai_cost_budget" "free_tier" {
  exempt = {
    name     = "free-tier-exemption"
    priority = 100

    scope = {
      each    = null
      filters = jsonencode([{ field = "tags.tier", op = "eq", value = "free" }])
    }

    validity = {
      from  = "2026-01-01T00:00:00Z"
      until = "2026-12-31T23:59:59Z"
    }
  }
}
