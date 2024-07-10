resource "konnect_api_product" "my_apiproduct" {
  description = "Text describing the API product"
  name        = "API Product"
  portal_ids = [
    "25a2624c-49fc-4764-99e1-224ed819f200",
  ]
}