resource "konnect_cloud_gateway_configuration" "my_cloudgatewayconfiguration" {
  api_access        = "public"
  control_plane_geo = "us"
  control_plane_id  = "0949471e-b759-45ba-87ab-ee63fb781388"
  dataplane_groups = [
    {
      autoscale = {
        configuration_data_plane_group_autoscale_autopilot = {
          base_rps = 1
          kind     = "autopilot"
          max_rps  = 1000
        }
        configuration_data_plane_group_autoscale_static = {
          instance_type       = "medium"
          kind                = "static"
          requested_instances = 3
        }
      }
      cloud_gateway_network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
      environment = [
        {
          name  = "KONG_LOG_LEVEL"
          value = "INFO"
        }
      ]
      provider = "aws"
      region   = "us-east-2"
    }
  ]
  version = "3.2"
}