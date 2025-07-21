resource "konnect_gateway_control_plane" "vaultcp" {
  name         = "Terraform Control Plane For Vault"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_vault" "my_vault" {
  name   = "env"
  prefix = "my-env-vault"
  config = jsonencode({
    prefix = "bbbb"
    base64_decode = false
  })
  control_plane_id = konnect_gateway_control_plane.vaultcp.id
}
