terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "2.2.0"
    }
  }
}

provider "konnect" {
  # Configuration options
}