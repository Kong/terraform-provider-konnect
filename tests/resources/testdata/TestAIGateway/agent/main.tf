resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - agent"
  name         = "tf-test-aigw-agent"
}

resource "konnect_ai_gateway_agent" "my_aigatewayagent" {
  provider = konnect-beta
  access = {
    acls = {
      deny = [
        "groupb"
      ]
    }
  }
  config = {
    route = {
      hosts = [
        "foo.example.com"
      ]
    }
    url = "https://booking-agent.internal.kongair.com"
  }
  display_name = "Test TF Flight Booking Agent"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-flight-booking-agent"
  type         = "a2a"
}
