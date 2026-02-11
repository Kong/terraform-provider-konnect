terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
    }
  }
}


provider "konnect" {
  server_url            = "https://us.api.konghq.com"
  konnect_access_token = "$$TOKEN"// your token
}


# resource "konnect_gateway_control_plane" "my_gatewaycontrolplane" {
#   cluster_type  = "CLUSTER_TYPE_CONTROL_PLANE"
#   description   = "A test control plane for exploration."
#   name = "Test Control Plane"
# }
#
# resource "konnect_gateway_service" "example_service" {
#   name     = "demo-service"
#   protocol = "http"
#   host     = "httpbin.org"
#   port     = 80
#   control_plane_id = "c41eb79d-43d1-4d46-b597-a866aa2e9b38"
# }
#
# resource "konnect_gateway_route" "anything" {
#   methods = ["GET"]
#   name    = "Anything"
#   paths   = ["/anything"]
#
#   strip_path = false
#
#   control_plane_id = "c41eb79d-43d1-4d46-b597-a866aa2e9b38"
#   service = {
#     id = "cb0526ad-f364-4665-a246-ca737bd0de67"
#   }
# }
#
#
# # Apply a rate limit of 5 requests per minute
# resource "konnect_gateway_plugin_rate_limiting" "my_rate_limiting_plugin" {
#   enabled = true
#   config = {
#     minute = 5
#     policy = "local"
#   }
#
#   protocols        = ["http", "https"]
#   control_plane_id = "c41eb79d-43d1-4d46-b597-a866aa2e9b38"
#   route = {
#     id = "cd6d8866-f669-4fb1-b7ef-faeea9bfadb5"
#   }
# }


resource "konnect_event_gateway" "my_event_gateway" {
  provider = konnect
  name = "my_test_event_gateway"
}


resource "konnect_event_gateway_backend_cluster" "my_event_gateway_backend_cluster" {
  provider = konnect
name = "my_test_event_gateway_backend_cluster"
authentication = {
anonymous = {}
}
bootstrap_servers = [
"10.0.0.1:8080"
]
tls = {
enabled = false
insecure_skip_verify = true
}
  gateway_id = konnect_event_gateway.my_event_gateway.id
}

resource "konnect_event_gateway_listener" "test_listener" {
  provider = konnect
  name        = "test-listener"
  description = "Test listener configuration"
  addresses   = ["0.0.0.0"]
  ports       = ["9092"]
  gateway_id  = konnect_event_gateway.my_event_gateway.id
}


resource "konnect_event_gateway_virtual_cluster" "my_event_gateway_virtual_cluster" {
  provider = konnect
name      = "my_test_event_gateway_virtual_cluster"
dns_label = "vcluster-1"
acl_mode  = "passthrough"
authentication = [
{
anonymous = {}
}
]
destination = {
id = konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster.id
}
  gateway_id = konnect_event_gateway.my_event_gateway.id
}

resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "tf_test_egw_forward_policy" {
  provider = konnect
  name        = "tf_test_egw_forward_policy"
  description = "Test forward to virtual cluster policy"
  enabled     = true

  config = {
    port_mapping = {
      advertised_host = "broker.example.com"
      bootstrap_port  = "at_start"
      min_broker_id   = 0
      destination = {
        virtual_cluster_reference_by_name = {
          name = konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster.name
        }
      }
    }
  }
  listener_id = konnect_event_gateway_listener.test_listener.id
  gateway_id                = konnect_event_gateway.my_event_gateway.id
}
