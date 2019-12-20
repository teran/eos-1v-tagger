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
	r.Equal([]*Film{
		{
			ID:                  ptrInt64(139),
			CameraID:            ptrInt64(1),
			Title:               ptrString("SampleTest film #139"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "09/28/2019T10:21:32", tz, TimestampFormatUS),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data"),
			Frames: []*Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					FocalLength:          ptrInt64(24),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/40"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(400),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrString("Evaluative"),
					ShootingMode:         ptrString("Aperture-priority AE"),
					ExposureCompensation: ptrFloat64(0),
					FlashCompensation:    ptrFloat64(0),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrString("One-Shot AF"),
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:18", tz, TimestampFormatUS),
					MultipleExposure:     ptrString("OFF"),
					BatteryLoadedDate:    nil,
					Remarks:              ptrString("test frame #1"),
				},
				{
					Flag:                 ptrBool(true),
					Number:               ptrInt64(2),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/60"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(400),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrString("Evaluative"),
					ShootingMode:         ptrString("Aperture-priority AE"),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrString("One-Shot AF"),
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:29", tz, TimestampFormatUS),
					MultipleExposure:     ptrString("OFF"),
					BatteryLoadedDate:    nil,
					ExposureCompensation: ptrFloat64(-5),
					FlashCompensation:    ptrFloat64(-4.5),
					Remarks:              ptrString("test frame #2"),
				},
			},
		},
		{
			ID:                  ptrInt64(140),
			CameraID:            ptrInt64(1),
			Title:               ptrString("SampleTest film #139 part II"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "10/07/2019T22:55:58", tz, TimestampFormatUS),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data 2"),
			Frames: []*Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					FocalLength:          ptrInt64(14),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/1600"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(200),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrString("Evaluative"),
					ShootingMode:         ptrString("Program AE"),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrString("One-Shot AF"),
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:38", tz, TimestampFormatUS),
					MultipleExposure:     ptrString("OFF"),
					BatteryLoadedDate:    nil,
					ExposureCompensation: ptrFloat64(1),
					FlashCompensation:    ptrFloat64(2),
					Remarks:              ptrString("test frame remarks #1"),
				},
				{
					Flag:                 ptrBool(true),
					Number:               ptrInt64(2),
					FocalLength:          ptrInt64(16),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/1250"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(800),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrString("Evaluative"),
					ShootingMode:         ptrString("Aperture-priority AE"),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrString("One-Shot AF"),
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:55", tz, TimestampFormatUS),
					MultipleExposure:     ptrString("OFF"),
					BatteryLoadedDate:    nil,
					ExposureCompensation: ptrFloat64(-1),
					FlashCompensation:    ptrFloat64(-2),
					Remarks:              ptrString("test frame remarks #2"),
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
	r.Equal([]*Film{
		{
			ID:                  ptrInt64(139),
			CameraID:            ptrInt64(1),
			Title:               ptrString("SampleTest film #139"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "28/09/2019T10:21:32", tz, TimestampFormatEU),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data"),
			Frames: []*Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					FocalLength:          ptrInt64(24),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/40"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(400),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrString("Evaluative"),
					ExposureCompensation: ptrFloat64(0),
					FlashCompensation:    ptrFloat64(0),
					ShootingMode:         ptrString("Aperture-priority AE"),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrString("One-Shot AF"),
					Timestamp:            mustParseTimestamp(t, "15/07/2019T20:02:18", tz, TimestampFormatEU),
					MultipleExposure:     ptrString("OFF"),
					BatteryLoadedDate:    nil,
					Remarks:              ptrString("test frame #1"),
				},
				{
					Flag:                 ptrBool(true),
					Number:               ptrInt64(2),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/60"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(400),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrString("Evaluative"),
					ShootingMode:         ptrString("Aperture-priority AE"),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrString("One-Shot AF"),
					Timestamp:            nil,
					MultipleExposure:     ptrString("OFF"),
					BatteryLoadedDate:    nil,
					ExposureCompensation: ptrFloat64(-5),
					FlashCompensation:    ptrFloat64(-4.5),
					Remarks:              ptrString("test frame #2"),
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

func mustParseTimestamp(t *testing.T, ts string, tz *time.Location, tf string) *time.Time {
	r := require.New(t)

	tts := strings.Split(ts, "T")
	r.Len(tts, 2)

	tt, err := parseTimestamp(tts[0], tts[1], tz, tf)
	r.NoError(err)

	return tt
}

func ptrInt64(i int64) *int64       { return &i }
func ptrFloat64(f float64) *float64 { return &f }
func ptrString(s string) *string    { return &s }
func ptrBool(b bool) *bool          { return &b }
