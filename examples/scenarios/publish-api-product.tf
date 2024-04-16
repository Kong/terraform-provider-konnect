# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

# Create a new Control Plane
resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_HYBRID"
  auth_type    = "pinned_client_certs"

  proxy_urls = [
    {
      host     = "example.com",
      port     = 443,
      protocol = "https"
    }
  ]
}

# Configure a service and a route that we can use to test
resource "konnect_gateway_service" "httpbin" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_route" "anything" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  strip_path = false

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

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
