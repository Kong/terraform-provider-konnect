resource "konnect_gateway_partial" "my_gatewaypartial" {
  config = jsonencode({
				"port": 6379,
			})
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  name             = "my_tf_redis_ee_partial"
  tags = [
    "tf-test"
  ]
  type       = "redis-ee"
}