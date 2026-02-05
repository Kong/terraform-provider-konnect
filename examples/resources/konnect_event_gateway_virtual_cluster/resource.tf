resource "konnect_event_gateway_virtual_cluster" "my_eventgatewayvirtualcluster" {
  acl_mode = "enforce_on_gateway"
  authentication = [
    {
      sasl_plain = {
        mediation = "passthrough"
        principals = [
          {
            password = "$${vault.env['MY_ENV_VAR']}"
            username = "...my_username..."
          }
        ]
      }
    }
  ]
  description = "...my_description..."
  destination = {
    id = "759b5471-3de4-485c-b7d3-6e8cb8929d81"
  }
  dns_label  = "vcluster-1"
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  namespace = {
    additional = {
      consumer_groups = [
        {
          glob = {
            glob = "...my_glob..."
          }
        }
      ]
      topics = [
        {
          exact_list = {
            conflict = "warn"
            exact_list = [
              {
                backend = "...my_backend..."
              }
            ]
          }
        }
      ]
    }
    mode   = "hide_prefix"
    prefix = "...my_prefix..."
  }
}