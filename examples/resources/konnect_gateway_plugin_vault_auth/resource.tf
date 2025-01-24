resource "konnect_gateway_plugin_vault_auth" "my_gatewaypluginvaultauth" {
  config = {
    access_token_name = "...my_access_token_name..."
    anonymous         = "...my_anonymous..."
    hide_credentials  = false
    run_on_preflight  = false
    secret_token_name = "...my_secret_token_name..."
    tokens_in_body    = true
    vault             = "...my_vault..."
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = true
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  protocols = [
    "tcp"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}