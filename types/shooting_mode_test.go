package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShootingMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name         string
		input        string
		expOutput    *ShootingMode
		expEXIFValue EXIFValue
		expString    string
		expError     error
	}

	tcs := []testCase{
		{
			name:      "Program AE",
			input:     "Program AE",
			expOutput: PtrShootingMode(ShootingModeProgramAE),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Program AE",
				"Canon:CanonExposureMode":             "Program AE",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Program AE",
		},
		{
			name:      "Shutter-speed-priority AE",
			input:     "Shutter-speed-priority AE",
			expOutput: PtrShootingMode(ShootingModeShutterSpeedPriorityAE),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Shutter speed priority AE",
				"Canon:CanonExposureMode":             "Shutter speed priority AE",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Shutter-speed-priority AE",
		},
		{
			name:      "Aperture-priority AE",
			input:     "Aperture-priority AE",
			expOutput: PtrShootingMode(ShootingModeAperturePriorityAE),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Aperture-priority AE",
				"Canon:CanonExposureMode":             "Aperture-priority AE",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Aperture-priority AE",
		},
		{
			name:      "Depth-of-field AE",
			input:     "Depth-of-field AE",
			expOutput: PtrShootingMode(ShootingModeDepthOfFieldAE),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Not Defined",
				"Canon:CanonExposureMode":             "Depth-of-field AE",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Depth-of-field AE",
		},
		{
			name:      "Manual exposure",
			input:     "Manual exposure",
			expOutput: PtrShootingMode(ShootingModeManualExposure),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Manual",
				"Canon:CanonExposureMode":             "Manual",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Manual exposure",
		},
		{
			name:      "Bulb",
			input:     "Bulb",
			expOutput: PtrShootingMode(ShootingModeBulb),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Bulb",
				"Canon:CanonExposureMode":             "Bulb",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Bulb",
		},
		{
			name:      "Program AE with spaces",
			input:     "      Program AE    ",
			expOutput: PtrShootingMode(ShootingModeProgramAE),
			expEXIFValue: EXIFValue{
				"ExifIFD:ExposureProgram":             "Program AE",
				"Canon:CanonExposureMode":             "Program AE",
				"CanonCustom:PF1DisableShootingModes": "Off",
				"CanonCustom:PF6PresetShootingModes":  "Off",
			},
			expString: "Program AE",
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
			r.Equalf(tc.expString, sm.String(), tc.name)
			r.Equalf(tc.expEXIFValue, sm.EXIFValue(), tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}
