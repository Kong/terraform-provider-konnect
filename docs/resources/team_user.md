---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_team_user Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  TeamUser Resource
---

# konnect_team_user (Resource)

TeamUser Resource

## Example Usage

```terraform
resource "konnect_team_user" "my_teamuser" {
  team_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  user_id = "df120cb4-f60b-47bc-a2f8-6a28e6a3c63b"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `team_id` (String) ID of the team. Requires replacement if changed.
- `user_id` (String) The user ID for the user being added to a team. Requires replacement if changed.
