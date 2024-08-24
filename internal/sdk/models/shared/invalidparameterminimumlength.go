// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Rule - invalid parameters rules
type Rule string

const (
	RuleMinLength    Rule = "min_length"
	RuleMinDigits    Rule = "min_digits"
	RuleMinLowercase Rule = "min_lowercase"
	RuleMinUppercase Rule = "min_uppercase"
	RuleMinSymbols   Rule = "min_symbols"
)

func (e Rule) ToPointer() *Rule {
	return &e
}
func (e *Rule) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "min_length":
		fallthrough
	case "min_digits":
		fallthrough
	case "min_lowercase":
		fallthrough
	case "min_uppercase":
		fallthrough
	case "min_symbols":
		*e = Rule(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Rule: %v", v)
	}
}

type InvalidParameterMinimumLength struct {
	Field string `json:"field"`
	// invalid parameters rules
	Rule    Rule    `json:"rule"`
	Minimum int64   `json:"minimum"`
	Source  *string `json:"source,omitempty"`
	Reason  string  `json:"reason"`
}

func (o *InvalidParameterMinimumLength) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *InvalidParameterMinimumLength) GetRule() Rule {
	if o == nil {
		return Rule("")
	}
	return o.Rule
}

func (o *InvalidParameterMinimumLength) GetMinimum() int64 {
	if o == nil {
		return 0
	}
	return o.Minimum
}

func (o *InvalidParameterMinimumLength) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *InvalidParameterMinimumLength) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}
