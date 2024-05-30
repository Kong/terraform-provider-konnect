package custom

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

type CustomPluginResourceModel struct {
	ID             types.String  `tfsdk:"id"`
	Name           types.String  `tfsdk:"name"`
	Config         types.Dynamic `tfsdk:"config"`
	ControlPlaneID types.String  `tfsdk:"control_plane_id"`
}

func (r *CustomPluginResourceModel) ToSharedPluginInput() (shared.PluginInput, error) {
	config := r.Config.UnderlyingValue()

	// Use .String() to convert the dynamic value to a string,
	// then unmarshal the JSON string into a map[string]any
	var configJson map[string]any
	err := json.Unmarshal([]byte(config.String()), &configJson)
	if err != nil {
		return shared.PluginInput{}, err
	}

	return shared.PluginInput{
		Name:   sdk.String(r.Name.ValueString()),
		Config: configJson,
	}, nil
}
