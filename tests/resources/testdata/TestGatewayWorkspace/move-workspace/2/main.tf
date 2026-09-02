resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Workspace Move"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "target" {
  name        = "tf-acceptance-workspace-move-target"
  description = "Destination workspace for the move test"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Step 2: identical except for `workspace`. The attribute carries
# RequiresReplaceIfConfigured, so moving an entity between workspaces must
# destroy and re-create it rather than silently updating in place - an entity
# cannot be relocated across workspaces by the API.
resource "konnect_gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace        = konnect_gateway_workspace.target.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
