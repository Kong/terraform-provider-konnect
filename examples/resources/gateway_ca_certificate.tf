resource "konnect_gateway_ca_certificate" "my_ca_certificate" {
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
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
