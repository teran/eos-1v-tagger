package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlashMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name         string
		input        string
		expOutput    *FlashMode
		expEXIFValue EXIFValue
		expError     error
	}

	tcs := []testCase{
		{
			name:      "ON",
			input:     "ON",
			expOutput: ptrFlashMode(FlashModeOn),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "On",
				"Canon:FlashBits":      "(none)",
				"ExifIFD:Flash":        "On, Fired",
			},
		},
		{
			name:      "OFF",
			input:     "OFF",
			expOutput: ptrFlashMode(FlashModeOff),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "Off",
				"Canon:FlashBits":      "(none)",
				"ExifIFD:Flash":        "Off, Did not fire",
			},
		},
		{
			name:      "E-TTL",
			input:     "E-TTL",
			expOutput: ptrFlashMode(FlashModeETTL),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "On",
				"Canon:FlashBits":      "E-TTL",
				"ExifIFD:Flash":        "Auto, Fired",
			},
		},
		{
			name:      "A-TTL",
			input:     "A-TTL",
			expOutput: ptrFlashMode(FlashModeATTL),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "On",
				"Canon:FlashBits":      "A-TTL",
				"ExifIFD:Flash":        "Auto, Fired",
			},
		},
		{
			name:      "TTL autoflash",
			input:     "TTL autoflash",
			expOutput: ptrFlashMode(FlashModeTTLAutoflash),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "Auto",
				"Canon:FlashBits":      "TTL",
				"ExifIFD:Flash":        "Auto, Fired",
			},
		},
		{
			name:      "Manual flash",
			input:     "Manual flash",
			expOutput: ptrFlashMode(FlashModeManualFlash),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "On",
				"Canon:FlashBits":      "Manual",
				"ExifIFD:Flash":        "On, Fired",
			},
		},
		{
			name:      "TTL autoflash with spaces",
			input:     "      TTL autoflash    ",
			expOutput: ptrFlashMode(FlashModeTTLAutoflash),
			expEXIFValue: EXIFValue{
				"Canon:CanonFlashMode": "Auto",
				"Canon:FlashBits":      "TTL",
				"ExifIFD:Flash":        "Auto, Fired",
			},
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
			r.Equalf(tc.expEXIFValue, sm.EXIFValue(), tc.name)
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
