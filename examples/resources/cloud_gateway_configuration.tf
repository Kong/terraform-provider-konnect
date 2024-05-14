resource "konnect_cloud_gateway_configuration" "my_cloudgatewayconfiguration" {
  version           = "3.6"
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
}
