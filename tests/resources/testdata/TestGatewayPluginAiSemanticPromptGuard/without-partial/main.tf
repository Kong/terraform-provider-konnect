resource "konnect_gateway_control_plane" "prompt_guard_cp" {
  name         = "Terraform Control Plane For AI Semantic Prompt Guard Plugin No Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_semantic_prompt_guard" "my_plugin" {
  enabled = true

  config = {
    embeddings = {
      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.openai_api_key"
      }
      model = {
        name     = "text-embedding-3-small"
        provider = "openai"
      }
    }

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
      redis = {}
    }

    rules = {
      allow_prompts = ["How do I reset my password?"]
      deny_prompts  = ["How do I hack a server?"]
    }

    search = {
      threshold = 0.7
    }
  }

  control_plane_id = konnect_gateway_control_plane.prompt_guard_cp.id
}
