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
			ID:                  types.PtrInt64(139),
			CameraID:            types.PtrInt64(1),
			Title:               types.PtrString("SampleTest film #139"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "09/28/2019T10:21:32", tz, types.TimestampFormatUS),
			FrameCount:          types.PtrInt64(2),
			ISO:                 types.PtrInt64(400),
			Remarks:             types.PtrString("test remarks data"),
			Frames: []*types.Frame{
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(1),
					FocalLength:          types.PtrInt64(24),
					MaxAperture:          types.PtrAperture(1.4),
					Tv:                   types.PtrString("1/40"),
					Av:                   types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(400),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					ExposureCompensation: types.PtrFloat64(0),
					FlashCompensation:    types.PtrFloat64(0),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:18", tz, types.TimestampFormatUS),
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					BatteryLoadedDate:    nil,
					Remarks:              types.PtrString("test frame #1"),
				},
				{
					Flag:                 types.PtrBool(true),
					Number:               types.PtrInt64(2),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					Tv:                   types.PtrString("1/60"),
					Av:                   types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(400),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/7/2019T20:02:29", tz, types.TimestampFormatUS),
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					BatteryLoadedDate:    nil,
					ExposureCompensation: types.PtrFloat64(-5),
					FlashCompensation:    types.PtrFloat64(-4.5),
					Remarks:              types.PtrString("test frame #2"),
				},
			},
		},
		{
			ID:                  types.PtrInt64(140),
			CameraID:            types.PtrInt64(1),
			Title:               types.PtrString("SampleTest film #139 part II"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "10/07/2019T22:55:58", tz, types.TimestampFormatUS),
			FrameCount:          types.PtrInt64(2),
			ISO:                 types.PtrInt64(400),
			Remarks:             types.PtrString("test remarks data 2"),
			Frames: []*types.Frame{
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(1),
					FocalLength:          types.PtrInt64(14),
					MaxAperture:          types.PtrAperture(1.4),
					Tv:                   types.PtrString("1/1600"),
					Av:                   types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(200),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeProgramAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:38", tz, types.TimestampFormatUS),
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					BatteryLoadedDate:    nil,
					ExposureCompensation: types.PtrFloat64(1),
					FlashCompensation:    types.PtrFloat64(2),
					Remarks:              types.PtrString("test frame remarks #1"),
				},
				{
					Flag:                 types.PtrBool(true),
					Number:               types.PtrInt64(2),
					FocalLength:          types.PtrInt64(16),
					MaxAperture:          types.PtrAperture(1.4),
					Tv:                   types.PtrString("1/1250"),
					Av:                   types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(800),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "10/13/2019T14:55:55", tz, types.TimestampFormatUS),
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					BatteryLoadedDate:    nil,
					ExposureCompensation: types.PtrFloat64(-1),
					FlashCompensation:    types.PtrFloat64(-2),
					Remarks:              types.PtrString("test frame remarks #2"),
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
			ID:                  types.PtrInt64(139),
			CameraID:            types.PtrInt64(1),
			Title:               types.PtrString("SampleTest film #139"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "28/09/2019T10:21:32", tz, types.TimestampFormatEU),
			FrameCount:          types.PtrInt64(2),
			ISO:                 types.PtrInt64(400),
			Remarks:             types.PtrString("test remarks data"),
			Frames: []*types.Frame{
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(1),
					FocalLength:          types.PtrInt64(24),
					MaxAperture:          types.PtrAperture(1.4),
					Tv:                   types.PtrString("1/40"),
					Av:                   types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(400),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ExposureCompensation: types.PtrFloat64(0),
					FlashCompensation:    types.PtrFloat64(0),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "15/07/2019T20:02:18", tz, types.TimestampFormatEU),
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					BatteryLoadedDate:    nil,
					Remarks:              types.PtrString("test frame #1"),
				},
				{
					Flag:                 types.PtrBool(true),
					Number:               types.PtrInt64(2),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					Tv:                   types.PtrString("1/60"),
					Av:                   types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(400),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            nil,
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					BatteryLoadedDate:    nil,
					ExposureCompensation: types.PtrFloat64(-5),
					FlashCompensation:    types.PtrFloat64(-4.5),
					Remarks:              types.PtrString("test frame #2"),
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
			ID:                  types.PtrInt64(119),
			CameraID:            types.PtrInt64(1),
			Title:               types.PtrString("sample test film#119"),
			FilmLoadedTimestamp: mustParseTimestamp(t, "09/28/2019T10:21:32", tz, types.TimestampFormatUS),
			FrameCount:          types.PtrInt64(2),
			ISO:                 types.PtrInt64(400),
			Remarks:             types.PtrString("test remarks data"),
			Frames: []*types.Frame{
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(1),
					ISO:                  types.PtrInt64(400),
					ExposureCompensation: types.PtrFloat64(-6.4),
				},
				{
					Flag:        types.PtrBool(false),
					Number:      types.PtrInt64(2),
					ISO:         types.PtrInt64(400),
					FocalLength: types.PtrInt64(35),
				},
				{
					Flag:        types.PtrBool(false),
					Number:      types.PtrInt64(3),
					FocalLength: types.PtrInt64(35),
					MaxAperture: types.PtrAperture(1.4),
					ISO:         types.PtrInt64(400),
				},
				{
					Flag:        types.PtrBool(false),
					Number:      types.PtrInt64(4),
					FocalLength: types.PtrInt64(35),
					MaxAperture: types.PtrAperture(1.4),
					ISO:         types.PtrInt64(400),
					Tv:          types.PtrString("1/1000"),
				},
				{
					Flag:        types.PtrBool(false),
					Number:      types.PtrInt64(5),
					FocalLength: types.PtrInt64(35),
					MaxAperture: types.PtrAperture(1.4),
					ISO:         types.PtrInt64(400),
					Tv:          types.PtrString("1/1000"),
					Av:          types.PtrAperture(1.4),
				},
				{
					Flag:        types.PtrBool(false),
					Number:      types.PtrInt64(6),
					FocalLength: types.PtrInt64(35),
					MaxAperture: types.PtrAperture(1.4),
					ISO:         types.PtrInt64(640),
					Tv:          types.PtrString("1/1000"),
					Av:          types.PtrAperture(1.4),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(7),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(200),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(8),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(100),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(9),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(50),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(10),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(500),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(11),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(640),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
				},
				{
					Flag:                 types.PtrBool(true),
					Number:               types.PtrInt64(12),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(1280),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(13),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(3200),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
				},
				{
					Flag:                 types.PtrBool(false),
					Number:               types.PtrInt64(14),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(1600),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					Timestamp:            mustParseTimestamp(t, "07/10/2019T20:02:18", tz, types.TimestampFormatEU),
				},
				{
					Flag:                 types.PtrBool(true),
					Number:               types.PtrInt64(15),
					FocalLength:          types.PtrInt64(35),
					MaxAperture:          types.PtrAperture(1.4),
					ISO:                  types.PtrInt64(64),
					Tv:                   types.PtrString("1/1000"),
					Av:                   types.PtrAperture(1.4),
					ExposureCompensation: types.PtrFloat64(1.3),
					FlashCompensation:    types.PtrFloat64(-3.2),
					FlashMode:            types.PtrFlashMode(types.FlashModeOff),
					MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
					ShootingMode:         types.PtrShootingMode(types.ShootingModeAperturePriorityAE),
					FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
					AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
					MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
					Timestamp:            mustParseTimestamp(t, "07/10/2019T20:02:18", tz, types.TimestampFormatEU),
					Remarks:              types.PtrString("test frame #1"),
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
