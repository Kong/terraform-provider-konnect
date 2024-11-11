# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

# Create a new Control Plane
resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
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

resource "konnect_portal" "my_portal" {
  name                      = "My New Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  is_public                 = false
  rbac_enabled              = false
}

# API Product configuration
resource "konnect_api_product" "httpbin" {
  name        = "HTTPBin Product"
  description = "This product productizes the HTTPBin service"

  # Which portals to publish the API product to
  # Konnect only supports a single portal at the moment
  # so we can rely on the first portal in the list
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

resource "konnect_api_product_specification" "httpbin_v1_spec" {
  name                   = "httpbin.yaml"
  content                = base64encode(file("./httpbin.yaml"))
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

resource "konnect_api_product_document" "my_apiproductdocument" {
  title          = "How to create a document in Konnect DocumentHub"
  content        = "## How To"
  slug           = "create-a-document"
  status         = "published"
  api_product_id = konnect_api_product.httpbin.id
}

resource "konnect_api_product_document" "my_child_apiproductdocument" {
  title              = "Child Page"
  content            = "## Child Page"
  slug               = "child-page"
  status             = "published"
  api_product_id     = konnect_api_product.httpbin.id
  parent_document_id = konnect_api_product_document.my_apiproductdocument.id
}
