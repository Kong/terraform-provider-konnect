resource "konnect_team" "my_team" {
  name        = "My Terraform Team"
  description = "This is a team that is managed by Terraform"

  labels = {
    example = "here"
  }
}