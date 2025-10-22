resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI GCP Model Armor Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_gcp_model_armor" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    project_id                = "my-project-id"
    location_id               = "location_id"
    template_id               = "template_id"
    guarding_mode             = "BOTH"
    gcp_use_service_account   = true
    gcp_service_account_json  = "{}"
    reveal_failure_categories = true
    request_failure_message   = "Your request was blocked by content policies."
    response_failure_message  = "The model's response was filtered for safety."
    timeout                   = 15000
    response_buffer_size      = 4096
    text_source               = "last_message"
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
