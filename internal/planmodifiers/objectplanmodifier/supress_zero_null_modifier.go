package objectplanmodifier

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/types/basetypes"

    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Object = ObjectSupressZeroNullModifierPlanModifier{}

type ObjectSupressZeroNullModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v ObjectSupressZeroNullModifierPlanModifier) Description(_ context.Context) string {
    return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v ObjectSupressZeroNullModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

// Validate performs the plan modification.
func (v ObjectSupressZeroNullModifierPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
    if resp.PlanValue.IsNull() && allNestedAttributesAreNull(req.StateValue) {
        resp.PlanValue = req.StateValue
    }
}

func SupressZeroNullModifier() planmodifier.Object {
    return ObjectSupressZeroNullModifierPlanModifier{}
}

// allNestedAttributesAreNull recursively checks whether every attribute in the given types.Object
// (including any nested types.Object attributes) is null.
func allNestedAttributesAreNull(obj basetypes.ObjectValue) bool {
    // If the entire object is null, then all its attributes are considered null.
    if obj.IsNull() {
        return true
    }

    // The Object type (an alias to basetypes.ObjectValue) typically provides access
    // to its attributes via an Attributes (or similar) field.
    // Here we assume the field is named "Attributes" of type map[string]attr.Value.
    // (If your actual type uses a different field or accessor, adjust accordingly.)
    for _, val := range obj.Attributes() {
        // If the attribute itself is null, then we're fine with that attribute.
        if val.IsNull() {
            continue
        }

        // If the attribute is not null, check if it is itself an object.
        // If so, recursively inspect its attributes.
        if nestedObj, ok := val.(basetypes.ObjectValue); ok {
            if !allNestedAttributesAreNull(nestedObj) {
                // If any nested object contains a non-null attribute, then return false.
                return false
            }
        } else {
            // For any attribute that is not an object and is not null,
            // we conclude that not all attributes are null.
            return false
        }

        // (Optional) You might log the key/value here for debugging:
        // log.Printf("Attribute %q is null or recursively null", key)
    }

    // If we finish iterating without finding a non-null attribute,
    // then all attributes are null.
    return true
}
