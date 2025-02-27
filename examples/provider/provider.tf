terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "2.3.0"
    }
  }
}

provider "konnect" {
  # Configuration options
}