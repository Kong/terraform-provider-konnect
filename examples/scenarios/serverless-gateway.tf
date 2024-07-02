# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Serverless TF Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_SERVERLESS"
  auth_type    = "pinned_client_certs"
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

resource "konnect_serverless_cloud_gateway" "my_scgw" {
  control_plane = {
    id     = konnect_gateway_control_plane.tfdemo.id
    prefix = replace(replace(konnect_gateway_control_plane.tfdemo.config.control_plane_endpoint, "https://", ""), ".us.cp0.konghq.com", "")
    region = "us"
  }
  cluster_cert     = file("./tls.crt")
  cluster_cert_key = file("./tls.key")
}

output "scgw_endpoint" {
  value = konnect_serverless_cloud_gateway.my_scgw.gateway_endpoint
}
