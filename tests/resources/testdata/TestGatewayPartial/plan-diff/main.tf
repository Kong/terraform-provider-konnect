resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_partial" "my_gatewaypartial" {
  config = jsonencode({
				"cluster_max_redirections": 5,
				"cluster_nodes": null,
				"connect_timeout": 2000,
				"connection_is_proxied": false,
				"database": 0,
				"host": "127.0.0.1",
				"keepalive_backlog": null,
				"keepalive_pool_size": 256,
				"password": null,
				"port": 6379,
				"read_timeout": 3000,
				"send_timeout": 2004,
				"sentinel_master": null,
				"sentinel_nodes": null,
				"sentinel_password": null,
				"sentinel_role": null,
				"sentinel_username": null,
				"server_name": null,
				"ssl": false,
				"ssl_verify": false,
				"username": null
			})
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  name             = "my_tf_redis_ee_partial"
  tags = [
    "tf-test"
  ]
  type       = "redis-ee"
}