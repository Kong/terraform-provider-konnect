package custom

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
)

type CustomPluginResourceModel struct {
	ID             types.String               `tfsdk:"id"`
	Name           types.String               `tfsdk:"name"`
	Config         types.Dynamic              `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer       `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer       `tfsdk:"consumer_group"`
	ControlPlaneID types.String               `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                `tfsdk:"created_at"`
	Enabled        types.Bool                 `tfsdk:"enabled"`
	InstanceName   types.String               `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering `tfsdk:"ordering"`
	Protocols      []types.String             `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer       `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer       `tfsdk:"service"`
	Tags           []types.String             `tfsdk:"tags"`
	UpdatedAt      types.Int64                `tfsdk:"updated_at"`
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

	pluginInput := shared.PluginInput{
		Name:   sdk.String(r.Name.ValueString()),
		Config: configJson,
	}

	if r.Enabled.ValueBoolPointer() != nil {
		pluginInput.Enabled = sdk.Bool(r.Enabled.ValueBool())
	}

	if r.InstanceName.ValueStringPointer() != nil {
		pluginInput.InstanceName = sdk.String(r.InstanceName.ValueString())
	}

	if r.Protocols != nil {
		protocols := []shared.Protocols{}
		for _, protocol := range r.Protocols {
			protocols = append(protocols, shared.Protocols(protocol.ValueString()))
		}
		pluginInput.Protocols = protocols
	}

	if r.Consumer != nil {
		pluginInput.Consumer = &shared.PluginConsumer{
			ID: sdk.String(r.Consumer.ID.ValueString()),
		}
	}

	if r.ConsumerGroup != nil {
		pluginInput.ConsumerGroup = &shared.PluginConsumerGroup{
			ID: sdk.String(r.ConsumerGroup.ID.ValueString()),
		}
	}

	if r.Route != nil {
		pluginInput.Route = &shared.PluginRoute{
			ID: sdk.String(r.Route.ID.ValueString()),
		}
	}

	if r.Service != nil {
		pluginInput.Service = &shared.PluginService{
			ID: sdk.String(r.Service.ID.ValueString()),
		}
	}

	if r.Tags != nil {
		tags := []string{}
		for _, tag := range r.Tags {
			tags = append(tags, tag.ValueString())
		}
		pluginInput.Tags = tags
	}

	if r.Ordering != nil {
		orderingBefore := []string{}
		orderingAfter := []string{}
		if r.Ordering.Before != nil && r.Ordering.Before.Access != nil {
			for _, before := range r.Ordering.Before.Access {
				orderingBefore = append(orderingBefore, before.ValueString())
			}
		}

		if r.Ordering.After != nil && r.Ordering.After.Access != nil {
			for _, after := range r.Ordering.After.Access {
				orderingAfter = append(orderingAfter, after.ValueString())
			}
			pluginInput.Ordering.After = &shared.After{
				Access: orderingAfter,
			}
		}

		if len(orderingBefore) > 0 || len(orderingAfter) > 0 {
			pluginInput.Ordering = &shared.Ordering{}
			if len(orderingBefore) > 0 {
				pluginInput.Ordering.Before = &shared.Before{
					Access: orderingBefore,
				}
			}
			if len(orderingAfter) > 0 {
				pluginInput.Ordering.After = &shared.After{
					Access: orderingAfter,
				}
			}
		}
	}

	return pluginInput, nil
}
