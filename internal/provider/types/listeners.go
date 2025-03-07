// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Listeners struct {
	CrossMesh types.Bool              `tfsdk:"cross_mesh"`
	Hostname  types.String            `tfsdk:"hostname"`
	Port      types.Int64             `tfsdk:"port"`
	Protocol  *Mode                   `tfsdk:"protocol"`
	Resources *Resources              `tfsdk:"resources"`
	Tags      map[string]types.String `tfsdk:"tags"`
	TLS       *MeshGatewayItemTLS     `tfsdk:"tls"`
}
