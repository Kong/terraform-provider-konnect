resource "konnect_gateway_config_store_secret" "my_gatewayconfigstoresecret" {
  config_store_id  = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  key              = "ConfigStoreSecretKey"
  value            = "...my_value..."
}