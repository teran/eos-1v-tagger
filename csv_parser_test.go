package tagger

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTwoFilmsInSingleCSV(t *testing.T) {
	r := require.New(t)

	tz, err := LocationByTimeZone("CET")
	r.NoError(err)

	p, err := NewCSVParser("testdata/two-films.csv", tz, TimestampFormatUS)
	r.NoError(err)
	r.NotNil(p)

	defer func() {
		err := p.Close()
		r.NoError(err)
	}()

	film, err := p.Parse()
	r.NoError(err)
	r.Equal([]Film{
		{
			ID:                  139,
			CameraID:            1,
			Title:               "SampleTest film #139",
			FilmLoadedTimestamp: mustParseTimestamp(t, "09/28/2019T10:21:32", tz, TimestampFormatUS),
			FrameCount:          2,
			ISO:                 400,
			Remarks:             "test remarks data",
			Frames: []*Frame{
				{
					Flag:              false,
					Number:            1,
					FocalLength:       24,
					MaxAperture:       1.4,
					Tv:                "1/40",
					Av:                1.4,
					ISO:               400,
					FlashMode:         "OFF",
					MeteringMode:      "Evaluative",
					ShootingMode:      "Aperture-priority AE",
					FilmAdvanceMode:   "Single-frame",
					AFMode:            "One-Shot AF",
					Timestamp:         mustParseTimestamp(t, "10/7/2019T20:02:18", tz, TimestampFormatUS),
					MultipleExposure:  "OFF",
					BatteryLoadedDate: time.Time{},
					Remarks:           "test frame #1",
				},
				{
					Flag:                 true,
					Number:               2,
					FocalLength:          35,
					MaxAperture:          1.4,
					Tv:                   "1/60",
					Av:                   1.4,
					ISO:                  400,
					FlashMode:            "OFF",
					MeteringMode:         "Evaluative",
					ShootingMode:         "Aperture-priority AE",
					FilmAdvanceMode:      "Single-frame",
					AFMode:               "One-Shot AF",
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:29", tz, TimestampFormatUS),
					MultipleExposure:     "OFF",
					BatteryLoadedDate:    time.Time{},
					ExposureCompensation: -5,
					FlashCompensation:    -4.5,
					Remarks:              "test frame #2",
				},
			},
		},
		{
			ID:                  140,
			CameraID:            1,
			Title:               "SampleTest film #139 part II",
			FilmLoadedTimestamp: mustParseTimestamp(t, "10/07/2019T22:55:58", tz, TimestampFormatUS),
			FrameCount:          2,
			ISO:                 400,
			Remarks:             "test remarks data 2",
			Frames: []*Frame{
				{
					Flag:                 false,
					Number:               1,
					FocalLength:          14,
					MaxAperture:          1.4,
					Tv:                   "1/1600",
					Av:                   1.4,
					ISO:                  200,
					FlashMode:            "OFF",
					MeteringMode:         "Evaluative",
					ShootingMode:         "Program AE",
					FilmAdvanceMode:      "Single-frame",
					AFMode:               "One-Shot AF",
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:38", tz, TimestampFormatUS),
					MultipleExposure:     "OFF",
					BatteryLoadedDate:    time.Time{},
					ExposureCompensation: 1,
					FlashCompensation:    2,
					Remarks:              "test frame remarks #1",
				},
				{
					Flag:                 true,
					Number:               2,
					FocalLength:          16,
					MaxAperture:          1.4,
					Tv:                   "1/1250",
					Av:                   1.4,
					ISO:                  800,
					FlashMode:            "OFF",
					MeteringMode:         "Evaluative",
					ShootingMode:         "Aperture-priority AE",
					FilmAdvanceMode:      "Single-frame",
					AFMode:               "One-Shot AF",
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:55", tz, TimestampFormatUS),
					MultipleExposure:     "OFF",
					BatteryLoadedDate:    time.Time{},
					ExposureCompensation: -1,
					FlashCompensation:    -2,
					Remarks:              "test frame remarks #2",
				},
			},
		},
	}, film)
}

func TestFilmWithPartialTimestampsEUFormatted(t *testing.T) {
	r := require.New(t)

	tz, err := LocationByTimeZone("CET")
	r.NoError(err)

	p, err := NewCSVParser("testdata/film-with-partial-timestamps-eu.csv", tz, TimestampFormatEU)
	r.NoError(err)
	r.NotNil(p)

	defer func() {
		err := p.Close()
		r.NoError(err)
	}()

	film, err := p.Parse()
	r.NoError(err)
	r.Equal([]Film{
		{
			ID:                  139,
			CameraID:            1,
			Title:               "SampleTest film #139",
			FilmLoadedTimestamp: mustParseTimestamp(t, "28/09/2019T10:21:32", tz, TimestampFormatEU),
			FrameCount:          2,
			ISO:                 400,
			Remarks:             "test remarks data",
			Frames: []*Frame{
				{
					Flag:              false,
					Number:            1,
					FocalLength:       24,
					MaxAperture:       1.4,
					Tv:                "1/40",
					Av:                1.4,
					ISO:               400,
					FlashMode:         "OFF",
					MeteringMode:      "Evaluative",
					ShootingMode:      "Aperture-priority AE",
					FilmAdvanceMode:   "Single-frame",
					AFMode:            "One-Shot AF",
					Timestamp:         mustParseTimestamp(t, "15/07/2019T20:02:18", tz, TimestampFormatEU),
					MultipleExposure:  "OFF",
					BatteryLoadedDate: time.Time{},
					Remarks:           "test frame #1",
				},
				{
					Flag:                 true,
					Number:               2,
					FocalLength:          35,
					MaxAperture:          1.4,
					Tv:                   "1/60",
					Av:                   1.4,
					ISO:                  400,
					FlashMode:            "OFF",
					MeteringMode:         "Evaluative",
					ShootingMode:         "Aperture-priority AE",
					FilmAdvanceMode:      "Single-frame",
					AFMode:               "One-Shot AF",
					Timestamp:            time.Time{},
					MultipleExposure:     "OFF",
					BatteryLoadedDate:    time.Time{},
					ExposureCompensation: -5,
					FlashCompensation:    -4.5,
					Remarks:              "test frame #2",
				},
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

func mustParseTimestamp(t *testing.T, ts string, tz *time.Location, tf string) time.Time {
	r := require.New(t)

	tts := strings.Split(ts, "T")
	r.Len(tts, 2)

	tt, err := parseTimestamp(tts[0], tts[1], tz, tf)
	r.NoError(err)

	return tt
}
