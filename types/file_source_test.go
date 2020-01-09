package types

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestFileSourceViaSetter(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput FileSource
		expLayout string
		expString string
		expError  error
	}

	tcs := []testCase{
		{
			name:      "Digital Camera",
			input:     "Digital Camera",
			expOutput: FileSourceDigitalCamera,
			expString: "Digital Camera",
		},
		{
			name:      "Reflection Print Scanner",
			input:     "Reflection Print Scanner",
			expOutput: FileSourceReflectionPrintScanner,
			expString: "Reflection Print Scanner",
		},
		{
			name:      "Film Scanner",
			input:     "Film Scanner",
			expOutput: FileSourceFilmScanner,
			expString: "Film Scanner",
		},
		{
			name:     "undefined value",
			input:    "blah",
			expError: errors.New("Unknown value `blah` for time format"),
		},
	}

	for _, tc := range tcs {
		fs := new(FileSource)
		err := fs.Set(tc.input)
		if tc.expError == nil {
			r.NoError(err)
			r.Equalf(tc.expOutput, *fs, tc.name)
			r.Equalf(tc.expString, fs.String(), tc.name)
		} else {
			r.Error(err)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}
