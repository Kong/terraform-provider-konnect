terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "2.4.0"
    }
  }
}

provider "konnect" {
  # Configuration options
}