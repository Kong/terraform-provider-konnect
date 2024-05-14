resource "konnect_cloud_gateway_network" "my_cloudgatewaynetwork" {
  name   = "Terraform Network"
  region = "eu-west-1"
  availability_zones = [
    "euw1-az1",
    "euw1-az2",
    "euw1-az3"
  ]

  firewall = {
    allowed_cidr_blocks = [
      "0.0.0.0/0"
    ]
  }

  cidr_block      = "192.168.0.0/16"
  ddos_protection = false

  cloud_gateway_provider_account_id = data.konnect_cloud_gateway_provider_account_list.my_cloudgatewayprovideraccountlist.data[0].id
}
