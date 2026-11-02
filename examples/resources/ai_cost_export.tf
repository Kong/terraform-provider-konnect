# A one-time export copies the usage events of a time range into files to
# download. It runs asynchronously: `status` follows its progress and `files`
# lists the data files once it succeeds. Everything is set at creation, so a
# change to any attribute replaces the export.
resource "konnect_ai_cost_export" "usage_november" {
  dataset = "usage_events"

  # `csv` accepts `none` or `gzip`, `parquet` accepts `zstd`.
  format      = "csv"
  compression = "gzip"

  # A fixed day in the past. The API rejects a range that ends after the day in
  # progress, so an example cannot use a range near today.
  from      = "2026-09-01T00:00:00Z"
  to        = "2026-09-02T00:00:00Z"
  time_zone = "Europe/Budapest"

  # Omit `filters` to export every event in the range. The shape is a record keyed
  # by usage field, each value naming one operator. A `tags.<key>` key filters the
  # value of that tag.
  filters = jsonencode({
    provider           = { eq = "openai" }
    "tags.cost_center" = { in = ["platform", "ml"] }
  })

  labels = {
    reason = "november-close"
  }
}
