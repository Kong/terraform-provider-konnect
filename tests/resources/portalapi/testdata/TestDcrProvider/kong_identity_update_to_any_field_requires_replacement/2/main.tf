resource "konnect_dcr_provider" "my_dcrprovider" {
  kong_identity = {
    dcr_config = {}
    issuer     = "https://issuer-changed.example.com"
    name       = "my-dcr-provider"
    labels = {
      team = "platform"
    }
  }
}
