resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "My AI Gateway"
  name         = "my-ai-gateway"
}

# pgvector datastore, used by semantic cache and RAG policies.
resource "konnect_ai_gateway_datastore" "my_vectordb" {
  gateway_id = konnect_ai_gateway.my_aigateway.id

  vectordb = {
    name         = "my-vectordb"
    display_name = "Vector DB"
    description  = "pgvector store for semantic cache and RAG"

    config = {
      host     = "pgvector.example.com"
      database = "kong-pgvector"
      user     = "kong"
      password = "demo-pgvector-password"
    }
  }
}

# Redis datastore, shared by the rate limiting and semantic cache policies.
resource "konnect_ai_gateway_datastore" "my_redis" {
  gateway_id = konnect_ai_gateway.my_aigateway.id

  redis_ce = {
    name         = "my-redis"
    display_name = "Shared Redis"
    description  = "Redis shared by rate limiting and semantic cache"

    config = {
      host     = "redis.example.com"
      port     = 6379
      username = "default"
      password = "demo-redis-password"
      ssl      = true
    }
  }
}

# Token vault datastore, which stores the upstream OAuth tokens agents receive.
# Keys are vault references, never literal key material. Keep the previous
# versions in the list so existing tokens can still be decrypted.
resource "konnect_ai_gateway_datastore" "my_token_vault" {
  gateway_id = konnect_ai_gateway.my_aigateway.id

  token_vault = {
    name         = "my-token-vault"
    display_name = "Token Vault"
    description  = "Stores upstream OAuth tokens issued to agents"

    config = {
      directory       = konnect_identity_directory.my_directory.name
      current_version = "v2"

      keys = [
        {
          key     = "{vault://hcv/token-vault-key-v1}"
          version = "v1"
        },
        {
          key     = "{vault://hcv/token-vault-key-v2}"
          version = "v2"
        },
      ]
    }
  }
}
