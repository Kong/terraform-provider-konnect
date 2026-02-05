resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
  aws_resource_endpoint_gateway = {
    dns_config = [
      {
        domain_proxy_list = [
          "foobar.com",
        ]
        remote_dns_server_ip_addresses = [
          "10.0.0.2",
        ]
      }
    ]
    name = "us-east-2 transit gateway"
    transit_gateway_attachment_config = {
      kind          = "aws-resource-endpoint-attachment"
      ram_share_arn = "...my_ram_share_arn..."
      resource_config = [
        {
          domain_name        = "...my_domain_name..."
          resource_config_id = "...my_resource_config_id..."
        }
      ]
    }
  }
  aws_transit_gateway = {
    cidr_blocks = [
      "10.0.0.0/8",
      "100.64.0.0/10",
      "172.16.0.0/12",
    ]
    dns_config = [
      {
        domain_proxy_list = [
          "foobar.com",
        ]
        remote_dns_server_ip_addresses = [
          "10.0.0.2",
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
      "10.0.0.0/8",
      "100.64.0.0/10",
      "172.16.0.0/12",
    ]
    dns_config = [
      {
        domain_proxy_list = [
          "foobar.com",
        ]
        remote_dns_server_ip_addresses = [
          "10.0.0.2",
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
          "foobar.com",
        ]
        remote_dns_server_ip_addresses = [
          "10.0.0.2",
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
  azure_vhub_peering_gateway = {
    dns_config = [
      {
        domain_proxy_list = [
          "foobar.com",
        ]
        remote_dns_server_ip_addresses = [
          "10.0.0.2",
        ]
      }
    ]
    name = "us-east-2 transit gateway"
  }
  gcp_vpc_peering_transit_gateway = {
    dns_config = [
      {
        domain_proxy_list = [
          "foobar.com",
        ]
        remote_dns_server_ip_addresses = [
          "10.0.0.2",
        ]
      }
    ]
    name = "us-east-2 transit gateway"
    transit_gateway_attachment_config = {
      kind            = "gcp-vpc-peering-attachment"
      peer_project_id = "...my_peer_project_id..."
      peer_vpc_name   = "...my_peer_vpc_name..."
    }
  }
  network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
}