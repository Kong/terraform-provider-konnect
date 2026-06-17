terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "3.18.1"
    }
  }
}

provider "konnect" {
  server_url = "..." # Optional - can use KONNECT_SERVER_URL environment variable
}