package types

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestTimestampFormat(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput TimestampFormat
		expLayout string
		expString string
		expError  error
	}

	tcs := []testCase{
		{
			name:      "timestmap format EU",
			input:     "EU",
			expOutput: TimestampFormatEU,
			expString: "EU",
			expLayout: "2/1/2006T15:04:05",
		},
		{
			name:      "timestmap format US",
			input:     "US",
			expOutput: TimestampFormatUS,
			expString: "US",
			expLayout: "1/2/2006T15:04:05",
		},
		{
			name:     "unexpected value",
			input:    "blah",
			expError: errors.New("Invalid value for timestamp format"),
		},
	}

	for _, tc := range tcs {
		tsf, err := NewTimestampFormat(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, *tsf, tc.name)
			r.Equalf(tc.expLayout, tsf.TimeLayout(), tc.name)
			r.Equalf(tc.expString, tsf.String(), tc.name)
		} else {
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}

func TestTimestampFormatViaSetter(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput TimestampFormat
		expLayout string
		expString string
		expError  error
	}

	tcs := []testCase{
		{
			name:      "timestmap format EU",
			input:     `"EU"`,
			expOutput: TimestampFormatEU,
			expString: "EU",
			expLayout: "2/1/2006T15:04:05",
		},
		{
			name:      "timestmap format US",
			input:     `"US"`,
			expOutput: TimestampFormatUS,
			expString: "US",
			expLayout: "1/2/2006T15:04:05",
		},
		{
			name:     "unexpected value",
			input:    `"blah"`,
			expError: errors.New("Unknown value `\"blah\"` for time format"),
		},
		{
			name:     "no quotes",
			input:    "blah",
			expError: errors.New("invalid syntax"),
		},
	}

	for _, tc := range tcs {
		tsf := new(TimestampFormat)
		err := tsf.Set(tc.input)
		if tc.expError == nil {
			r.NoError(err)
			r.Equalf(tc.expOutput, *tsf, tc.name)
			r.Equalf(tc.expLayout, tsf.TimeLayout(), tc.name)
			r.Equalf(tc.expString, tsf.String(), tc.name)
		} else {
			r.Error(err)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}
