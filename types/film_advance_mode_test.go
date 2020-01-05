package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilmAdvanceMode(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expOutput *FilmAdvanceMode
		expError  error
	}

	tcs := []testCase{
		{
			name:      "Single-frame",
			input:     "Single-frame",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceModeSingleFrame),
		},
		{
			name:      "Continuous (body only)",
			input:     "Continuous (body only)",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceModeContinuousBodyOnly),
		},
		{
			name:      "Low-speed continuous",
			input:     "Low-speed continuous",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceModeLowSpeedContinuous),
		},
		{
			name:      "High-speed continuous",
			input:     "High-speed continuous",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceModeHighSpeedContinuous),
		},
		{
			name:      "Ultra-high-speed continuous",
			input:     "Ultra-high-speed continuous",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceModeUltraHighSpeedContinuous),
		},
		{
			name:      "2-sec. self-timer",
			input:     "2-sec. self-timer",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceMode2secSelfTimer),
		},
		{
			name:      "10-sec. self-timer",
			input:     "10-sec. self-timer",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceMode10secSelfTimer),
		},
		{
			name:      "Ultra-high-speed continuous with spaces",
			input:     "      Ultra-high-speed continuous    ",
			expOutput: PtrFilmAdvanceMode(FilmAdvanceModeUltraHighSpeedContinuous),
		},
		{
			name:     "some random text",
			input:    "aksjdfghq3",
			expError: errors.New("error parsing FilmAdvanceMode: unknown value `aksjdfghq3`"),
		},
		{
			name:     "empty string",
			input:    "",
			expError: errors.New("empty value"),
		},
	}

	for _, tc := range tcs {
		sm, err := FilmAdvanceModeFromString(tc.input)
		if tc.expError == nil {
			r.Equalf(tc.expOutput, sm, tc.name)
			r.NoErrorf(err, tc.name)
		} else {
			r.Errorf(err, tc.name)
			r.Equalf(tc.expError.Error(), err.Error(), tc.name)
		}
	}
}
