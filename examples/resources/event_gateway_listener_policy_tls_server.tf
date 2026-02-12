resource "konnect_event_gateway_listener_policy_tls_server" "tls" {
  gateway_id  = konnect_event_gateway.demo.id
  listener_id = konnect_event_gateway_listener.listener.id

  config = {
    certificates = [
      { certificate = "-----BEGIN CERTIFICATE-----\n...PEM...\n-----END CERTIFICATE-----", key = "REPLACE_KEY" }
    ]
  }
}
