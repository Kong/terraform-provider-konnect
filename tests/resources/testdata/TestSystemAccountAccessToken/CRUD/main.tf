resource "konnect_system_account" "konnect_sys_account" {
    name        = "provider-token-test-system-account"
    description = "description"
    konnect_managed = false
}

# These timestamps need to be in the future, so need to be bumped every year.
# Dynamic functions like https://developer.hashicorp.com/terraform/language/functions/timestamp are evaluated at apply time,
# and so cannot be used here - otherwise lead to false diffs.
resource "konnect_system_account_access_token" "my_sys_token_1" {
    name       = "first-token-local-offset"
    account_id = konnect_system_account.konnect_sys_account.id
    expires_at = "2026-09-22T07:00:00-05:30"
}

resource "konnect_system_account_access_token" "my_sys_token_2" {
    name       = "second-token-with-ms"
    account_id = konnect_system_account.konnect_sys_account.id
    expires_at = "2026-09-22T07:00:00.112Z"
}

resource "konnect_system_account_access_token" "my_sys_token_3" {
    name       = "third-token"
    account_id = konnect_system_account.konnect_sys_account.id
    expires_at = "2026-09-22T07:00:00Z"
}
