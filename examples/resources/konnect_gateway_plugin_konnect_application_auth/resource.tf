resource "konnect_gateway_plugin_konnect_application_auth" "my_gatewaypluginkonnectapplicationauth" {
  config = {
    auth_type = "v2-strategies"
    key_names = [
      "..."
    ]
    scope = "...my_scope..."
    v2_strategies = {
      key_auth = [
        {
          config = {
            key_names = [
              "..."
            ]
          }
          strategy_id = "...my_strategy_id..."
        }
      ]
      openid_connect = [
        {
          config = {
            anonymous = "...my_anonymous..."
            audience = [
              "..."
            ]
            audience_claim = [
              "..."
            ]
            audience_required = [
              "..."
            ]
            auth_methods = [
              "client_credentials"
            ]
            authenticated_groups_claim = [
              "..."
            ]
            authorization_cookie_domain    = "...my_authorization_cookie_domain..."
            authorization_cookie_http_only = false
            authorization_cookie_name      = "...my_authorization_cookie_name..."
            authorization_cookie_path      = "...my_authorization_cookie_path..."
            authorization_cookie_same_site = "Default"
            authorization_cookie_secure    = true
            authorization_endpoint         = "...my_authorization_endpoint..."
            authorization_query_args_client = [
              "..."
            ]
            authorization_query_args_names = [
              "..."
            ]
            authorization_query_args_values = [
              "..."
            ]
            authorization_rolling_timeout = 8.13
            bearer_token_cookie_name      = "...my_bearer_token_cookie_name..."
            bearer_token_param_type = [
              "cookie"
            ]
            by_username_ignore_case = true
            cache_introspection     = false
            cache_token_exchange    = false
            cache_tokens            = true
            cache_tokens_salt       = "...my_cache_tokens_salt..."
            cache_ttl               = 5.66
            cache_ttl_max           = 3.33
            cache_ttl_min           = 1.4
            cache_ttl_neg           = 4.76
            cache_ttl_resurrect     = 8.57
            cache_user_info         = false
            claims_forbidden = [
              "..."
            ]
            client_alg = [
              "RS256"
            ]
            client_arg = "...my_client_arg..."
            client_auth = [
              "self_signed_tls_client_auth"
            ]
            client_credentials_param_type = [
              "body"
            ]
            client_id = [
              "..."
            ]
            client_jwk = [
              {
                alg    = "...my_alg..."
                crv    = "...my_crv..."
                d      = "...my_d..."
                dp     = "...my_dp..."
                dq     = "...my_dq..."
                e      = "...my_e..."
                issuer = "...my_issuer..."
                k      = "...my_k..."
                key_ops = [
                  "..."
                ]
                kid = "...my_kid..."
                kty = "...my_kty..."
                n   = "...my_n..."
                oth = "...my_oth..."
                p   = "...my_p..."
                q   = "...my_q..."
                qi  = "...my_qi..."
                r   = "...my_r..."
                t   = "...my_t..."
                use = "...my_use..."
                x   = "...my_x..."
                x5c = [
                  "..."
                ]
                x5t             = "...my_x5t..."
                x5t_number_s256 = "...my_x5t_number_s256..."
                x5u             = "...my_x5u..."
                y               = "...my_y..."
              }
            ]
            client_secret = [
              "..."
            ]
            cluster_cache_redis = {
              cluster_max_redirections = 0
              cluster_nodes = [
                {
                  ip   = "...my_ip..."
                  port = 20075
                }
              ]
              connect_timeout       = 1087513919
              connection_is_proxied = true
              database              = 5
              host                  = "...my_host..."
              keepalive_backlog     = 21059882
              keepalive_pool_size   = 1814531755
              password              = "...my_password..."
              port                  = 61462
              read_timeout          = 764507817
              send_timeout          = 1625033015
              sentinel_master       = "...my_sentinel_master..."
              sentinel_nodes = [
                {
                  host = "...my_host..."
                  port = 27278
                }
              ]
              sentinel_password = "...my_sentinel_password..."
              sentinel_role     = "any"
              sentinel_username = "...my_sentinel_username..."
              server_name       = "...my_server_name..."
              ssl               = true
              ssl_verify        = false
              username          = "...my_username..."
            }
            cluster_cache_strategy = "off"
            consumer_by = [
              "username"
            ]
            consumer_claim = [
              "..."
            ]
            consumer_optional = true
            credential_claim = [
              "..."
            ]
            disable_session = [
              "bearer"
            ]
            discovery_headers_names = [
              "..."
            ]
            discovery_headers_values = [
              "..."
            ]
            display_errors = true
            domains = [
              "..."
            ]
            downstream_access_token_header     = "...my_downstream_access_token_header..."
            downstream_access_token_jwk_header = "...my_downstream_access_token_jwk_header..."
            downstream_headers_claims = [
              "..."
            ]
            downstream_headers_names = [
              "..."
            ]
            downstream_id_token_header          = "...my_downstream_id_token_header..."
            downstream_id_token_jwk_header      = "...my_downstream_id_token_jwk_header..."
            downstream_introspection_header     = "...my_downstream_introspection_header..."
            downstream_introspection_jwt_header = "...my_downstream_introspection_jwt_header..."
            downstream_refresh_token_header     = "...my_downstream_refresh_token_header..."
            downstream_session_id_header        = "...my_downstream_session_id_header..."
            downstream_user_info_header         = "...my_downstream_user_info_header..."
            downstream_user_info_jwt_header     = "...my_downstream_user_info_jwt_header..."
            dpop_proof_lifetime                 = 4.13
            dpop_use_nonce                      = true
            enable_hs_signatures                = false
            end_session_endpoint                = "...my_end_session_endpoint..."
            expose_error_code                   = true
            extra_jwks_uris = [
              "..."
            ]
            forbidden_destroy_session = true
            forbidden_error_message   = "...my_forbidden_error_message..."
            forbidden_redirect_uri = [
              "..."
            ]
            groups_claim = [
              "..."
            ]
            groups_required = [
              "..."
            ]
            hide_credentials          = false
            http_proxy                = "...my_http_proxy..."
            http_proxy_authorization  = "...my_http_proxy_authorization..."
            http_version              = 5.05
            https_proxy               = "...my_https_proxy..."
            https_proxy_authorization = "...my_https_proxy_authorization..."
            id_token_param_name       = "...my_id_token_param_name..."
            id_token_param_type = [
              "query"
            ]
            ignore_signature = [
              "client_credentials"
            ]
            introspect_jwt_tokens              = true
            introspection_accept               = "application/json"
            introspection_check_active         = true
            introspection_endpoint             = "...my_introspection_endpoint..."
            introspection_endpoint_auth_method = "none"
            introspection_headers_client = [
              "..."
            ]
            introspection_headers_names = [
              "..."
            ]
            introspection_headers_values = [
              "..."
            ]
            introspection_hint = "...my_introspection_hint..."
            introspection_post_args_client = [
              "..."
            ]
            introspection_post_args_names = [
              "..."
            ]
            introspection_post_args_values = [
              "..."
            ]
            introspection_token_param_name = "...my_introspection_token_param_name..."
            issuer                         = "...my_issuer..."
            issuers_allowed = [
              "..."
            ]
            jwt_session_claim  = "...my_jwt_session_claim..."
            jwt_session_cookie = "...my_jwt_session_cookie..."
            keepalive          = true
            leeway             = 5
            login_action       = "response"
            login_methods = [
              "password"
            ]
            login_redirect_mode = "fragment"
            login_redirect_uri = [
              "..."
            ]
            login_tokens = [
              "id_token"
            ]
            logout_methods = [
              "GET"
            ]
            logout_post_arg  = "...my_logout_post_arg..."
            logout_query_arg = "...my_logout_query_arg..."
            logout_redirect_uri = [
              "..."
            ]
            logout_revoke               = false
            logout_revoke_access_token  = false
            logout_revoke_refresh_token = true
            logout_uri_suffix           = "...my_logout_uri_suffix..."
            max_age                     = 7.73
            mtls_introspection_endpoint = "...my_mtls_introspection_endpoint..."
            mtls_revocation_endpoint    = "...my_mtls_revocation_endpoint..."
            mtls_token_endpoint         = "...my_mtls_token_endpoint..."
            no_proxy                    = "...my_no_proxy..."
            password_param_type = [
              "header"
            ]
            preserve_query_args                               = false
            proof_of_possession_auth_methods_validation       = false
            proof_of_possession_dpop                          = "strict"
            proof_of_possession_mtls                          = "strict"
            pushed_authorization_request_endpoint             = "...my_pushed_authorization_request_endpoint..."
            pushed_authorization_request_endpoint_auth_method = "private_key_jwt"
            redirect_uri = [
              "..."
            ]
            redis = {
              cluster_max_redirections = 2
              cluster_nodes = [
                {
                  ip   = "...my_ip..."
                  port = 18265
                }
              ]
              connect_timeout       = 2071449819
              connection_is_proxied = false
              database              = 5
              host                  = "...my_host..."
              keepalive_backlog     = 1917297556
              keepalive_pool_size   = 749569645
              password              = "...my_password..."
              port                  = 48312
              prefix                = "...my_prefix..."
              read_timeout          = 982818545
              send_timeout          = 238944243
              sentinel_master       = "...my_sentinel_master..."
              sentinel_nodes = [
                {
                  host = "...my_host..."
                  port = 52207
                }
              ]
              sentinel_password = "...my_sentinel_password..."
              sentinel_role     = "master"
              sentinel_username = "...my_sentinel_username..."
              server_name       = "...my_server_name..."
              socket            = "...my_socket..."
              ssl               = true
              ssl_verify        = true
              username          = "...my_username..."
            }
            rediscovery_lifetime     = 3.59
            refresh_token_param_name = "...my_refresh_token_param_name..."
            refresh_token_param_type = [
              "body"
            ]
            refresh_tokens                        = false
            require_proof_key_for_code_exchange   = true
            require_pushed_authorization_requests = false
            require_signed_request_object         = true
            resolve_distributed_claims            = true
            response_mode                         = "query"
            response_type = [
              "..."
            ]
            reverify                        = true
            revocation_endpoint             = "...my_revocation_endpoint..."
            revocation_endpoint_auth_method = "tls_client_auth"
            revocation_token_param_name     = "...my_revocation_token_param_name..."
            roles_claim = [
              "..."
            ]
            roles_required = [
              "..."
            ]
            run_on_preflight = false
            scopes = [
              "..."
            ]
            scopes_claim = [
              "..."
            ]
            scopes_required = [
              "..."
            ]
            search_user_info                  = false
            session_absolute_timeout          = 7.82
            session_audience                  = "...my_session_audience..."
            session_cookie_domain             = "...my_session_cookie_domain..."
            session_cookie_http_only          = false
            session_cookie_name               = "...my_session_cookie_name..."
            session_cookie_path               = "...my_session_cookie_path..."
            session_cookie_same_site          = "Lax"
            session_cookie_secure             = true
            session_enforce_same_subject      = false
            session_hash_storage_key          = true
            session_hash_subject              = true
            session_idling_timeout            = 7.54
            session_memcached_host            = "...my_session_memcached_host..."
            session_memcached_port            = 30895
            session_memcached_prefix          = "...my_session_memcached_prefix..."
            session_memcached_socket          = "...my_session_memcached_socket..."
            session_remember                  = true
            session_remember_absolute_timeout = 6.68
            session_remember_cookie_name      = "...my_session_remember_cookie_name..."
            session_remember_rolling_timeout  = 7.84
            session_request_headers = [
              "timeout"
            ]
            session_response_headers = [
              "idling-timeout"
            ]
            session_rolling_timeout       = 2.38
            session_secret                = "...my_session_secret..."
            session_storage               = "cookie"
            session_store_metadata        = false
            ssl_verify                    = true
            timeout                       = 6.82
            tls_client_auth_cert_id       = "...my_tls_client_auth_cert_id..."
            tls_client_auth_ssl_verify    = false
            token_cache_key_include_scope = true
            token_endpoint                = "...my_token_endpoint..."
            token_endpoint_auth_method    = "client_secret_jwt"
            token_exchange_endpoint       = "...my_token_exchange_endpoint..."
            token_headers_client = [
              "..."
            ]
            token_headers_grants = [
              "password"
            ]
            token_headers_names = [
              "..."
            ]
            token_headers_prefix = "...my_token_headers_prefix..."
            token_headers_replay = [
              "..."
            ]
            token_headers_values = [
              "..."
            ]
            token_post_args_client = [
              "..."
            ]
            token_post_args_names = [
              "..."
            ]
            token_post_args_values = [
              "..."
            ]
            unauthorized_destroy_session = false
            unauthorized_error_message   = "...my_unauthorized_error_message..."
            unauthorized_redirect_uri = [
              "..."
            ]
            unexpected_redirect_uri = [
              "..."
            ]
            upstream_access_token_header     = "...my_upstream_access_token_header..."
            upstream_access_token_jwk_header = "...my_upstream_access_token_jwk_header..."
            upstream_headers_claims = [
              "..."
            ]
            upstream_headers_names = [
              "..."
            ]
            upstream_id_token_header          = "...my_upstream_id_token_header..."
            upstream_id_token_jwk_header      = "...my_upstream_id_token_jwk_header..."
            upstream_introspection_header     = "...my_upstream_introspection_header..."
            upstream_introspection_jwt_header = "...my_upstream_introspection_jwt_header..."
            upstream_refresh_token_header     = "...my_upstream_refresh_token_header..."
            upstream_session_id_header        = "...my_upstream_session_id_header..."
            upstream_user_info_header         = "...my_upstream_user_info_header..."
            upstream_user_info_jwt_header     = "...my_upstream_user_info_jwt_header..."
            userinfo_accept                   = "application/jwt"
            userinfo_endpoint                 = "...my_userinfo_endpoint..."
            userinfo_headers_client = [
              "..."
            ]
            userinfo_headers_names = [
              "..."
            ]
            userinfo_headers_values = [
              "..."
            ]
            userinfo_query_args_client = [
              "..."
            ]
            userinfo_query_args_names = [
              "..."
            ]
            userinfo_query_args_values = [
              "..."
            ]
            using_pseudo_issuer = true
            verify_claims       = false
            verify_nonce        = false
            verify_parameters   = true
            verify_signature    = true
          }
          strategy_id = "...my_strategy_id..."
        }
      ]
    }
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
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
    "tcp"
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