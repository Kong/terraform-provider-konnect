resource "konnect_gateway_plugin_ai_lakera_guard" "my_gatewaypluginailakeraguard" {
  config = {
    api_key                   = "...my_api_key..."
    guarding_mode             = "INPUT"
    lakera_service_url        = "...my_lakera_service_url..."
    project_id                = "...my_project_id..."
    request_failure_message   = "...my_request_failure_message..."
    response_buffer_size      = 9.27
    response_failure_message  = "...my_response_failure_message..."
    reveal_failure_categories = true
    stop_on_error             = true
    text_source               = "last_message"
    timeout                   = 4.57
    verify_ssl                = false
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 5
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
    "grpc"
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
  updated_at = 6
}