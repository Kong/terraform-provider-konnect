resource "konnect_gateway_control_plane" "my_cp1" {
  name         = "Terraform Test Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
  auth_type    = "pinned_client_certs"
}

resource "konnect_portal" "my_portal" {
  name                      = "My New Test Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tf-example-domain.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_api_product" "my_apiproduct" {
  description = "Text describing the API product"
  labels = {
    key = "value"
  }
  name = "API Product"
  portal_ids = [
    konnect_portal.my_portal.id
  ]
}

resource "konnect_gateway_service" "my_service" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.my_cp1.id
}

resource "konnect_api_product_version" "my_apiproductversion" {
  api_product_id = konnect_api_product.my_apiproduct.id
  gateway_service = {
    control_plane_id = konnect_gateway_control_plane.my_cp1.id
    id               = konnect_gateway_service.my_service.id
  }
  name = "v1"
}

resource "konnect_api_product_specification" "my_apiproductspecification" {
  api_product_id         = konnect_api_product.my_apiproduct.id
  api_product_version_id = konnect_api_product_version.my_apiproductversion.id
  content                =  "TXkgWUFNTCBvciBKU09OIGZvcm1hdHRlZCBPQVMgY29udGVudA=="
  name                   = "oas.yaml"
}
