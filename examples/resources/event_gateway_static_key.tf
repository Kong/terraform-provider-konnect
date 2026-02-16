resource "konnect_event_gateway_static_key" "statickey" {
  name  = "tf-test-static-key"
  value = "$${vault.env[\"MY_ENV_VAR\"]}"
  gateway_id = konnect_event_gateway.demo.id
}
