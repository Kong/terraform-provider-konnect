---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_system_account_team Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  SystemAccountTeam Resource
---

# konnect_system_account_team (Resource)

SystemAccountTeam Resource

## Example Usage

```terraform
resource "konnect_system_account_team" "my_systemaccountteam" {
  account_id = "501e14fa-d023-4728-b2a2-a14691ae5c05"
  team_id    = "...my_team_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `team_id` (String) ID of the team. Requires replacement if changed.

### Optional

- `account_id` (String) ID of the system account. Requires replacement if changed.