package listplanmodifier

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.List = ListSupressZeroNullModifierPlanModifier{}

type ListSupressZeroNullModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v ListSupressZeroNullModifierPlanModifier) Description(_ context.Context) string {
    return "Treats null and empty arrays the same"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v ListSupressZeroNullModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

// Validate performs the plan modification.
func (v ListSupressZeroNullModifierPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
    if len(resp.PlanValue.Elements()) == 0 && len(req.StateValue.Elements()) == 0 && req.ConfigValue.IsNull() {
        resp.PlanValue = req.StateValue
    }
    //resp.Diagnostics.AddAttributeError(
    //	req.Path,
    //	"TODO: implement plan modifier SupressZeroNullModifier logic",
    //	req.Path.String()+": "+v.Description(ctx),
    //)
}

func SupressZeroNullModifier() planmodifier.List {
    return ListSupressZeroNullModifierPlanModifier{}
}
