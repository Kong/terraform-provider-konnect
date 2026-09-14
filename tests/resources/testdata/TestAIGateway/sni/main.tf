terraform {
  required_providers {
    tls = {
      source  = "hashicorp/tls"
      version = "~> 4.0"
    }
  }
}

resource "tls_private_key" "sni_cert" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "sni_cert" {
  private_key_pem = tls_private_key.sni_cert.private_key_pem

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
  display_name = "TF Test AIGW - sni"
  name         = "tf-test-aigw-sni"
}

resource "konnect_ai_gateway_certificate" "my_aigatewaycertificate" {
  provider = konnect-beta
  cert     = tls_self_signed_cert.sni_cert.cert_pem
  key      = tls_private_key.sni_cert.private_key_pem

  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-cert-for-sni"
}

resource "konnect_ai_gateway_sni" "my_aigatewaysni" {
  provider     = konnect-beta
  certificate  = konnect_ai_gateway_certificate.my_aigatewaycertificate.name
  display_name = "TF Test SNI"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  hostname     = "example.com"
  name         = "tf-test-sni"
}
