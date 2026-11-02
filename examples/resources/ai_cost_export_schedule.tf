# An export schedule creates one export for each of its periods as the period
# ends. `cadence` and `time_zone` cannot change after creation, so a change to
# either replaces the schedule. Set `status` to `paused` to stop it.
resource "konnect_ai_cost_export_schedule" "monthly_focus" {
  name    = "monthly-focus"
  cadence = "monthly"

  # `focus` writes FinOps FOCUS cost and usage data, one file per month. It needs
  # the `focus` block, and `parquet` goes with `zstd`.
  dataset     = "focus"
  format      = "parquet"
  compression = "zstd"

  focus = {
    version = "1.4"
  }

  status    = "active"
  time_zone = "UTC"

  # Only usage from these tags is exported. Omit `filters` to export everything.
  filters = jsonencode({
    "tags.cost_center" = { in = ["platform", "ml"] }
  })

  labels = {
    owner = "finops"
  }
}
