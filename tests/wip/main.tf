
data "konnect_gateway_control_plane_list" "my_gatewaycontrolplanelist" {
  filter = {
    name = {
      eq = "demo123"
    }
  }
  page_number = 1
  page_size   = 1
}

output "gateway_control_plane_list" {
  value = data.konnect_gateway_control_plane_list.my_gatewaycontrolplanelist.data[0]
}
