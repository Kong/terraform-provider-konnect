# A principal is the identity a client authenticates as. type = "agent" marks a
# deployed AI agent, type = "user" a person in the organization.
resource "konnect_identity_principal" "my_principal" {
  directory_id = konnect_identity_directory.my_directory.id
  display_name = "Payments Service"
  description  = "Principal used by the payments service"
  type         = "agent"

  metadata = jsonencode({
    team = "payments"
  })
}
