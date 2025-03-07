// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshOPAItem struct {
	CreationTime     types.String            `tfsdk:"creation_time"`
	Labels           map[string]types.String `tfsdk:"labels"`
	Mesh             types.String            `tfsdk:"mesh"`
	ModificationTime types.String            `tfsdk:"modification_time"`
	Name             types.String            `tfsdk:"name"`
	Spec             MeshOPAItemSpec         `tfsdk:"spec"`
	Type             types.String            `tfsdk:"type"`
}
