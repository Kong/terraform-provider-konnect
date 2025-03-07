package stringplanmodifier

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = StringRequiresReplaceModifierPlanModifier{}

type StringRequiresReplaceModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v StringRequiresReplaceModifierPlanModifier) Description(_ context.Context) string {
    return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v StringRequiresReplaceModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

// Validate performs the plan modification.
func (v StringRequiresReplaceModifierPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
    resp.RequiresReplace = true
}

func RequiresReplaceModifier() planmodifier.String {
    return StringRequiresReplaceModifierPlanModifier{}
}
