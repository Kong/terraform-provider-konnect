resource "konnect_ai_gateway_cost_management_tags" "my_aigatewaycostmanagementtags" {
  headers = [
    {
      name    = "tftest-tenant-id"
      tag_key = "tftest-tenant-updated"
    }
  ]
  prefixes = [
    {
      name = "tftest-co"
    }
  ]
}
