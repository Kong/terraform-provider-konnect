terraform {
  required_providers {
    tls = {
      source  = "hashicorp/tls"
      version = "~> 4.0"
    }
  }
}

resource "tls_private_key" "ca" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "ca" {
  private_key_pem      = tls_private_key.ca.private_key_pem
  is_ca_certificate    = true

  subject {
    common_name  = "Test CA"
    organization = "Test Organization"
  }

  validity_period_hours = 8760
  early_renewal_hours   = 720

  allowed_uses = [
    "cert_signing",
  ]
}

resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - ca-certificate"
  name         = "tf-test-aigw-ca-cert"
}

resource "konnect_ai_gateway_ca_certificate" "my_aigatewaycacertificate" {
  provider = konnect-beta
  cert     = trimspace(tls_self_signed_cert.ca.cert_pem)

  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-ca-certificate"
}
