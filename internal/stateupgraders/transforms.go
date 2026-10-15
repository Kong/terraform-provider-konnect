package stateupgraders

// defaultWorkspaceIfMissing backfills `workspace` on state written before the
// attribute existed on this resource. `workspace` is a flat top-level
// attribute on every affected resource, so one transform covers all of them.
func defaultWorkspaceIfMissing(m map[string]any) {
	if v, ok := m["workspace"]; !ok || v == nil || v == "" {
		m["workspace"] = "default"
	}
}
