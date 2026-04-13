data "konnect_cloud_gateway_network" "test_network_by_name" {
  filter = {
    name = {
      contains = "my-network"
    }
  }
}