package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMeteringMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput *MeteringMode
		expError  error
	}

	tcs := []testCase{
		{
			name:      "Evaluative",
			input:     "Evaluative",
			expOutput: ptrMeteringMode(MeteringModeEvaluative),
		},
		{
			name:      "Partial",
			input:     "Partial",
			expOutput: ptrMeteringMode(MeteringModePartial),
		},
		{
			name:      "Spot",
			input:     "Spot",
			expOutput: ptrMeteringMode(MeteringModeSpot),
		},
		{
			name:      "Center Averaging",
			input:     "Center Averaging",
			expOutput: ptrMeteringMode(MeteringModeCenterAveraging),
		},
		{
			name:      "Center Averaging with spaces",
			input:     "      Center Averaging    ",
			expOutput: ptrMeteringMode(MeteringModeCenterAveraging),
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing MeteringMode: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("error parsing MeteringMode: unknown value ``"),
		},
	}

	for _, tc := range tcs {
		sm, err := MeteringModeFromString(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, sm, tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}

func ptrMeteringMode(mm MeteringMode) *MeteringMode {
	if mm == "" {
		return nil
	}
	return &mm
}
