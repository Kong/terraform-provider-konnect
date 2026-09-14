resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - data-plane-certificate"
  name         = "tf-test-aigw-dp-cert"
}

resource "konnect_ai_gateway_data_plane_certificate" "my_aigatewaydataplanecertificate" {
  provider = konnect-beta
  cert     = <<EOF
-----BEGIN CERTIFICATE-----
MIIB4TCCAYugAwIBAgIUAenxUyPjkSLCe2BQXoBMBacqgLowDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yNDEwMjgyMDA3NDlaFw0zNDEw
MjYyMDA3NDlaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwXDANBgkqhkiG9w0BAQEF
AANLADBIAkEAyzipjrbAaLO/yPg7lL1dLWzhqNdc3S4YNR7f1RG9whWhbsPE2z42
e6WGFf9hggP6xjG4qbU8jFVczpd1UPwGbQIDAQABo1MwUTAdBgNVHQ4EFgQUkPPB
ghj+iHOHAKJlC1gLbKT/ZHQwHwYDVR0jBBgwFoAUkPPBghj+iHOHAKJlC1gLbKT/
ZHQwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAANBALfy49GvA2ld+u+G
Koxa8kCt7uywoqu0hfbBfUT4HqmXPvsuhz8RinE5ltxId108vtDNlD/+bKl+N5Ub
qKjBs0k=
-----END CERTIFICATE-----
EOF

  gateway_id = konnect_ai_gateway.my_aigateway.id
  title      = "tf-test-dp-cert"
}
