resource "konnect_mesh_multi_zone_service" "my_meshmultizoneservice" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    ports = [
      {
        app_protocol = "...my_app_protocol..."
        name         = "...my_name..."
        port         = 5
      }
    ]
    selector = {
      mesh_service = {
        match_labels = {
          key = "value",
        }
      }
    }
  }
  type = "MeshMultiZoneService"
}