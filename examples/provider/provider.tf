terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "3.6.0"
    }
  }
}

provider "konnect" {
  # Configuration options
}