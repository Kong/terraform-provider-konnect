package stringplanmodifier

import (
    "context"
    "encoding/json"
    "fmt"
    "reflect"

    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = StringArbitraryJSONModifierPlanModifier{}

type StringArbitraryJSONModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v StringArbitraryJSONModifierPlanModifier) Description(_ context.Context) string {
    return "Uses deep equal to check if the underlying JSON object changed."
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v StringArbitraryJSONModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

// Validate performs the plan modification.
func (v StringArbitraryJSONModifierPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
    if !req.StateValue.IsNull() {
        equal, err := compareJSON(resp.PlanValue.ValueString(), req.StateValue.ValueString())
        if err != nil {
            resp.Diagnostics.AddAttributeError(
                req.Path,
                "must be a valid json",
                req.Path.String()+": "+v.Description(ctx),
            )
            fmt.Println("Error comparing JSONs:", err)
            return
        }

        if equal {
            resp.PlanValue = req.StateValue
        }
    }
}

func ArbitraryJSONModifier() planmodifier.String {
    return StringArbitraryJSONModifierPlanModifier{}
}

// compareJSON takes two JSON strings, parses them into interface{} values,
// and compares whether they represent equivalent data structures.
func compareJSON(jsonStr1, jsonStr2 string) (bool, error) {
    var obj1, obj2 interface{}

    // Parse the first JSON string.
    if err := json.Unmarshal([]byte(jsonStr1), &obj1); err != nil {
        return false, fmt.Errorf("error parsing first JSON: %w", err)
    }

    // Parse the second JSON string.
    if err := json.Unmarshal([]byte(jsonStr2), &obj2); err != nil {
        return false, fmt.Errorf("error parsing second JSON: %w", err)
    }

    // Use reflect.DeepEqual to compare the parsed JSON objects.
    return reflect.DeepEqual(obj1, obj2), nil
}
