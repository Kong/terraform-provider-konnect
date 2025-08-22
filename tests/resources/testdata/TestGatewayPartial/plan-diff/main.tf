resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}
resource "konnect_gateway_partial" "my_gatewaypartial_ce" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
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
    name       = "my_tf_redis_ce_partial"
    tags = [
      "redisce"
    ]
  }
}


resource "konnect_gateway_partial" "my_gatewaypartial_ee" {
  redis_ee = {
    name = "my_tf_redis_ee_partial"
    config = {
      username = "rediseeusername"
      connect_timeout = 2000
      connection_is_proxied = false
      database = 0
      host = "127.0.0.1"
      keepalive_backlog = 0
      keepalive_pool_size = 256
      password = "rediseepassword"
      port = 6379
      read_timeout = 2000
      send_timeout = 2000
      server_name = "redis.example.com"
      ssl = true
      ssl_verify = false
    }
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}