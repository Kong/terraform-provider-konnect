resource "konnect_gateway_partial" "my_gatewaypartialce" {
  control_plane_id = konnect_gateway_control_plane.tfdemo2.id
  redis_ce = {
    config = {
      database    = 5
      host        = "localhost"
      password    = "rediscepass"
      port        = 6379
      server_name = "redisceservername"
      ssl         = true
      ssl_verify  = true
      timeout     = 1000
      username    = "redisceusername"
    }
    name       = "redis-ce-partial"
    tags = [
      "redisce"
    ]
  }
}