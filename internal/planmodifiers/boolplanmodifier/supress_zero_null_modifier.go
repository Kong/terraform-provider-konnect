package boolplanmodifier

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Bool = BoolSupressZeroNullModifierPlanModifier{}

type BoolSupressZeroNullModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v BoolSupressZeroNullModifierPlanModifier) Description(_ context.Context) string {
    return "Treats null and false the same"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v BoolSupressZeroNullModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

// Validate performs the plan modification.
func (v BoolSupressZeroNullModifierPlanModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
    // req.State.Raw is null on the first request
    //println("{" + req.Path.String() +
    //    " - state: " + req.StateValue.String() +
    //    ", plan: " + req.PlanValue.String() +
    //    ", config: " + req.ConfigValue.String() +
    //    ", resp: " + resp.PlanValue.String() +
    //    "}")
}

func SupressZeroNullModifier() planmodifier.Bool {
    return BoolSupressZeroNullModifierPlanModifier{}
}
