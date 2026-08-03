resource "konnect_dcr_provider" "my_dcrprovider" {
  kong_identity = {
    dcr_config = {}
    issuer     = "https://issuer.example.com"
    name       = "my-dcr-provider-renamed"
    labels = {
      team = "identity"
    }
  }
}
