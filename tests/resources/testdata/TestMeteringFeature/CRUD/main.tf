resource "konnect_metering_feature" "my_meteringfeature" {
  key  = "tftestfeaturekey"
  name = "tftestfeaturename"
  unit_cost = {
    manual = {
      amount = "1.00"
      type   = "manual"
    }
  }
}
