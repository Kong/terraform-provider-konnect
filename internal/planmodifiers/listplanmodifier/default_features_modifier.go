package listplanmodifier

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.List = ListDefaultFeaturesModifierPlanModifier{}

type ListDefaultFeaturesModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v ListDefaultFeaturesModifierPlanModifier) Description(_ context.Context) string {
    return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v ListDefaultFeaturesModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

// Validate performs the plan modification.
func (v ListDefaultFeaturesModifierPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
    //if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
    //    var l = types.List{}
    //    diags := tfsdk.ValueFrom(ctx, defaultFeatures, attr.TypeWithAttributeTypes{}, &l)
    //
    //    if diags.HasError() {
    //        resp.Diagnostics.AddAttributeError(
    //            req.Path,
    //            "Problems converting",
    //            req.Path.String()+": "+v.Description(ctx),
    //        )
    //    }
    //
    //    resp.PlanValue = l
    //}
}

func DefaultFeaturesModifier() planmodifier.List {
    return ListDefaultFeaturesModifierPlanModifier{}
}
