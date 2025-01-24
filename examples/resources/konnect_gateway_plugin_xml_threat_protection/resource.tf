resource "konnect_gateway_plugin_xml_threat_protection" "my_gatewaypluginxmlthreatprotection" {
  config = {
    allow_dtd = true
    allowed_content_types = [
      "..."
    ]
    attribute             = 6
    bla_max_amplification = 6.05
    bla_threshold         = 1030
    buffer                = 7
    checked_content_types = [
      "..."
    ]
    comment         = 1
    document        = 4
    entity          = 1
    entityname      = 5
    entityproperty  = 7
    localname       = 4
    max_attributes  = 6
    max_children    = 7
    max_depth       = 8
    max_namespaces  = 6
    namespace_aware = false
    namespaceuri    = 0
    pidata          = 10
    pitarget        = 9
    prefix          = 3
    text            = 3
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