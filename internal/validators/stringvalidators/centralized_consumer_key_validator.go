package stringvalidators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = StringCentralizedConsumerKeyValidatorValidator{}

type StringCentralizedConsumerKeyValidatorValidator struct{}

// Description describes the validation in plain text formatting.
func (v StringCentralizedConsumerKeyValidatorValidator) Description(_ context.Context) string {
	return "Validates that type == 'legacy' if the secret field is provided"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v StringCentralizedConsumerKeyValidatorValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v StringCentralizedConsumerKeyValidatorValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {

	// If attribute configuration is unknown, delay the validation until it is known.
	if req.ConfigValue.IsUnknown() {
		return
	}

	var typeValue string
	if err := req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName("type"), &typeValue); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Configuration",
			"Failed to retrieve the 'type' attribute.",
		)
		return
	}

	secretValue := req.ConfigValue.ValueString()
	if secretValue != "" && typeValue != "legacy" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Configuration",
			"The 'secret' field can only be set if 'type' is set to 'legacy'.",
		)
		return
	}

	if secretValue == "" && typeValue == "legacy" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Configuration",
			"The 'secret' field must be provided if 'type' is set to 'legacy'.",
		)
		return
	}
}

func CentralizedConsumerKeyValidator() validator.String {
	return StringCentralizedConsumerKeyValidatorValidator{}
}
