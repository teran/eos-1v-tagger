package tagger

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	tagger "github.com/teran/eos-1v-tagger"
	types "github.com/teran/eos-1v-tagger/types"
)

func TestTwoFilmsInSingleCSV(t *testing.T) {
	r := require.New(t)

	tz, err := tagger.LocationByTimeZone("CET")
	r.NoError(err)

	p, err := New("testdata/two-films.csv", tz, types.TimestampFormatUS)
	r.NoError(err)
	r.NotNil(p)

	defer func() {
		err := p.Close()
		r.NoError(err)
	}()

	film, err := p.Parse()
	r.NoError(err)
	r.Equal([]*types.Film{
		{
			ID:                  ptrInt64(139),
			CameraID:            ptrInt64(1),
			Title:               ptrString("SampleTest film #139"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "09/28/2019T10:21:32", tz, types.TimestampFormatUS),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data"),
			Frames: []*types.Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					FocalLength:          ptrInt64(24),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/40"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(400),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					ExposureCompensation: ptrFloat64(0),
					FlashCompensation:    ptrFloat64(0),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:18", tz, types.TimestampFormatUS),
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
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:29", tz, types.TimestampFormatUS),
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
			FilmLoadedTimestamp: mustParseTimestamp(t, "10/07/2019T22:55:58", tz, types.TimestampFormatUS),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data 2"),
			Frames: []*types.Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					FocalLength:          ptrInt64(14),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/1600"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(200),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeProgramAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:38", tz, types.TimestampFormatUS),
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
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:55", tz, types.TimestampFormatUS),
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

	tz, err := tagger.LocationByTimeZone("CET")
	r.NoError(err)

	p, err := New("testdata/film-with-partial-timestamps-eu.csv", tz, types.TimestampFormatEU)
	r.NoError(err)
	r.NotNil(p)

	defer func() {
		err := p.Close()
		r.NoError(err)
	}()

	film, err := p.Parse()
	r.NoError(err)
	r.Equal([]*types.Film{
		{
			ID:                  ptrInt64(139),
			CameraID:            ptrInt64(1),
			Title:               ptrString("SampleTest film #139"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "28/09/2019T10:21:32", tz, types.TimestampFormatEU),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data"),
			Frames: []*types.Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					FocalLength:          ptrInt64(24),
					MaxAperture:          ptrFloat64(1.4),
					Tv:                   ptrString("1/40"),
					Av:                   ptrFloat64(1.4),
					ISO:                  ptrInt64(400),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ExposureCompensation: ptrFloat64(0),
					FlashCompensation:    ptrFloat64(0),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "15/07/2019T20:02:18", tz, types.TimestampFormatEU),
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
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
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

func TestPartialData(t *testing.T) {
	r := require.New(t)

	tz, err := tagger.LocationByTimeZone("CET")
	r.NoError(err)

	p, err := New("testdata/partial-data.csv", tz, types.TimestampFormatUS)
	r.NoError(err)
	r.NotNil(p)

	film, err := p.Parse()
	r.NoError(err)
	r.Equal([]*types.Film{
		{
			ID:                  ptrInt64(119),
			CameraID:            ptrInt64(1),
			Title:               ptrString("sample test film#119"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "09/28/2019T10:21:32", tz, types.TimestampFormatUS),
			FrameCount:          ptrInt64(2),
			ISO:                 ptrInt64(400),
			Remarks:             ptrString("test remarks data"),
			Frames: []*types.Frame{
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(1),
					ISO:                  ptrInt64(400),
					ExposureCompensation: ptrFloat64(-6.4),
				},
				{
					Flag:        ptrBool(false),
					Number:      ptrInt64(2),
					ISO:         ptrInt64(400),
					FocalLength: ptrInt64(35),
				},
				{
					Flag:        ptrBool(false),
					Number:      ptrInt64(3),
					FocalLength: ptrInt64(35),
					MaxAperture: ptrFloat64(1.4),
					ISO:         ptrInt64(400),
				},
				{
					Flag:        ptrBool(false),
					Number:      ptrInt64(4),
					FocalLength: ptrInt64(35),
					MaxAperture: ptrFloat64(1.4),
					ISO:         ptrInt64(400),
					Tv:          ptrString("1/1000"),
				},
				{
					Flag:        ptrBool(false),
					Number:      ptrInt64(5),
					FocalLength: ptrInt64(35),
					MaxAperture: ptrFloat64(1.4),
					ISO:         ptrInt64(400),
					Tv:          ptrString("1/1000"),
					Av:          ptrFloat64(1.4),
				},
				{
					Flag:        ptrBool(false),
					Number:      ptrInt64(6),
					FocalLength: ptrInt64(35),
					MaxAperture: ptrFloat64(1.4),
					ISO:         ptrInt64(640),
					Tv:          ptrString("1/1000"),
					Av:          ptrFloat64(1.4),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(7),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(200),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(8),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(100),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(9),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(50),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(10),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(500),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(11),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(640),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
				},
				{
					Flag:                 ptrBool(true),
					Number:               ptrInt64(12),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(1280),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(13),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(3200),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
				},
				{
					Flag:                 ptrBool(false),
					Number:               ptrInt64(14),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(1600),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "07/10/2019T20:02:18", tz, types.TimestampFormatEU),
				},
				{
					Flag:                 ptrBool(true),
					Number:               ptrInt64(15),
					FocalLength:          ptrInt64(35),
					MaxAperture:          ptrFloat64(1.4),
					ISO:                  ptrInt64(64),
					Tv:                   ptrString("1/1000"),
					Av:                   ptrFloat64(1.4),
					ExposureCompensation: ptrFloat64(1.3),
					FlashCompensation:    ptrFloat64(-3.2),
					FlashMode:            ptrString("OFF"),
					MeteringMode:         ptrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         ptrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      ptrString("Single-frame"),
					AFMode:               ptrAFMode(types.AFModeOneShotAF),
					MultipleExposure:     ptrString("OFF"),
					Timestamp:            mustParseTimestamp(t, "07/10/2019T20:02:18", tz, types.TimestampFormatEU),
					Remarks:              ptrString("test frame #1"),
				},
			},
		},
	}, film)
}

func TestEmptyFrame(t *testing.T) {
	r := require.New(t)

	tz, err := tagger.LocationByTimeZone("CET")
	r.NoError(err)

	p, err := New("testdata/empty-frame.csv", tz, types.TimestampFormatUS)
	r.NoError(err)
	r.NotNil(p)

	_, err = p.Parse()
	r.Error(err)
	r.Equal(ErrEmptyFrame, err)
}

func mustParseTimestamp(t *testing.T, ts string, tz *time.Location, tf string) *time.Time {
	r := require.New(t)

	tts := strings.Split(ts, "T")
	r.Len(tts, 2)

	tt, err := parseTimestamp(tts[0], tts[1], tz, tf)
	r.NoError(err)

	return tt
}

func ptrInt64(i int64) *int64                                   { return &i }
func ptrFloat64(f float64) *float64                             { return &f }
func ptrString(s string) *string                                { return &s }
func ptrBool(b bool) *bool                                      { return &b }
func ptrMeteringMode(mm types.MeteringMode) *types.MeteringMode { return &mm }
func ptrShootingMode(sm types.ShootingMode) *types.ShootingMode { return &sm }
func ptrAFMode(am types.AFMode) *types.AFMode                   { return &am }
