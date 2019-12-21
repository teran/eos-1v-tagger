package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAFMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput *AFMode
		expError  error
	}

	tcs := []testCase{
		{
			name:      "One-Shot AF",
			input:     "One-Shot AF",
			expOutput: ptrAFMode(AFModeOneShotAF),
		},
		{
			name:      "AI Servo AF",
			input:     "AI Servo AF",
			expOutput: ptrAFMode(AFModeAIServoAF),
		},
		{
			name:      "Manual focus",
			input:     "Manual focus",
			expOutput: ptrAFMode(AFModeManualFocus),
		},
		{
			name:      "Manual focus with spaces",
			input:     "      Manual focus    ",
			expOutput: ptrAFMode(AFModeManualFocus),
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing AFMode: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("error parsing AFMode: unknown value ``"),
		},
	}

	for _, tc := range tcs {
		sm, err := AFModeFromString(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, sm, tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}

func ptrAFMode(am AFMode) *AFMode {
	if am == "" {
		return nil
	}
	return &am
}