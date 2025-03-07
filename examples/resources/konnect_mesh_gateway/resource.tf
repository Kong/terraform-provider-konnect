resource "konnect_mesh_gateway" "my_meshgateway" {
  conf = {
    listeners = [
      {
        cross_mesh = true
        hostname   = "...my_hostname..."
        port       = 5
        protocol = {
          integer = 4
          str     = "...my_str..."
        }
        resources = {
          connection_limit = 5
        }
        tags = {
          key = "value"
        }
        tls = {
          certificates = [
            {
              type = "{ \"see\": \"documentation\" }"
            }
          ]
          mode = {
            integer = 4
            str     = "...my_str..."
          }
          options = {
            # ...
          }
        }
      }
    ]
  }
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  selectors = [
    {
      match = {
        key = "value"
      }
    }
  ]
  tags = {
    key = "value"
  }
  type = "...my_type..."
}