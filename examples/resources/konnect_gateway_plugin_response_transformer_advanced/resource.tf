resource "konnect_gateway_plugin_response_transformer_advanced" "my_gatewaypluginresponsetransformeradvanced" {
  config = {
    add = {
      headers = [
        "..."
      ]
      if_status = [
        "..."
      ]
      json = [
        "..."
      ]
      json_types = [
        "number"
      ]
    }
    allow = {
      json = [
        "..."
      ]
    }
    append = {
      headers = [
        "..."
      ]
      if_status = [
        "..."
      ]
      json = [
        "..."
      ]
      json_types = [
        "number"
      ]
    }
    dots_in_keys = true
    remove = {
      headers = [
        "..."
      ]
      if_status = [
        "..."
      ]
      json = [
        "..."
      ]
    }
    rename = {
      headers = [
        "..."
      ]
      if_status = [
        "..."
      ]
    }
    replace = {
      body = "...my_body..."
      headers = [
        "..."
      ]
      if_status = [
        "..."
      ]
      json = [
        "..."
      ]
      json_types = [
        "number"
      ]
    }
    transform = {
      functions = [
        "..."
      ]
      if_status = [
        "..."
      ]
      json = [
        "..."
      ]
    }
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
}