resource "konnect_gateway_control_plane" "my_gatewaycontrolplane" {
  name          = "CGW Test"
  cluster_type  = "CLUSTER_TYPE_CONTROL_PLANE"
  cloud_gateway = true
  description   = "A test control plane for exploration."
  auth_type     = "pinned_client_certs"
  proxy_urls    = []
}

data "konnect_cloud_gateway_provider_account_list" "my_cloudgatewayprovideraccountlist" {
  page_number = 1
  page_size   = 1
}

resource "konnect_cloud_gateway_network" "my_cloudgatewaynetwork" {
  name   = "Terraform Network"
  region = "eu-west-1"
  availability_zones = [
    "euw1-az1",
    "euw1-az2",
    "euw1-az3"
  ]

  cidr_block      = "192.168.0.0/16"

  cloud_gateway_provider_account_id = data.konnect_cloud_gateway_provider_account_list.my_cloudgatewayprovideraccountlist.data[0].id
}

resource "konnect_cloud_gateway_configuration" "my_cloudgatewayconfiguration" {
  api_access        = "private+public"
  control_plane_geo = "us"
  control_plane_id  = konnect_gateway_control_plane.my_gatewaycontrolplane.id
  dataplane_groups = [
    {
      provider = "aws"
      region   = "eu-west-1"
      autoscale = {
        configuration_data_plane_group_autoscale_autopilot = {
          kind     = "autopilot"
          base_rps = 10
          max_rps  = 100
        }

        #configuration_data_plane_group_autoscale_static = {
        #  kind                = "static"
        #  instance_type       = "small"
        #  requested_instances = 1
        #}
      }
      cloud_gateway_network_id = konnect_cloud_gateway_network.my_cloudgatewaynetwork.id
    },
  ]
  version = "3.6"
}

resource "konnect_cloud_gateway_custom_domain" "my_cloudgatewaycustomdomain" {
  control_plane_geo = "us"
  control_plane_id  = konnect_gateway_control_plane.my_gatewaycontrolplane.id
  domain            = "api.example.com"
}

#resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
#  name = "Terraform Transit Gateway"
#
#  cidr_blocks = [
#    "192.168.0.0/24",
#  ]
#
#  transit_gateway_attachment_config = {
#    aws_transit_gateway_attachment_config = {
#      kind               = "aws-transit-gateway-attachment"
#      ram_share_arn      = "arn:aws:ram:eu-west-1:111122223333:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12"
#      transit_gateway_id = "arn:aws:ec2:eu-west-1:111122223333:transit-gateway/tgw-0262a0e521EXAMPLE"
#    }
#  }
#
#  network_id = konnect_cloud_gateway_network.my_cloudgatewaynetwork.id
#}
