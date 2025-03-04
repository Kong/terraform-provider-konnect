resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
  aws_vpc_peering_gateway = {
    cidr_blocks = [
      "192.168.0.0/24",
    ]
    dns_config = [
      {
        domain_proxy_list = [
          "example.com"
        ]
        remote_dns_server_ip_addresses = [
          "192.168.1.1"
        ]
      }
    ]
    name = "Terraform vpc peering gateway"
    transit_gateway_attachment_config = {
      kind            = "aws-vpc-peering-attachment"
      peer_account_id = "...my_peer_account_id..."
      peer_vpc_id     = "...my_peer_vpc_id..."
      peer_vpc_region = "...my_peer_vpc_region..."
    }
  }
  network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
}