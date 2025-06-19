resource "konnect_cloud_gateway_private_dns" "my_cloudgatewayprivatedns" {
  name       = "my private dns name"
  network_id = konnect_cloud_gateway_network.my_cloudgatewaynetwork.id
  private_dns_attachment_config = {
    aws_private_dns_resolver_attachment_config = {
      kind = "aws-outbound-resolver"
      dns_config = {
        "mysubdomain.mydomain.com" = {
          remote_dns_server_ip_addresses = [
            "10.0.0.1"
          ]
        },
      }
    }
  }
}