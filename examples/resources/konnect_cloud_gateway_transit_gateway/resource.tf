resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
  aws_transit_gateway = {
    cidr_blocks = [
      "..."
    ]
    dns_config = [
      {
        domain_proxy_list = [
          "..."
        ]
        remote_dns_server_ip_addresses = [
          "..."
        ]
      }
    ]
    name = "us-east-2 transit gateway"
    transit_gateway_attachment_config = {
      kind               = "aws-transit-gateway-attachment"
      ram_share_arn      = "...my_ram_share_arn..."
      transit_gateway_id = "...my_transit_gateway_id..."
    }
  }
  aws_vpc_peering_gateway = {
    cidr_blocks = [
      "..."
    ]
    dns_config = [
      {
        domain_proxy_list = [
          "..."
        ]
        remote_dns_server_ip_addresses = [
          "..."
        ]
      }
    ]
    name = "us-east-2 transit gateway"
    transit_gateway_attachment_config = {
      kind            = "aws-vpc-peering-attachment"
      peer_account_id = "...my_peer_account_id..."
      peer_vpc_id     = "...my_peer_vpc_id..."
      peer_vpc_region = "...my_peer_vpc_region..."
    }
  }
  azure_transit_gateway = {
    dns_config = [
      {
        domain_proxy_list = [
          "..."
        ]
        remote_dns_server_ip_addresses = [
          "..."
        ]
      }
    ]
    name = "us-east-2 transit gateway"
    transit_gateway_attachment_config = {
      kind                = "azure-vnet-peering-attachment"
      resource_group_name = "...my_resource_group_name..."
      subscription_id     = "...my_subscription_id..."
      tenant_id           = "...my_tenant_id..."
      vnet_name           = "...my_vnet_name..."
    }
  }
  network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
}