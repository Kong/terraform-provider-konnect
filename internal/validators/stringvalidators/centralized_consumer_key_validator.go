package stringvalidators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = StringCentralizedConsumerKeyValidatorValidator{}

type StringCentralizedConsumerKeyValidatorValidator struct{}

// Description describes the validation in plain text formatting.
func (v StringCentralizedConsumerKeyValidatorValidator) Description(_ context.Context) string {
	return "TODO: add validator description"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v StringCentralizedConsumerKeyValidatorValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v StringCentralizedConsumerKeyValidatorValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"TODO: implement stringvalidator CentralizedConsumerKeyValidator logic",
		req.Path.String()+": "+v.Description(ctx),
	)
}

func CentralizedConsumerKeyValidator() validator.String {
	return StringCentralizedConsumerKeyValidatorValidator{}
}
