resource "konnect_gateway_control_plane" "api_product_version_linked_cp" {
  name         = "Terraform Control Plane For Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_service" "my_httpbin_api_product_version_svc" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.api_product_version_linked_cp.id
}


resource "konnect_portal" "my_portal" {
  name                      = "APIProductVersionNullPortal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "api-product-version-null-test.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_api_product" "my_apiproduct" {
  description = "API product description"

  name = "API Product for version with null description"
  portal_ids = [
    konnect_portal.my_portal.id
  ]
}

resource "konnect_api_product_version" "my_null_api_product_version" {
  api_product_id = konnect_api_product.my_apiproduct.id
  name           = "v1"
  gateway_service = {
    control_plane_id = konnect_gateway_control_plane.api_product_version_linked_cp.id
    id               = konnect_gateway_service.my_httpbin_api_product_version_svc.id
  }
}
