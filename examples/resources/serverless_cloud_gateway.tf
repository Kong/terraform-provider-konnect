resource "konnect_serverless_cloud_gateway" "my_scgw" {
  control_plane = {
    id     = konnect_gateway_control_plane.tfdemo.id
    prefix = replace(replace(konnect_gateway_control_plane.tfdemo.config.control_plane_endpoint, "https://", ""), ".us.cp0.konghq.com", "")
    region = "us"
  }
  cluster_cert     = file("./tls.crt")
  cluster_cert_key = file("./tls.key")
}