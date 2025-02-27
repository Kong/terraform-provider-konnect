resource "konnect_api_product" "my_apiproduct" {
  description = "Text describing the API product"
  labels = {
    key = "value",
  }
  name = "API Product"
  portal_ids = [
  ]
  public_labels = {
    key = "value",
  }
}