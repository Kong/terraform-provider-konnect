resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AI Gateway"
  name         = "tf-test-ai-gateway"
}
