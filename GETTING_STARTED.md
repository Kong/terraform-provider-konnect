# Getting Started

This provider is available on the [Terraform registry](https://registry.terraform.io/providers/Kong/konnect/latest).

## Sample manifest

Place the following content in `main.tf`, set your `personal_access_token` then run `terraform apply`.

```hcl
# Configure the provider to use your Kong Konnect account
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
    }
  }
}

provider "konnect" {
  personal_access_token = "kpat_YOUR_PAT"
  server_url            = "https://us.api.konghq.com"
}

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

resource "konnect_gateway_route" "hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  strip_path = false

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

# Secure the service with a basic-auth plugin
resource "konnect_gateway_plugin_basic_auth" "basic_auth" {
  enabled          = true
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
  config = {
    hide_credentials = false
  }
}

# Create a consumer and a basic auth credential for that consumer
resource "konnect_gateway_consumer" "alice" {
  username         = "alice"
  custom_id        = "alice"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_basic_auth" "my_basicauth" {
  username = "alice-test"
  password = "demo"

  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
```
