resource "konnect_event_gateway" "test_egw" {
  name = "test_policy_chain_egw"
}

resource "konnect_event_gateway_backend_cluster" "test_backend_cluster" {
    name = "test_policy_chain_backend_cluster"
    gateway_id = konnect_event_gateway.test_egw.id
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
}

resource "konnect_event_gateway_virtual_cluster" "test_virtual_cluster" {
    name      = "test_policy_chain_virtual_cluster"
    dns_label = "vcluster-policy-chain"
    acl_mode  = "enforce_on_gateway"
    gateway_id = konnect_event_gateway.test_egw.id
    authentication = [
        {
            anonymous = {}
        }
    ]
    destination = {
        id = konnect_event_gateway_backend_cluster.test_backend_cluster.id
    }
}

resource "konnect_event_gateway_listener" "test_listener" {
    name        = "test-listener-policy-chain"
    description = "Listener for policy chain test"
    addresses   = ["0.0.0.0"]
    ports       = ["9092", "9093", "9094", "9095"]
    gateway_id = konnect_event_gateway.test_egw.id
}

resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "test_listener_policy" {
    name        = "test-listener-forward-policy"
    description = "Test forward policy for chain"
    enabled     = true

    config = {
        port_mapping = {
            advertised_host = "broker.example.com"
            bootstrap_port  = "at_start"
            min_broker_id   = 0
            destination = {
                id = konnect_event_gateway_virtual_cluster.test_virtual_cluster.id
            }
        }
    }
    listener_id = konnect_event_gateway_listener.test_listener.id
    gateway_id = konnect_event_gateway.test_egw.id
}

resource "konnect_event_gateway_listener_policy_tls_server" "test_listener_policy_tls" {
    name        = "test-listener-tls-policy"
    description = "Test TLS policy for chain"
    enabled     = true

    config = {
        allow_plaintext = false
        certificates = [
            {
                certificate = "-----BEGIN CERTIFICATE-----\nMIIBxjCCAUygAwIBAgIUX9TaLbWF76yQc8IGR+YRbeiDlHkwCgYIKoZIzj0EAwIwGjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMB4XDTI0MDMwMTE0MzkxNloXDTI3MDMwMTE0MzkxNlowGjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEcMndCotXzeZ9vGAMfDfZ7UxUuP5bcIrwwUOI8YlpMdvB12HvjtS7O0/ONr3fBeCWagRuitPEqd4b3EJuD8kuFUMt+2A09N6KY1YDJWgKHei7rzKgrefzVt11XgBiDsUBo1MwUTAdBgNVHQ4EFgQUIrdAC8p02h60GZW0Jlh2Vcg/WeMwHwYDVR0jBBgwFoAUIrdAC8p02h60GZW0Jlh2Vcg/WeMwDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNoADBlAjBYb+yQf33sItlmsONLc41Agtx73FMEN7LfWA85OtlkMie1N1x0mj08pzS/Xc1VONwCMQDN9sBn3Kody0gse+EXYSuPPj1oo9jmFB9/xrpz35YpDATvuyhH8xwSJ4xMuxQiduc=\n-----END CERTIFICATE-----"
                key = "-----BEGIN PRIVATE KEY-----\nMIG2AgEAMBAGByqGSM49AgEGBSuBBAAiBIGeMIGbAgEBBDDLuRX+uzSbstvLWsQrWwuGK4AdjLU/tN9A/fn03gxNvppKw++SBtnLyB+9YZ29YA+hZANiAARwyd0Ki1fN5n28YAx8N9ntTFS4/ltwivDBQ4jxiWkx28HXYe+O1Ls7T842vd8F4JZqBG6K08Sp3hvcQm4PyS4VQy37YDT03opjVgMlaAod6LuvMqCt5/NW3XVeAGIOxQE=\n-----END PRIVATE KEY-----"
            }
        ]
        versions = {
            min = "TLSv1.2"
            max = "TLSv1.3"
        }
    }

    listener_id = konnect_event_gateway_listener.test_listener.id
    gateway_id = konnect_event_gateway.test_egw.id
}

resource "konnect_event_gateway_listener_policy_chain" "test_listener_policy_chain" {
    policies = [
        {
            id = konnect_event_gateway_listener_policy_forward_to_virtual_cluster.test_listener_policy.id
        },
        {
            id = konnect_event_gateway_listener_policy_tls_server.test_listener_policy_tls.id
        }
    ]

    listener_id = konnect_event_gateway_listener.test_listener.id
    gateway_id = konnect_event_gateway.test_egw.id
}
