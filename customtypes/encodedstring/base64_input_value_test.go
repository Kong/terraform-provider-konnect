//go:build unit

package encodedstring

import (
	"context"
	"testing"
)

func TestBase64InputValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                           string
		stateOrConfigValue             string
		currentValue                   string
		expectedSemanticEqualityOutput bool
	}{
		{
			name:                           "semantically equal values",
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
	}

	ctx := context.Background()

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			givenValue := NewBase64InputValue(testCase.stateOrConfigValue)
			currentValue := NewBase64InputValue(testCase.currentValue)

			areEqual, _ := currentValue.StringSemanticEquals(ctx, givenValue)

			if areEqual != testCase.expectedSemanticEqualityOutput {
				t.Errorf("Unexpected difference in Base64Input semantic equality, got: %t, expected: %t", areEqual, testCase.expectedSemanticEqualityOutput)
			}
		})
	}
}
