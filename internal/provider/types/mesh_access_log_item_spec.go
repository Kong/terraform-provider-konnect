// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type MeshAccessLogItemSpec struct {
	From      []From                      `tfsdk:"from"`
	Rules     []Rules                     `tfsdk:"rules"`
	TargetRef *MeshAccessLogItemTargetRef `tfsdk:"target_ref"`
	To        []From                      `tfsdk:"to"`
}
