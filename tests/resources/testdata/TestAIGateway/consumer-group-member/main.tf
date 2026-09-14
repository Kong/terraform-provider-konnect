resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - consumer-group-member"
  name         = "tf-test-aigw-consumer-group-member"
}

resource "konnect_ai_gateway_consumer" "my_aigatewayconsumer" {
  provider     = konnect-beta
  display_name = "TF Test Consumer"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-consumer"
  type         = "api-key"
}

resource "konnect_ai_gateway_consumer_group" "my_aigatewayconsumergroup" {
  provider     = konnect-beta
  display_name = "TF Test Consumer Group"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-consumers"

  # consumer and consumer_group are both direct children of the gateway; this single
  # local edge serializes their deletes so this fixture stays at 1 concurrent delete.
  depends_on = [konnect_ai_gateway_consumer.my_aigatewayconsumer]
}

resource "konnect_ai_gateway_consumer_group_member" "my_aigatewayconsumergroupmember" {
  provider          = konnect-beta
  consumer_id       = konnect_ai_gateway_consumer.my_aigatewayconsumer.id
  consumer_group_id = konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup.id
  gateway_id        = konnect_ai_gateway.my_aigateway.id
}
