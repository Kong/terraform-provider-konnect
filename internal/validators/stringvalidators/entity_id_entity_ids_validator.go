package stringvalidators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = StringEntityIdEntityIdsValidator{}

type StringEntityIdEntityIdsValidator struct{}

// Description describes the validation in plain text formatting.
func (v StringEntityIdEntityIdsValidator) Description(_ context.Context) string {
	return "entity_id and entity_ids are mutually exclusive"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v StringEntityIdEntityIdsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v StringEntityIdEntityIdsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the configured value is unknown or null there is nothing to compare, so
	// return early. The framework calls this validator again once the value is
	// known, so the check still runs in the apply for a value that comes from a
	// resource created in the same apply.
	//
	// The gap that remains is the diagnostic's phase, not a missed check: for an
	// entity_id that is unknown at plan time the error appears during the apply,
	// after any resources the id depends on were created. The schema description
	// says "the provider rejects" rather than "rejected at plan time" for that
	// reason. No configuration tested sent both fields to the API: the API's 400
	// is a backstop for callers that are not this provider.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	var entityIds types.List
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName("entity_ids"), &entityIds)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if entityIds.IsNull() || entityIds.IsUnknown() || len(entityIds.Elements()) == 0 {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Conflicting Configuration",
		"The 'entity_id' and 'entity_ids' fields are mutually exclusive. Only one of them can be set.",
	)
}

func EntityIDEntityIdsValidator() validator.String {
	return StringEntityIdEntityIdsValidator{}
}
