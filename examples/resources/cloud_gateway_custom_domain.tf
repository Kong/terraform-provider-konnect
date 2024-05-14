resource "konnect_cloud_gateway_custom_domain" "my_cloudgatewaycustomdomain" {
  control_plane_geo = "us"
  control_plane_id  = konnect_gateway_control_plane.my_gatewaycontrolplane.id
  domain            = "api.example.com"
}
