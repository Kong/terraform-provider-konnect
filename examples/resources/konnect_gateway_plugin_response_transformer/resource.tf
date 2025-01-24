resource "konnect_gateway_plugin_response_transformer" "my_gatewaypluginresponsetransformer" {
  config = {
    add = {
      headers = [
        "..."
      ]
      json = [
        "..."
      ]
      json_types = [
        "number"
      ]
    }
    append = {
      headers = [
        "..."
      ]
      json = [
        "..."
      ]
      json_types = [
        "number"
      ]
    }
    remove = {
      headers = [
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
      json = [
        "..."
      ]
    }
    replace = {
      headers = [
        "..."
      ]
      json = [
        "..."
      ]
      json_types = [
        "string"
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
  protocols = [
    "wss"
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