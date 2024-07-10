resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
  cidr_blocks = [
    "...",
  ]
  name       = "us-east-2 transit gateway"
  network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
  transit_gateway_attachment_config = {
    aws_transit_gateway_attachment_config = {
      kind               = "aws-transit-gateway-attachment"
      ram_share_arn      = "...my_ram_share_arn..."
      transit_gateway_id = "...my_transit_gateway_id..."
    }
  }
  transit_gateway_id = "0850820b-d153-4a2a-b9be-7d2204779139"
}