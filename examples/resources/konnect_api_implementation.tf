resource "konnect_api_implementation" "my_apiimplementation" {
  api_id = konnect_api.my_api.id
  service_reference = {
    service = {
      control_plane_id = konnect_gateway_control_plane.tfdemo.id
      id               = konnect_gateway_service.httpbin.id
    }
  }
}
