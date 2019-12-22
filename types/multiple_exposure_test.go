package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultipleExposure(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput *MultipleExposure
		expError  error
	}

	tcs := []testCase{
		{
			name:      "ON",
			input:     "ON",
			expOutput: ptrMultipleExposure(MultipleExposureOn),
		},
		{
			name:      "OFF",
			input:     "OFF",
			expOutput: ptrMultipleExposure(MultipleExposureOff),
		},
		{
			name:      "OFF with spaces",
			input:     "      OFF    ",
			expOutput: ptrMultipleExposure(MultipleExposureOff),
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing MultipleExposure: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("error parsing MultipleExposure: unknown value ``"),
		},
	}

	for _, tc := range tcs {
		sm, err := MultipleExposureFromString(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, sm, tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}

func ptrMultipleExposure(me MultipleExposure) *MultipleExposure {
	if me == "" {
		return nil
	}
	return &me
}
