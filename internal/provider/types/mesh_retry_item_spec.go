// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type MeshRetryItemSpec struct {
	TargetRef *MeshAccessLogItemTargetRef `tfsdk:"target_ref"`
	To        []MeshRetryItemTo           `tfsdk:"to"`
}
