resource "konnect_portal" "tfdemo" {
  name        = "Lookup Portal"
  description = "This is a sample description"
}

data "konnect_portal" "my_portal" {
  filter = {
    name = {
      eq = "Lookup Portal"
    }
  }
  depends_on = [konnect_portal.tfdemo]
}

output "portal" {
  value = data.konnect_portal.my_portal.name
}
