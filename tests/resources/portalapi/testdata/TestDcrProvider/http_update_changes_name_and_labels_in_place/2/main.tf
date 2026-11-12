resource "konnect_dcr_provider" "my_dcrprovider" {
  http = {
    issuer = "https://issuer.example.com"
    name   = "my-dcr-provider-renamed"
    dcr_config = {
      api_key      = "tfacctestapikey"
      dcr_base_url = "https://dcr.example.com"
    }
    labels = {
      team = "identity"
    }
  }
}
