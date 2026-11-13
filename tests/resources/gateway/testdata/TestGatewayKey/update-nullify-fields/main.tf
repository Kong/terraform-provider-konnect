resource "konnect_gateway_control_plane" "gateway_key_cp" {
  name         = "Terraform Control Plane For Gateway Key"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_key_set" "my_keyset" {
  name             = "demo-keyset"
  control_plane_id = konnect_gateway_control_plane.gateway_key_cp.id
}

resource "konnect_gateway_key" "my_nullable_key" {
  name = "my-test-nullify-jwk"
  kid  = "mykeyid"
  set = {
    id = konnect_gateway_key_set.my_keyset.id
  }
  control_plane_id = konnect_gateway_control_plane.gateway_key_cp.id
  pem              = {
    private_key = <<-EOT
          -----BEGIN RSA PRIVATE KEY-----
          MIICXAIBAAKBgG79jfVbRQHMEV/qaskL6WVeClxYD17q3hWBr6xLpVENedhky8ps
          dXVoqLkbWMW+uYHB9vYgXb9yI+Mu9EYQ4y6lmi18nPX3qkYsNHCkItcjAnVSHkeC
          h+S7lsHvnTALu/0y/emldJiWqh38wKnlZxE+2QcdELKTomExyeuI6Xp9AgMBAAEC
          gYACCm1yxnPVXaAwKEpAWhS3hWwhWIkn0J+2u4S0YUuv2tSTsv7AQRBD0WHQzXzr
          Xd9hUGSvD9dJbtrUEYfyRds3eR6j5QOUd9ITgya4Bsyk9G2YLoSa0xRRIO1EG4gw
          WU8LsePrvjtBDvN1TipmCy0LpS4W2A26D/AzGnIdVi5SbQJBALr0q+dV4xGnLU8N
          6DHeU+fwf9sc5SYkOQ0xoPBW8Tp2oaYRTguyCqBrDUc5TySxHFTR1hW8a7FVSwE2
          xgNmDdMCQQCX+uhE3xWTQQ+uJn7GrfBp541XZV/g7Ur3Lk9akGpDe/S9O0qgIk6b
          2JxeECBrgAhpV0eEjWacIzm+caBA9xRvAkEAtw/VM5QZ37M+3mXTsuDsb/RCfZ7x
          kkNtycmKuUuosYJwDlhrf3A6j5HDUrZ/FIKIJ1XAt0/kc6P86ZdklPddGQJAEu4S
          rjlnM49AB74Nvmt97YME7OTYm2iTFNS015/zTHKNGmDfO7DqP6ksWN8DWsB4y74u
          T0lZhYsxrxHyFwVviwJBAJFjgPNOdiiLGCZH+jTxn13AYeK+FRCrpLEQUzyNlcHB
          DFgOK+hv6mppQlmApBsC82rTSsBjl1l7R1qHDFQcZR8=
          -----END RSA PRIVATE KEY-----
      EOT
    public_key  = <<-EOT
          -----BEGIN PUBLIC KEY-----
          MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgG79jfVbRQHMEV/qaskL6WVeClxY
          D17q3hWBr6xLpVENedhky8psdXVoqLkbWMW+uYHB9vYgXb9yI+Mu9EYQ4y6lmi18
          nPX3qkYsNHCkItcjAnVSHkeCh+S7lsHvnTALu/0y/emldJiWqh38wKnlZxE+2Qcd
          ELKTomExyeuI6Xp9AgMBAAE=
          -----END PUBLIC KEY-----
      EOT
  }
}
