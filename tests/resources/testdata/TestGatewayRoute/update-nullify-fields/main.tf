resource "konnect_gateway_control_plane" "routecp" {
  name         = "Terraform Control Plane For Route"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_certificate" "my_sni_certificate" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIIBxjCCAUygAwIBAgIUX9TaLbWF76yQc8IGR+YRbeiDlHkwCgYIKoZIzj0EAwIw
GjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMB4XDTI0MDMwMTE0MzkxNloXDTI3
MDMwMTE0MzkxNlowGjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMHYwEAYHKoZI
zj0CAQYFK4EEACIDYgAEcMndCotXzeZ9vGAMfDfZ7UxUuP5bcIrwwUOI8YlpMdvB
12HvjtS7O0/ONr3fBeCWagRuitPEqd4b3EJuD8kuFUMt+2A09N6KY1YDJWgKHei7
rzKgrefzVt11XgBiDsUBo1MwUTAdBgNVHQ4EFgQUIrdAC8p02h60GZW0Jlh2Vcg/
WeMwHwYDVR0jBBgwFoAUIrdAC8p02h60GZW0Jlh2Vcg/WeMwDwYDVR0TAQH/BAUw
AwEB/zAKBggqhkjOPQQDAgNoADBlAjBYb+yQf33sItlmsONLc41Agtx73FMEN7Lf
WA85OtlkMie1N1x0mj08pzS/Xc1VONwCMQDN9sBn3Kody0gse+EXYSuPPj1oo9jm
FB9/xrpz35YpDATvuyhH8xwSJ4xMuxQiduc=
-----END CERTIFICATE-----
EOF
  key              = <<EOF
-----BEGIN PRIVATE KEY-----
MIG2AgEAMBAGByqGSM49AgEGBSuBBAAiBIGeMIGbAgEBBDDLuRX+uzSbstvLWsQr
WwuGK4AdjLU/tN9A/fn03gxNvppKw++SBtnLyB+9YZ29YA+hZANiAARwyd0Ki1fN
5n28YAx8N9ntTFS4/ltwivDBQ4jxiWkx28HXYe+O1Ls7T842vd8F4JZqBG6K08Sp
3hvcQm4PyS4VQy37YDT03opjVgMlaAod6LuvMqCt5/NW3XVeAGIOxQE=
-----END PRIVATE KEY-----
EOF
  control_plane_id = konnect_gateway_control_plane.routecp.id
}

resource "konnect_gateway_sni" "my_route_sni" {
  name = "example.com"
  certificate = {
    id = konnect_gateway_certificate.my_sni_certificate.id
  }
  control_plane_id = konnect_gateway_control_plane.routecp.id
}

resource "konnect_gateway_route" "my_nullify_route" {
  methods = ["GET"]
  name    = "my-route-name"
  paths   = ["/anything"]
  hosts = [
    "httpbin.org"
  ]
  snis = [
    konnect_gateway_sni.my_route_sni.id
  ]

  control_plane_id = konnect_gateway_control_plane.routecp.id
}

