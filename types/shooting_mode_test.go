package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShootingMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput *ShootingMode
		expError  error
	}

	tcs := []testCase{
		{
			name:      "Program AE",
			input:     "Program AE",
			expOutput: ptrShootingMode(ShootingModeProgramAE),
		},
		{
			name:      "Shutter-speed-priority AE",
			input:     "Shutter-speed-priority AE",
			expOutput: ptrShootingMode(ShootingModeShutterSpeedPriorityAE),
		},
		{
			name:      "Aperture-priority AE",
			input:     "Aperture-priority AE",
			expOutput: ptrShootingMode(ShootingModeAperturePriorityAE),
		},
		{
			name:      "Depth-of-field AE",
			input:     "Depth-of-field AE",
			expOutput: ptrShootingMode(ShootingModeDepthOfFieldAE),
		},
		{
			name:      "Manual exposure",
			input:     "Manual exposure",
			expOutput: ptrShootingMode(ShootingModeManualExposure),
		},
		{
			name:      "Bulb",
			input:     "Bulb",
			expOutput: ptrShootingMode(ShootingModeBulb),
		},
		{
			name:      "Program AE with spaces",
			input:     "      Program AE    ",
			expOutput: ptrShootingMode(ShootingModeProgramAE),
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing ShootingMode: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("error parsing ShootingMode: unknown value ``"),
		},
	}

	for _, tc := range tcs {
		sm, err := ShootingModeFromString(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, sm, tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}

func ptrShootingMode(sm ShootingMode) *ShootingMode {
	if sm == "" {
		return nil
	}
	return &sm
}
