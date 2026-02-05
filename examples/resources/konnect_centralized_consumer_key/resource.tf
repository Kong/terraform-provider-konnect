resource "konnect_centralized_consumer_key" "my_centralizedconsumerkey" {
  consumer_id = "...my_consumer_id..."
  realm_id    = "75fbe6fe-cfd5-48fb-8590-5c6e31cc0859"
  secret      = "...my_secret..."
  tags = [
    "..."
  ]
  type = "legacy"
}