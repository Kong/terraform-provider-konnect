resource "konnect_application_auth_strategy" "my_applicationauthstrategy" {
  key_auth = {
    configs = {
      key_auth = {
        key_names = [
          "..."
        ]
        ttl = {
          unit  = "years"
          value = 11
        }
      }
    }
    display_name = "...my_display_name..."
    labels = {
      key = "value"
    }
    name          = "...my_name..."
    strategy_type = "key_auth"
  }
}