resource "konnect_cloud_gateway_network" "my_cloudgatewaynetwork" {
  availability_zones = [
    "use2-az1",
    "use2-az2",
    "use2-az3",
  ]
  cidr_block                        = "10.0.0.0/8"
  cloud_gateway_provider_account_id = "929b2449-c69f-44c4-b6ad-9ecec6f811ae"
  name                              = "us-east-2 network"
  region                            = "us-east-2"
}