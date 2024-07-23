# Create a Portal
resource "konnect_portal" "my_portal" {
  name                      = "My New Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tfdemo.example.com"
  is_public                 = false
  rbac_enabled              = false
}

# API Product configuration
resource "konnect_api_product" "httpbin" {
  name        = "HTTPBin Product"
  description = "This product productizes the HTTPBin service"

  portal_ids = [
    konnect_portal.my_portal.id
  ]
}

resource "konnect_api_product_version" "httpbin_v1" {
  api_product_id = konnect_api_product.httpbin.id
  name           = "v1"
  gateway_service = {
    control_plane_id = konnect_gateway_control_plane.tfdemo.id
    id               = konnect_gateway_service.httpbin.id
  }
}

# Read OAS from a URL
data "http" "httpbin_oas" {
  url = "https://petstore3.swagger.io/api/v3/openapi.json"
}

resource "konnect_api_product_specification" "httpbin_v1_spec" {
  name                   = "sample.json"
  content                = base64encode(tostring(data.http.httpbin_oas.response_body))
  api_product_id         = konnect_api_product.httpbin.id
  api_product_version_id = konnect_api_product_version.httpbin_v1.id
}

# Define an authentication strategy to be used by the product version
resource "konnect_application_auth_strategy" "my_applicationauthstrategy" {
  key_auth = {
    name          = "my-application-auth-strategy"
    key_names     = ["apikey"]
    display_name  = "My Test Strategy"
    strategy_type = "key_auth"
    configs = {
      key_auth = {
        key_names = ["apikey"]
      }
    }
  }
}

# Assign the product version to a portal
resource "konnect_portal_product_version" "my_portalproductversion" {
  application_registration_enabled = true
  auto_approve_registration        = true
  deprecated                       = false
  publish_status                   = "published"

  portal_id          = konnect_portal.my_portal.id
  product_version_id = konnect_api_product_version.httpbin_v1.id
  auth_strategy_ids = [
    konnect_application_auth_strategy.my_applicationauthstrategy.id
  ]
}
