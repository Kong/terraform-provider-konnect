resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For Kafka Consume Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_kafka_consume" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    bootstrap_servers = [
      {
        host = "bootstrap_server_host"
        port = 9092
    }]
    topics = [
      {
        name = "kafka_topic"
    }]
    mode = "server-sent-events"
    message_by_lua_functions = [<<EOF
return function(message)
  local m = type(message) == "table" and message or { value = message }
  if m.value == nil or (type(m.value)=="string" and #m.value==0) then
    return nil
  end
  return m
end
EOF
    ]
  }
  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
