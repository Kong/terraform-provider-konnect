resource "konnect_gateway_control_plane" "plugin_ai_semantic_cache_cp" {
  name         = "Terraform Control Plane For AI Semantic Cache Plugin With Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_partial" "pgvector_partial" {
  vectordb = {
    name = "my_tf_pgvector_partial"
    config = {
      dimensions      = 1536
      distance_metric = "cosine"
      strategy        = "pgvector"
      pgvector = {
        host     = "pgvector.example.com"
        port     = 5432
        database = "kong-pgvector"
        user     = "postgres"
      }
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ai_semantic_cache_cp.id
}

resource "konnect_gateway_plugin_ai_semantic_cache" "my_ai_semantic_cache" {
  enabled = true

  partials = [{
    id   = konnect_gateway_partial.pgvector_partial.id
    path = "config.vectordb"
  }]

  config = {
    embeddings = {
      model = {
        name     = "text-embedding-3-small"
        provider = "openai"
      }
    }

    vectordb = null
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ai_semantic_cache_cp.id
}
