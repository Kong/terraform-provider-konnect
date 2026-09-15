data "konnect_platform_available_regions" "my_platformavailableregions" {

}


output "platform_available_regions" {
  value = data.konnect_platform_available_regions.my_platformavailableregions.regions
}