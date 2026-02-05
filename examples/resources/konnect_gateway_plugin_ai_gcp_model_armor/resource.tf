resource "konnect_gateway_plugin_ai_gcp_model_armor" "my_gatewaypluginaigcpmodelarmor" {
  config = {
    enable_multi_language_detection = true
    gcp_service_account_json        = "...my_gcp_service_account_json..."
    gcp_use_service_account         = true
    guarding_mode                   = "INPUT"
    location_id                     = "...my_location_id..."
    project_id                      = "...my_project_id..."
    request_failure_message         = "...my_request_failure_message..."
    response_buffer_size            = 8.82
    response_failure_message        = "...my_response_failure_message..."
    reveal_failure_categories       = false
    source_language                 = "...my_source_language..."
    stop_on_error                   = false
    template_id                     = "...my_template_id..."
    text_source                     = "concatenate_all_content"
    timeout                         = 7.5
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 3
  enabled          = false
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpcs"
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
  updated_at = 4
}