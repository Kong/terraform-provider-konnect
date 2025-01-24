resource "konnect_mesh" "my_mesh" {
  constraints = {
    dataplane_proxy = {
      requirements = [
        {
          tags = {
            key = "value",
          }
        }
      ]
      restrictions = [
        {
          tags = {
            key = "value",
          }
        }
      ]
    }
  }
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  logging = {
    backends = [
      {
        conf   = "{ \"see\": \"documentation\" }"
        format = "...my_format..."
        name   = "...my_name..."
        type   = "...my_type..."
      }
    ]
    default_backend = "...my_default_backend..."
  }
  mesh_services = {
    mode = {
      integer = 5
      str     = "...my_str..."
    }
  }
  metrics = {
    backends = [
      {
        conf = "{ \"see\": \"documentation\" }"
        name = "...my_name..."
        type = "...my_type..."
      }
    ]
    enabled_backend = "...my_enabled_backend..."
  }
  mtls = {
    backends = [
      {
        conf = "{ \"see\": \"documentation\" }"
        dp_cert = {
          request_timeout = {
            nanos   = 5
            seconds = 5
          }
          rotation = {
            expiration = "...my_expiration..."
          }
        }
        mode = {
          integer = 4
          str     = "...my_str..."
        }
        name = "...my_name..."
        root_chain = {
          request_timeout = {
            nanos   = 9
            seconds = 4
          }
        }
        type = "...my_type..."
      }
    ]
    enabled_backend = "...my_enabled_backend..."
    skip_validation = true
  }
  name = "...my_name..."
  networking = {
    outbound = {
      passthrough = {
        value = true
      }
    }
  }
  routing = {
    default_forbid_mesh_external_service_access = false
    locality_aware_load_balancing               = false
    zone_egress                                 = false
  }
  skip_creating_initial_policies = [
    "..."
  ]
  tracing = {
    backends = [
      {
        conf = "{ \"see\": \"documentation\" }"
        name = "...my_name..."
        sampling = {
          value = 9.31
        }
        type = "...my_type..."
      }
    ]
    default_backend = "...my_default_backend..."
  }
  type = "...my_type..."
}