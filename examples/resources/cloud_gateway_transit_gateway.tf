resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
  name = "Terraform Transit Gateway"

  cidr_blocks = [
    "192.168.0.0/24",
  ]

  transit_gateway_attachment_config = {
    aws_transit_gateway_attachment_config = {
      kind               = "aws-transit-gateway-attachment"
      ram_share_arn      = "arn:aws:ram:eu-west-1:111122223333:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12"
      transit_gateway_id = "arn:aws:ec2:eu-west-1:111122223333:transit-gateway/tgw-0262a0e521EXAMPLE"
    }
  }

  network_id = konnect_cloud_gateway_network.my_cloudgatewaynetwork.id
}
