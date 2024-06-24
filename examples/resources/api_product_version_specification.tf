resource "konnect_api_product_specification" "httpbin_v1_spec" {
  name                   = "httpbin.yaml"
  content                = base64encode(file("./httpbin.yaml"))
  api_product_id         = konnect_api_product.httpbin.id
  api_product_version_id = konnect_api_product_version.httpbin_v1.id
}