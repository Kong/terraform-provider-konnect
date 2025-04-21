package stringplanmodifier

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type StringBase64InputEqualityPlanModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (v StringBase64InputEqualityPlanModifier) Description(_ context.Context) string {
	return "If the value in plan in base64 encoded, decodes it and compares it to state. If equal, suppresses the diff."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (v StringBase64InputEqualityPlanModifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Implements the plan modification logic
func (v StringBase64InputEqualityPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() || req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	planStr := req.PlanValue.ValueString()
	stateStr := req.StateValue.ValueString()

	if planStr == stateStr {
		// They are equal, do nothing.
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(planStr)
	if err != nil {
		// Not base64 encoded, do nothing.
		return
	}

	if strings.TrimSpace(string(decoded)) == strings.TrimSpace(stateStr) {
		// decoded value from plan is same as state - so override planned value to suppress diff
		resp.PlanValue = req.StateValue
	}
}

func Base64InputEquality() planmodifier.String {
	return StringBase64InputEqualityPlanModifier{}
}
