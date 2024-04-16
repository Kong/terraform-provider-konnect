resource "konnect_api_product" "httpbin" {
  name        = "HTTPBin Product"
  description = "This is a description for my API Product"

  portal_ids = [
    data.konnect_portal_list.my_portallist.data[0].id
  ]
}
