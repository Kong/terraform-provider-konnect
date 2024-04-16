resource "konnect_api_product_version" "httpbin_v1" {
  api_product_id = konnect_api_product.httpbin.id
  name           = "v1"
  gateway_service = {
    control_plane_id = konnect_gateway_control_plane.tfdemo.id
    id               = konnect_gateway_service.httpbin.id
  }
}
