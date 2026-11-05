package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/kong/terraform-provider-konnect/v3/internal/stateupgraders"
)

func init() {
	ctx := context.Background()
	p := &KonnectProvider{}

	for _, newResource := range p.Resources(ctx) {
		r := newResource()
		if _, ok := r.(resource.ResourceWithUpgradeState); !ok {
			continue
		}

		var metaResp resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "konnect"}, &metaResp)

		var schemaResp resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

		stateupgraders.RegisterSchemaType(metaResp.TypeName, schemaResp.Schema.Type().TerraformType(ctx))
	}
}
