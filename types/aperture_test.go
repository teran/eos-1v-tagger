package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAperture(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name         string
		input        string
		expOutput    *Aperture
		expEXIFValue EXIFValue
		expError     error
	}

	tcs := []testCase{
		{
			name:      "valid aperture value",
			input:     "2",
			expOutput: PtrAperture(2),
			expEXIFValue: EXIFValue{
				"Canon:FNumber":         "2.0",
				"Canon:TargetAperture":  "2.0",
				"ExifIFD:ApertureValue": "2.0",
				"ExifIFD:FNumber":       "2.0",
			},
		},
		{
			name:      "valid aperture value with spaces",
			input:     "      3.2    ",
			expOutput: PtrAperture(3.2),
			expEXIFValue: EXIFValue{
				"Canon:FNumber":         "3.2",
				"Canon:TargetAperture":  "3.2",
				"ExifIFD:ApertureValue": "3.2",
				"ExifIFD:FNumber":       "3.2",
			},
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("strconv.ParseFloat: parsing \"aksjdfghq3\": invalid syntax"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("empty value"),
		},
	}

	for _, tc := range tcs {
		sm, err := ApertureFromString(tc.input)
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
