data "konnect_gateway_control_plane_list" "my_gatewaycontrolplanelist" {
  filter = {
    name = {
      eq       = "tfdemo"
    }
  }
  page_number = 1
  page_size   = 10
  sort        = "created_at desc"
}