resource "konnect_centralized_consumer" "my_centralizedconsumer" {
  consumer_groups = [
    "..."
  ]
  custom_id = "...my_custom_id..."
  realm_id  = "44e96a35-843f-4590-b0ec-3a958e6297cb"
  tags = [
    "..."
  ]
  type     = "application"
  username = "...my_username..."
}