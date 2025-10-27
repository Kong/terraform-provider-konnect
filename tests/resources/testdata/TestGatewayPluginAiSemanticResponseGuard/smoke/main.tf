resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI Semantic Response Guard Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_semantic_response_guard" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

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

    search = {
      threshold = 0.7
    }

    vectordb = {
      strategy        = "pgvector"
      distance_metric = "cosine"
      threshold       = 0.7
      dimensions      = 1024

      pgvector = {
        host         = "pgvector_host"
        port         = 5432
        database     = "kong-pgvector"
        user         = "pgvector_user"
        password     = "pgvector_password"
        ssl          = false
        ssl_required = false
        ssl_verify   = false
        ssl_version  = "tlsv1_2"
        timeout      = 5000
      }
    }

    rules = {
      allow_responses = ["Troubleshooting networks and connectivity issues", "Managing cloud platforms (AWS, Azure, GCP)", "Security hardening and incident response strategies", "DevOps pipelines, automation, and observability", "Software engineering concepts and language syntax", "IT governance, compliance, and regulatory guidance", "Continuous integration and deployment practices", "Writing documentation and explaining technical concepts", "Operating system administration and configuration", "Best practices for collaboration and productivity tools"]
      deny_responses  = ["Unauthorized penetration testing or hacking tutorials", "Methods for bypassing software licensing or DRM", "Step-by-step instructions for exploiting vulnerabilities", "Techniques to evade or disable security controls", "Collecting or exposing personal or employee data", "Using AI for impersonation, phishing, or fraud", "Manipulative social engineering techniques", "Advice on breaking internal IT or security policies", "Entertainment, dating, or other non-work topics", "Political, religious, or otherwise sensitive discussions unrelated to work"]
    }
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
