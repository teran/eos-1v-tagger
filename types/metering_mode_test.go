package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMeteringMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name         string
		input        string
		expOutput    *MeteringMode
		expEXIFValue EXIFValue
		expString    string
		expError     error
	}

	tcs := []testCase{
		{
			name:      "Evaluative",
			input:     "Evaluative",
			expOutput: PtrMeteringMode(MeteringModeEvaluative),
			expString: "Evaluative",
			expEXIFValue: EXIFValue{
				"ExifIFD:MeteringMode":                "Multi-Segment",
				"Canon:MeteringMode":                  "Evaluative",
				"CanonCustom:PF2DisableMeteringModes": "Off",
			},
		},
		{
			name:      "Partial",
			input:     "Partial",
			expOutput: PtrMeteringMode(MeteringModePartial),
			expString: "Partial",
			expEXIFValue: EXIFValue{
				"ExifIFD:MeteringMode":                "Partial",
				"Canon:MeteringMode":                  "Partial",
				"CanonCustom:PF2DisableMeteringModes": "Off",
			},
		},
		{
			name:      "Spot",
			input:     "Spot",
			expOutput: PtrMeteringMode(MeteringModeSpot),
			expString: "Spot",
			expEXIFValue: EXIFValue{
				"ExifIFD:MeteringMode":                "Spot",
				"Canon:MeteringMode":                  "Spot",
				"CanonCustom:PF2DisableMeteringModes": "Off",
			},
		},
		{
			name:      "Center Averaging",
			input:     "Center Averaging",
			expOutput: PtrMeteringMode(MeteringModeCenterAveraging),
			expString: "Center Averaging",
			expEXIFValue: EXIFValue{
				"ExifIFD:MeteringMode":                "Center-weighted average",
				"Canon:MeteringMode":                  "Center-weighted average",
				"CanonCustom:PF2DisableMeteringModes": "Off",
			},
		},
		{
			name:      "Center Averaging with spaces",
			input:     "      Center Averaging    ",
			expOutput: PtrMeteringMode(MeteringModeCenterAveraging),
			expString: "Center Averaging",
			expEXIFValue: EXIFValue{
				"ExifIFD:MeteringMode":                "Center-weighted average",
				"Canon:MeteringMode":                  "Center-weighted average",
				"CanonCustom:PF2DisableMeteringModes": "Off",
			},
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing MeteringMode: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("empty value"),
		},
	}

	for _, tc := range tcs {
		sm, err := MeteringModeFromString(tc.input)
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

func TestMeteringModeWithUnexpectedValue(t *testing.T) {
	r := require.New(t)

	v := MeteringMode("blah")

	r.Equal("blah", v.String())
	r.Equal(EXIFValue{
		"Canon:MeteringMode":                  "Unknown",
		"CanonCustom:PF2DisableMeteringModes": "Off",
		"ExifIFD:MeteringMode":                "Unknown",
	}, v.EXIFValue())
}
