data "konnect_portal_list" "my_portallist" {
  page_number = 1
  page_size   = 10
  sort        = "...my_sort..."
}