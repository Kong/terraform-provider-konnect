resource "konnect_team" "my_team" {
  description = "The Identity Management (IDM) team."
  labels = {
    key = "value"
  }
  name = "IDM - Developers"
}