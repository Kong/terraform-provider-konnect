// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type MeshTimeoutItemSpec struct {
	From      []MeshTimeoutItemFrom       `tfsdk:"from"`
	Rules     []Rules                     `tfsdk:"rules"`
	TargetRef *MeshAccessLogItemTargetRef `tfsdk:"target_ref"`
	To        []MeshTimeoutItemFrom       `tfsdk:"to"`
}