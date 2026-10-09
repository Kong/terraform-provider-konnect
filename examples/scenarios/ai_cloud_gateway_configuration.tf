resource "konnect_ai_gateway" "my_aigateway" {
  name            = "my-ai-gateway"
  display_name    = "My AI Gateway"
  description     = "An AI Gateway for my organization."
  deployment_type = "managed"
}

resource "konnect_cloud_gateway_configuration" "my_ai_cloudgatewayconfiguration" {
  api_access        = "public"
  control_plane_geo = "us"
  control_plane_id  = konnect_ai_gateway.my_aigateway.id
  dataplane_groups = [
    {
      provider                 = "aws"
      region                   = "eu-west-1"
      cloud_gateway_network_id = konnect_cloud_gateway_network.my_cloudgatewaynetwork.id
    }
  ]
  kind = "dedicated.v0"
  type = "ai"
}