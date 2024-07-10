resource "konnect_cloud_gateway_custom_domain" "my_cloudgatewaycustomdomain" {
  control_plane_geo = "us"
  control_plane_id  = "0949471e-b759-45ba-87ab-ee63fb781388"
  custom_domain_id  = "39ed3790-085d-4605-9627-f96d86aaf425"
  domain            = "example.com"
}