package tagger

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCSVParser(t *testing.T) {
	r := require.New(t)

	tz, err := LocationByTimeZone("CET")
	r.NoError(err)

	p, err := NewCSVParser("testdata/sample.CSV", tz)
	r.NoError(err)
	r.NotNil(p)

	defer func() {
		err := p.Close()
		r.NoError(err)
	}()

	film, err := p.Parse()
	r.NoError(err)
	r.Equal(Film{
		ID:                  "03-758",
		Title:               "Sample",
		FilmLoadedTimestamp: mustParseTimestamp(t, "09/01/2010T14:00:00", tz),
		FrameCount:          36,
		ISO:                 200,
		Frames: []Frame{
			{
				Flag:              true,
				Number:            1,
				FocalLength:       11,
				MaxAperture:       22,
				Tv:                "1/8000",
				Av:                22,
				ISO:               200,
				FlashMode:         "OFF",
				MeteringMode:      "Evaluative",
				ShootingMode:      "Program AE",
				FilmAdvanceMode:   "Single-frame",
				AFMode:            "One-Shot AF",
				Timestamp:         mustParseTimestamp(t, "11/09/2010T18:31:26", tz),
				MultipleExposure:  "OFF",
				BatteryLoadedDate: time.Time{},
			},
			{
				Flag:                 true,
				Number:               2,
				FocalLength:          24,
				MaxAperture:          2.8,
				Tv:                   "15",
				Av:                   1.4,
				ISO:                  200,
				FlashMode:            "OFF",
				MeteringMode:         "Evaluative",
				ShootingMode:         "Program AE",
				FilmAdvanceMode:      "Single-frame",
				AFMode:               "One-Shot AF",
				Timestamp:            mustParseTimestamp(t, "12/09/2010T18:32:55", tz),
				MultipleExposure:     "OFF",
				BatteryLoadedDate:    time.Time{},
				ExposureCompensation: -5,
				FlashCompensation:    -4.5,
			},
		},
	}, film)
}

func TestIsFilmHeader(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expResult bool
	}

	tcs := []testCase{
		{
			name:      "normal film header",
			input:     `*,Film ID,03-758,Title,Sample,Date and time film loaded,9/1/2010,14:00:00,Frame count,36,ISO (DX),200`,
			expResult: true,
		},
		{
			name:      "normal film header with no asterisk",
			input:     `,Film ID,03-758,Title,Sample,Date and time film loaded,9/1/2010,14:00:00,Frame count,36,ISO (DX),200`,
			expResult: true,
		},
		{
			name:      "not a film header",
			input:     `blah`,
			expResult: false,
		},
		{
			name:      "empty line",
			input:     `blah`,
			expResult: false,
		},
	}

	for _, tc := range tcs {
		result := isFilmHeader(tc.input)
		r.Equalf(tc.expResult, result, tc.name)
	}
}

func TestIsFrameHeader(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     string
		expResult bool
	}

	tcs := []testCase{
		{
			name:      "normal film header",
			input:     `*,Frame No.,Focal length,Max. aperture,Tv,Av,ISO (M),Exposure compensation,Flash exposure compensation,Flash mode,Metering mode,Shooting mode,Film advance mode,AF mode,Bulb exposure time,Date,Time,Multiple exposure,Battery-loaded date,Battery-loaded time,Remarks`,
			expResult: true,
		},
		{
			name:      "normal film header with no asterisk",
			input:     `,Frame No.,Focal length,Max. aperture,Tv,Av,ISO (M),Exposure compensation,Flash exposure compensation,Flash mode,Metering mode,Shooting mode,Film advance mode,AF mode,Bulb exposure time,Date,Time,Multiple exposure,Battery-loaded date,Battery-loaded time,Remarks`,
			expResult: true,
		},
		{
			name:      "not a film header",
			input:     `blah`,
			expResult: false,
		},
		{
			name:      "empty line",
			input:     `blah`,
			expResult: false,
		},
	}

	for _, tc := range tcs {
		result := isFrameHeader(tc.input)
		r.Equalf(tc.expResult, result, tc.name)
	}
}

func mustParseTimestamp(t *testing.T, ts string, tz *time.Location) time.Time {
	r := require.New(t)

	tts := strings.Split(ts, "T")
	r.Len(tts, 2)

	tt := maybeParseTimestamp(tts[0], tts[1], tz)
	r.NotNil(tt)
	r.False(tt.IsZero())

	return tt
}
