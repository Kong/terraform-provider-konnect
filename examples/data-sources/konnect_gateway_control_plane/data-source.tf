data "konnect_gateway_control_plane" "my_gatewaycontrolplane" {
  filter = {
    cloud_gateway = true
    cluster_type = {
      eq  = "...my_eq..."
      neq = "...my_neq..."
    }
    id = {
      eq  = "...my_eq..."
      oeq = "...my_oeq..."
    }
    name = {
      contains = "...my_contains..."
      eq       = "...my_eq..."
      neq      = "...my_neq..."
    }
  }
  filter_labels = "key:value,existCheck"
  sort          = "created_at desc"
}