resource "konnect_event_gateway_listener_policy_tls_server" "my_eventgatewaylistenerpolicytlsserver" {
  config = {
    allow_plaintext = true
    certificates = [
      {
        certificate = "...my_certificate..."
        key         = "$${vault.env['MY_ENV_VAR']}"
      }
    ]
    versions = {
      max = "TLSv1.3"
      min = "TLSv1.3"
    }
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  listener_id = "f7d7b9be-5608-44c3-8f6a-46e055797c31"
  name        = "...my_name..."
}