resource "konnect_metering_customer" "my_meteringcustomer" {
  key  = "tftestsubcustomerkey"
  name = "tftestsubcustomername"
}

resource "konnect_metering_plan" "my_meteringplan" {
  billing_cadence = "P1M"
  currency        = "USD"
  key             = "tftestsubplankey"
  name            = "tftestsubplanname"
  phases = [
    {
      key  = "tftestphasekey"
      name = "tftestphasename"
      rate_cards = [
        {
          key  = "tftestratecardkey"
          name = "tftestratecardname"
          price = {
            free = {
              type = "free"
            }
          }
        }
      ]
    }
  ]
}

resource "konnect_metering_subscription" "my_meteringsubscription" {
  customer = {
    id = konnect_metering_customer.my_meteringcustomer.id
  }
  plan = {
    id = konnect_metering_plan.my_meteringplan.id
  }
}
