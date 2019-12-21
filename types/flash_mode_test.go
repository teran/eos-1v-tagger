package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlashMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput *FlashMode
		expError  error
	}

	tcs := []testCase{
		{
			name:      "ON",
			input:     "ON",
			expOutput: ptrFlashMode(FlashModeOn),
		},
		{
			name:      "OFF",
			input:     "OFF",
			expOutput: ptrFlashMode(FlashModeOff),
		},
		{
			name:      "E-TTL",
			input:     "E-TTL",
			expOutput: ptrFlashMode(FlashModeETTL),
		},
		{
			name:      "A-TTL",
			input:     "A-TTL",
			expOutput: ptrFlashMode(FlashModeATTL),
		},
		{
			name:      "TTL autoflash",
			input:     "TTL autoflash",
			expOutput: ptrFlashMode(FlashModeTTLAutoflash),
		},
		{
			name:      "Manual flash",
			input:     "Manual flash",
			expOutput: ptrFlashMode(FlashModeManualFlash),
		},
		{
			name:      "TTL autoflash with spaces",
			input:     "      TTL autoflash    ",
			expOutput: ptrFlashMode(FlashModeTTLAutoflash),
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing FlashMode: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("error parsing FlashMode: unknown value ``"),
		},
	}

	for _, tc := range tcs {
		sm, err := FlashModeFromString(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, sm, tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}

func ptrFlashMode(fm FlashMode) *FlashMode {
	if fm == "" {
		return nil
	}
	return &fm
}
