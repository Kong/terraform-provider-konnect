# A directory is the boundary that groups principals, their groups and their
# credentials. Set vault_enabled to use the token vault with this directory.
resource "konnect_identity_directory" "my_directory" {
  name          = "my-directory"
  description   = "Directory for the payments platform"
  vault_enabled = true

  allow_all_control_planes = true

  # How long a data plane caches a principal it authenticated from this
  # directory, and a principal it failed to authenticate.
  ttl_secs          = 600
  negative_ttl_secs = 600
}
