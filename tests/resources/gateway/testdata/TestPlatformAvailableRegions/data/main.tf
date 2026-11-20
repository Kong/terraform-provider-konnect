# 1. Declare the data source
data "konnect_platform_available_regions" "test" {}


output "platform_available_regions" {
  value = data.konnect_platform_available_regions.test.regions
}
