resource "konnect_cloud_gateway_custom_domain" "my_cloudgatewaycustomdomain" {
  control_plane_geo = "in"
  control_plane_id  = "0949471e-b759-45ba-87ab-ee63fb781388"
  domain            = "example.com"
}