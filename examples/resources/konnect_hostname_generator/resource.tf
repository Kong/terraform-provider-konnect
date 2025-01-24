resource "konnect_hostname_generator" "my_hostnamegenerator" {
  cp_id         = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  creation_time = "0001-01-01T00:00:00Z"
  labels = {
    key = "value",
  }
  modification_time = "0001-01-01T00:00:00Z"
  name              = "...my_name..."
  spec = {
    selector = {
      mesh_external_service = {
        match_labels = {
          key = "value",
        }
      }
      mesh_multi_zone_service = {
        match_labels = {
          key = "value",
        }
      }
      mesh_service = {
        match_labels = {
          key = "value",
        }
      }
    }
    template = "...my_template..."
  }
  type = "HostnameGenerator"
}