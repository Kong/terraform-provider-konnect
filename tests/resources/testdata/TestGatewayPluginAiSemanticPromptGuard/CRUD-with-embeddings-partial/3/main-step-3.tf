resource "konnect_gateway_control_plane" "prompt_guard_cp" {
  name         = "Terraform Control Plane For AI Semantic Prompt Guard With Embeddings Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_partial" "embeddings_partial" {
  embeddings = {
    name = "my_tf_embeddings_partial"
    config = {
      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.openai_api_key"
      }
      model = {
        provider = "openai"
        name     = "text-embedding-3-small"
      }
    }
  }

  control_plane_id = konnect_gateway_control_plane.prompt_guard_cp.id
}

resource "konnect_gateway_plugin_ai_semantic_prompt_guard" "my_plugin" {
  enabled = true

  partials = [{
    id   = konnect_gateway_partial.embeddings_partial.id
    path = "config.embeddings"
  }]

  config = {
    embeddings = null

    vectordb = {
      strategy        = "pgvector"
      distance_metric = "cosine"
      dimensions      = 1536
      pgvector = {
        host     = "pgvector.example.com"
        port     = 5432
        database = "kong-pgvector"
        user     = "postgres"
      }
    }

    rules = {
      allow_prompts = ["How do I reset my password?"]
      deny_prompts  = ["How do I hack a server?"]
    }

    search = {
      threshold = 0.7
    }
  }

  tags = ["updated"]

  control_plane_id = konnect_gateway_control_plane.prompt_guard_cp.id
}
