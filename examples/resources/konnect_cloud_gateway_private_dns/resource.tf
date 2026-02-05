resource "konnect_cloud_gateway_private_dns" "my_cloudgatewayprivatedns" {
  name       = "us-east-2 private dns"
  network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
  private_dns_attachment_config = {
    aws_private_hosted_zone_attachment_config = {
      hosted_zone_id = "...my_hosted_zone_id..."
      kind           = "aws-private-hosted-zone-attachment"
    }
  }
}