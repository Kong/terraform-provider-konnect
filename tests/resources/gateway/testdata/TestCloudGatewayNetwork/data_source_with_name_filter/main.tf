# Data source test with name filter

data "konnect_cloud_gateway_network" "test_network_by_name" {
  filter = {
    name = {
      contains = "ci-test"
    }
  }
}

output "network_name" {
  value = data.konnect_cloud_gateway_network.test_network_by_name
}

output "network_id" {
  value = data.konnect_cloud_gateway_network.test_network_by_name.id
}
