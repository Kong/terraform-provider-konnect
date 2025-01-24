resource "konnect_gateway_plugin_acme" "my_gatewaypluginacme" {
  config = {
    account_email = "...my_account_email..."
    account_key = {
      key_id  = "...my_key_id..."
      key_set = "...my_key_set..."
    }
    allow_any_domain = false
    api_uri          = "...my_api_uri..."
    cert_type        = "rsa"
    domains = [
      "..."
    ]
    eab_hmac_key            = "...my_eab_hmac_key..."
    eab_kid                 = "...my_eab_kid..."
    enable_ipv4_common_name = false
    fail_backoff_minutes    = 7.02
    preferred_chain         = "...my_preferred_chain..."
    renew_threshold_days    = 1.5
    rsa_key_size            = 3072
    storage                 = "shm"
    storage_config = {
      consul = {
        host    = "...my_host..."
        https   = false
        kv_path = "...my_kv_path..."
        port    = 27961
        timeout = 6.85
        token   = "...my_token..."
      }
      kong = {
        key = jsonencode("value"),
      }
      redis = {
        database = 6
        extra_options = {
          namespace  = "...my_namespace..."
          scan_count = 4.65
        }
        host        = "...my_host..."
        password    = "...my_password..."
        port        = 60103
        server_name = "...my_server_name..."
        ssl         = true
        ssl_verify  = false
        timeout     = 1734626631
        username    = "...my_username..."
      }
      shm = {
        shm_name = "...my_shm_name..."
      }
      vault = {
        auth_method     = "kubernetes"
        auth_path       = "...my_auth_path..."
        auth_role       = "...my_auth_role..."
        host            = "...my_host..."
        https           = false
        jwt_path        = "...my_jwt_path..."
        kv_path         = "...my_kv_path..."
        port            = 58658
        timeout         = 3.3
        tls_server_name = "...my_tls_server_name..."
        tls_verify      = true
        token           = "...my_token..."
      }
    }
    tos_accepted = true
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
    "grpcs"
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