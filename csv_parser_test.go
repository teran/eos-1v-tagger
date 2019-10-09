package tagger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCSVParser(t *testing.T) {
	r := require.New(t)

	p, err := NewCSVParser("testdata/sample.CSV")
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
		FilmLoadedTimestamp: mustParseTimestamp("09/01/2010T14:00:00"),
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
				ShoothingMode:     "Program AE",
				FilmAdvanceMode:   "Single-frame",
				AFMode:            "One-Shot AF",
				Timestamp:         time.Date(2010, 11, 9, 18, 31, 26, 0, time.UTC),
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
				ShoothingMode:        "Program AE",
				FilmAdvanceMode:      "Single-frame",
				AFMode:               "One-Shot AF",
				Timestamp:            time.Date(2010, 12, 9, 18, 32, 55, 0, time.UTC),
				MultipleExposure:     "OFF",
				BatteryLoadedDate:    time.Time{},
				ExposureCompensation: -5,
				FlashCompensation:    -4.5,
			},
		},
	}, film)
}

func mustParseTimestamp(ts string) time.Time {
	tt, err := time.Parse(TimestampFormat, ts)
	if err != nil {
		panic(err)
	}

	return tt
}
