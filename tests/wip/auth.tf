terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "999.99.9" # This ensures that we only test with the locally built provider
    }
  }
}

provider "konnect" {
  server_url = "https://us.api.konghq.com"
}

