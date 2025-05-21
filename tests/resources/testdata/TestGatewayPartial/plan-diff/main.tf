resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}
resource "konnect_gateway_partial" "my_gatewaypartial" {
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