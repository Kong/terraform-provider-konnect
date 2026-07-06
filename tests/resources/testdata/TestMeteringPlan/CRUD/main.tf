resource "konnect_metering_plan" "my_meteringplan" {
  billing_cadence = "P1Y"
  currency        = "USD"
  description     = "aaaa"
  key             = "resource_keyy"
  name            = "Myfirstplan"
  phases = [{
    name = "my phase name",
    key  = "myphasekeyy",
    rate_cards = [
      {
        billing_cadence = "P1Y",
        name            = "my rate card name",
        key             = "myratecardkeyy",
        price = {
          free = {}
        }
      }
    ]
  }]
}
