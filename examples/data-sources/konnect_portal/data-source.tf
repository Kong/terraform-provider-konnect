data "konnect_portal" "my_portal" {
  filter = {
    authentication_enabled    = false
    auto_approve_applications = true
    auto_approve_developers   = true
    canonical_domain = {
      contains  = "...my_contains..."
      eq        = "...my_eq..."
      neq       = "...my_neq..."
      ocontains = "...my_ocontains..."
      oeq       = "...my_oeq..."
    }
    default_api_visibility = {
      contains  = "...my_contains..."
      eq        = "...my_eq..."
      neq       = "...my_neq..."
      ocontains = "...my_ocontains..."
      oeq       = "...my_oeq..."
    }
    default_application_auth_strategy_id = {
      eq  = "...my_eq..."
      neq = "...my_neq..."
      oeq = "...my_oeq..."
    }
    default_domain = {
      contains  = "...my_contains..."
      eq        = "...my_eq..."
      neq       = "...my_neq..."
      ocontains = "...my_ocontains..."
      oeq       = "...my_oeq..."
    }
    default_page_visibility = {
      contains  = "...my_contains..."
      eq        = "...my_eq..."
      neq       = "...my_neq..."
      ocontains = "...my_ocontains..."
      oeq       = "...my_oeq..."
    }
    description = {
      contains  = "...my_contains..."
      eq        = "...my_eq..."
      neq       = "...my_neq..."
      ocontains = "...my_ocontains..."
      oeq       = "...my_oeq..."
    }
    id = {
      eq  = "...my_eq..."
      neq = "...my_neq..."
      oeq = "...my_oeq..."
    }
    name = {
      contains  = "...my_contains..."
      eq        = "...my_eq..."
      neq       = "...my_neq..."
      ocontains = "...my_ocontains..."
      oeq       = "...my_oeq..."
    }
    rbac_enabled = true
  }
  sort = "...my_sort..."
}