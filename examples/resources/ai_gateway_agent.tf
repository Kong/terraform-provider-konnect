resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "My AI Gateway"
  name         = "my-ai-gateway"
}

resource "konnect_ai_gateway_agent" "my_aigatewayagent" {
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "flight-booking-agent"
  display_name = "Flight Booking Agent"
  type         = "a2a"
  config = {
    url = "https://booking-agent.internal.example.com"
    route = {
      hosts = ["agents.example.com"]
    }
  }
  access = {
    acls = {
      allow = ["gold-tier-consumers"]
    }
  }
}
