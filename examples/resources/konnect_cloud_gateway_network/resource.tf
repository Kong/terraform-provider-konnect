resource "konnect_cloud_gateway_network" "my_cloudgatewaynetwork" {
  availability_zones = [
    "...",
  ]
  cidr_block                        = "10.0.0.0/8"
  cloud_gateway_provider_account_id = "929b2449-c69f-44c4-b6ad-9ecec6f811ae"
  ddos_protection                   = false
  name                              = "us-east-2 network"
  network_id                        = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
  region                            = "us-east-2"
}