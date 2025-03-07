resource "konnect_gateway_key" "my_gatewaykey" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id               = "...my_id..."
  jwk              = "...my_jwk..."
  kid              = "...my_kid..."
  name             = "...my_name..."
  pem = {
    private_key = "...my_private_key..."
    public_key  = "...my_public_key..."
  }
  set = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}