# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

# Configure a serverless gateway control plane
resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Serverless TF Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_SERVERLESS_V1"
  auth_type    = "pinned_client_certs"
  cloud_gateway = true
}

# Configure a serverless gateway configuration to create the dataplane
resource "konnect_cloud_gateway_configuration" "my_cloudgatewayconfiguration1" {
  control_plane_geo = "us"
  control_plane_id  = konnect_gateway_control_plane.tfdemo.id
  dataplane_groups = [
    {
      provider = "aws"
      region   = "us"
    }
  ]
  kind    = "serverless.v1"
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

resource "konnect_gateway_data_plane_client_certificate" "my_cert" {
  cert             = file("./tls.crt")
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Output the endpoint of the serverless gateway
output "scgw_endpoint" {
  value = konnect_gateway_control_plane.tfdemo.config.control_plane_endpoint
}
