resource "konnect_gateway_control_plane" "plugin_entitlement_enforcement_cp" {
  name         = "Terraform Control Plane For EntitlementEnforcement Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_entitlement_enforcement" "my_entitlement_enforcement" {
  enabled = true
  config = {
    api_token                   = "test-api-token-12345"
    entitlement_access_endpoint = "https://entitlements.example.com/v1/access"

    feature = {
      key = "premium-api-access"
    }

    redis = {
      host = "redis.example.com"
      port = 6379
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugin_entitlement_enforcement_cp.id
}
