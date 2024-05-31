# Fetch the list of portals so that we can refer
# to the Portal ID for publishing
data "konnect_portal_list" "my_portallist" {
  page_number = 1
  page_size   = 1
}

# API Product configuration
resource "konnect_api_product" "httpbin" {
  name        = "HTTPBin Product"
  description = "This product productizes the HTTPBin service"

  # Which portals to publish the API product to
  # Konnect only supports a single portal at the moment
  # so we can rely on the first portal in the list
  portal_ids = [
    data.konnect_portal_list.my_portallist.data[0].id
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

  portal_id          = data.konnect_portal_list.my_portallist.data[0].id
  product_version_id = konnect_api_product_version.httpbin_v1.id
  auth_strategy_ids = [
    konnect_application_auth_strategy.my_applicationauthstrategy.id
  ]
}
