resource "konnect_metering_meter" "my_meteringmeter" {
  aggregation = "count"
  event_type  = "prompt"
  key         = "tftestmeterkey"
  name        = "tftestmetername"
  description = "bbbb"
}
