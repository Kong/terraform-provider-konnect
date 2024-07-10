resource "konnect_cloud_gateway_configuration" "my_cloudgatewayconfiguration" {
  api_access        = "private"
  configuration_id  = "edaf40f9-9fb0-4ffe-bb74-4e763a6bd471"
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
      }
      cloud_gateway_network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
      created_at               = "2022-11-04T20:10:06.927Z"
      egress_ip_addresses = [
        "...",
      ]
      id = "cbb8872a-1f83-4806-bf69-fdf0b4783c7e"
      private_ip_addresses = [
        "...",
      ]
      provider   = "aws"
      region     = "us-east-2"
      state      = "created"
      updated_at = "2022-11-04T20:10:06.927Z"
    },
  ]
  version = "3.2"
}