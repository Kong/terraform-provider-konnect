resource "konnect_gateway_plugin_openid_connect" "my_gatewaypluginopenidconnect" {
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
      "session"
    ]
    authenticated_groups_claim = [
      "..."
    ]
    authorization_cookie_domain    = "...my_authorization_cookie_domain..."
    authorization_cookie_http_only = false
    authorization_cookie_name      = "...my_authorization_cookie_name..."
    authorization_cookie_path      = "...my_authorization_cookie_path..."
    authorization_cookie_same_site = "Default"
    authorization_cookie_secure    = false
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
    authorization_rolling_timeout = 1.26
    bearer_token_cookie_name      = "...my_bearer_token_cookie_name..."
    bearer_token_param_type = [
      "header"
    ]
    by_username_ignore_case = false
    cache_introspection     = true
    cache_token_exchange    = false
    cache_tokens            = false
    cache_tokens_salt       = "...my_cache_tokens_salt..."
    cache_ttl               = 4.51
    cache_ttl_max           = 8.18
    cache_ttl_min           = 0.48
    cache_ttl_neg           = 5.85
    cache_ttl_resurrect     = 0.5
    cache_user_info         = false
    claims_forbidden = [
      "..."
    ]
    client_alg = [
      "RS512"
    ]
    client_arg = "...my_client_arg..."
    client_auth = [
      "client_secret_jwt"
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
      cluster_max_redirections = 5
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 39126
        }
      ]
      connect_timeout       = 1007376275
      connection_is_proxied = false
      database              = 6
      host                  = "...my_host..."
      keepalive_backlog     = 513691764
      keepalive_pool_size   = 742855137
      password              = "...my_password..."
      port                  = 25288
      read_timeout          = 1652724306
      send_timeout          = 24704322
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 5690
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "master"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = true
      ssl_verify        = true
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
      "client_credentials"
    ]
    discovery_headers_names = [
      "..."
    ]
    discovery_headers_values = [
      "..."
    ]
    display_errors = false
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
    dpop_proof_lifetime                 = 9.34
    dpop_use_nonce                      = true
    enable_hs_signatures                = true
    end_session_endpoint                = "...my_end_session_endpoint..."
    expose_error_code                   = false
    extra_jwks_uris = [
      "..."
    ]
    forbidden_destroy_session = false
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
    hide_credentials          = true
    http_proxy                = "...my_http_proxy..."
    http_proxy_authorization  = "...my_http_proxy_authorization..."
    http_version              = 2.54
    https_proxy               = "...my_https_proxy..."
    https_proxy_authorization = "...my_https_proxy_authorization..."
    id_token_param_name       = "...my_id_token_param_name..."
    id_token_param_type = [
      "body"
    ]
    ignore_signature = [
      "session"
    ]
    introspect_jwt_tokens              = true
    introspection_accept               = "application/json"
    introspection_check_active         = false
    introspection_endpoint             = "...my_introspection_endpoint..."
    introspection_endpoint_auth_method = "client_secret_basic"
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
    leeway             = 4.43
    login_action       = "upstream"
    login_methods = [
      "password"
    ]
    login_redirect_mode = "fragment"
    login_redirect_uri = [
      "..."
    ]
    login_tokens = [
      "tokens"
    ]
    logout_methods = [
      "GET"
    ]
    logout_post_arg  = "...my_logout_post_arg..."
    logout_query_arg = "...my_logout_query_arg..."
    logout_redirect_uri = [
      "..."
    ]
    logout_revoke               = true
    logout_revoke_access_token  = false
    logout_revoke_refresh_token = false
    logout_uri_suffix           = "...my_logout_uri_suffix..."
    max_age                     = 0.81
    mtls_introspection_endpoint = "...my_mtls_introspection_endpoint..."
    mtls_revocation_endpoint    = "...my_mtls_revocation_endpoint..."
    mtls_token_endpoint         = "...my_mtls_token_endpoint..."
    no_proxy                    = "...my_no_proxy..."
    password_param_type = [
      "header"
    ]
    preserve_query_args                               = true
    proof_of_possession_auth_methods_validation       = true
    proof_of_possession_dpop                          = "strict"
    proof_of_possession_mtls                          = "off"
    pushed_authorization_request_endpoint             = "...my_pushed_authorization_request_endpoint..."
    pushed_authorization_request_endpoint_auth_method = "none"
    redirect_uri = [
      "..."
    ]
    redis = {
      cluster_max_redirections = 9
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 55819
        }
      ]
      connect_timeout       = 829309575
      connection_is_proxied = true
      database              = 2
      host                  = "...my_host..."
      keepalive_backlog     = 1420640006
      keepalive_pool_size   = 147781497
      password              = "...my_password..."
      port                  = 20220
      prefix                = "...my_prefix..."
      read_timeout          = 2120279470
      send_timeout          = 523577252
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 58352
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "any"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      socket            = "...my_socket..."
      ssl               = true
      ssl_verify        = true
      username          = "...my_username..."
    }
    rediscovery_lifetime     = 0.82
    refresh_token_param_name = "...my_refresh_token_param_name..."
    refresh_token_param_type = [
      "query"
    ]
    refresh_tokens                        = true
    require_proof_key_for_code_exchange   = true
    require_pushed_authorization_requests = true
    require_signed_request_object         = false
    resolve_distributed_claims            = true
    response_mode                         = "query.jwt"
    response_type = [
      "..."
    ]
    reverify                        = false
    revocation_endpoint             = "...my_revocation_endpoint..."
    revocation_endpoint_auth_method = "none"
    revocation_token_param_name     = "...my_revocation_token_param_name..."
    roles_claim = [
      "..."
    ]
    roles_required = [
      "..."
    ]
    run_on_preflight = true
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
    session_absolute_timeout          = 6.27
    session_audience                  = "...my_session_audience..."
    session_cookie_domain             = "...my_session_cookie_domain..."
    session_cookie_http_only          = false
    session_cookie_name               = "...my_session_cookie_name..."
    session_cookie_path               = "...my_session_cookie_path..."
    session_cookie_same_site          = "Strict"
    session_cookie_secure             = true
    session_enforce_same_subject      = false
    session_hash_storage_key          = false
    session_hash_subject              = false
    session_idling_timeout            = 9.33
    session_memcached_host            = "...my_session_memcached_host..."
    session_memcached_port            = 10230
    session_memcached_prefix          = "...my_session_memcached_prefix..."
    session_memcached_socket          = "...my_session_memcached_socket..."
    session_remember                  = false
    session_remember_absolute_timeout = 6.89
    session_remember_cookie_name      = "...my_session_remember_cookie_name..."
    session_remember_rolling_timeout  = 2.91
    session_request_headers = [
      "audience"
    ]
    session_response_headers = [
      "id"
    ]
    session_rolling_timeout       = 5.68
    session_secret                = "...my_session_secret..."
    session_storage               = "memcache"
    session_store_metadata        = true
    ssl_verify                    = true
    timeout                       = 0.75
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
      "client_credentials"
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
    userinfo_accept                   = "application/json"
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
    verify_claims       = true
    verify_nonce        = false
    verify_parameters   = true
    verify_signature    = false
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
    "tls_passthrough"
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