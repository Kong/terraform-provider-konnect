//go:build unit

package encodedstring

import (
	"context"
	"testing"
)

func TestBase64OrPlainInputValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                           string
		stateOrConfigValue             string
		currentValue                   string
		expectedSemanticEqualityOutput bool
	}{
		{
			name:                           "semantically equal values with encoded config/state value",
			stateOrConfigValue:             "c2FtcGxlIHZhbHVl",
			currentValue:                   "sample value",
			expectedSemanticEqualityOutput: true,
		},
		{
			name:                           "semantically unequal values",
			stateOrConfigValue:             "c2FtcGxlIHZhbHVl",
			currentValue:                   "not sample value",
			expectedSemanticEqualityOutput: false,
		},
		{
			name:                           "semantically equal plain values",
			stateOrConfigValue:             "#My document",
			currentValue:                   "#My document",
			expectedSemanticEqualityOutput: true,
		},
		{
			name:                           "semantically unequal plain values",
			stateOrConfigValue:             "#My document",
			currentValue:                   "#Not My document",
			expectedSemanticEqualityOutput: false,
		},
	}

	ctx := context.Background()

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			givenValue := NewBase64OrPlainInputValue(testCase.stateOrConfigValue)
			currentValue := NewBase64OrPlainInputValue(testCase.currentValue)

			areEqual, _ := currentValue.StringSemanticEquals(ctx, givenValue)

			if areEqual != testCase.expectedSemanticEqualityOutput {
				t.Errorf("Unexpected difference in Base64OrPlainInput semantic equality, got: %t, expected: %t", areEqual, testCase.expectedSemanticEqualityOutput)
			}
		})
	}
}
