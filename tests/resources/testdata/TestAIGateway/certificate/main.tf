terraform {
  required_providers {
    tls = {
      source  = "hashicorp/tls"
      version = "~> 4.0"
    }
  }
}

resource "tls_private_key" "cert" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "cert" {
  private_key_pem = tls_private_key.cert.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "Test Organization"
  }

  validity_period_hours = 8760
  early_renewal_hours   = 720

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - certificate"
  name         = "tf-test-aigw-certificate"
}

resource "konnect_ai_gateway_certificate" "my_aigatewaycertificate" {
  provider = konnect-beta
  cert     = tls_self_signed_cert.cert.cert_pem
  key      = tls_private_key.cert.private_key_pem

  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-certificate"
}
