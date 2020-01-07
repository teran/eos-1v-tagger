package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPtrBool(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		input     bool
		expOutput *bool
	}

	tcs := []testCase{
		{
			name:      "true to *true",
			input:     true,
			expOutput: func() *bool { t := true; return &t }(),
		},
		{
			name:      "false to *false",
			input:     false,
			expOutput: func() *bool { t := false; return &t }(),
		},
	}

	for _, tc := range tcs {
		r.Equalf(tc.expOutput, PtrBool(tc.input), tc.name)
	}
}

func TestPtrFloat64(t *testing.T) {
	r := require.New(t)

	r.Equal(func() *float64 { t := 3.2; return &t }(), PtrFloat64(3.2))
}

func TestPtrInt64(t *testing.T) {
	r := require.New(t)

	r.Equal(func() *int64 { t := int64(3); return &t }(), PtrInt64(3))
}

func TestPtrString(t *testing.T) {
	r := require.New(t)

	r.Equal(func() *string { t := "blah"; return &t }(), PtrString("blah"))
}

func TestPtrTime(t *testing.T) {
	r := require.New(t)
	ts := time.Now()

	r.Equal(func() *time.Time { return &ts }(), PtrTime(ts))
}

func TestPtrFilmAdvanceMode(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *FilmAdvanceMode { t := FilmAdvanceModeContinuousBodyOnly; return &t }(),
		PtrFilmAdvanceMode(FilmAdvanceModeContinuousBodyOnly))

	r.Nil(PtrFilmAdvanceMode(FilmAdvanceMode("")))
}

func TestPtrAFMode(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *AFMode { t := AFModeAIServoAF; return &t }(),
		PtrAFMode(AFModeAIServoAF))

	r.Nil(PtrAFMode(AFMode("")))
}

func TestPtrFlashMode(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *FlashMode { t := FlashModeETTL; return &t }(),
		PtrFlashMode(FlashModeETTL))

	r.Nil(PtrFlashMode(FlashMode("")))
}

func TestPtrMeteringMode(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *MeteringMode { t := MeteringModeCenterAveraging; return &t }(),
		PtrMeteringMode(MeteringModeCenterAveraging))

	r.Nil(PtrMeteringMode(MeteringMode("")))
}

func TestPtrMultipleExposure(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *MultipleExposure { t := MultipleExposureOn; return &t }(),
		PtrMultipleExposure(MultipleExposureOn))

	r.Nil(PtrMultipleExposure(MultipleExposure("")))
}

func TestPtrShootingMode(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *ShootingMode { t := ShootingModeDepthOfFieldAE; return &t }(),
		PtrShootingMode(ShootingModeDepthOfFieldAE))

	r.Nil(PtrShootingMode(ShootingMode("")))
}

func TestPtrFileSource(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *FileSource { t := FileSourceDigitalCamera; return &t }(),
		PtrFileSource(FileSourceDigitalCamera))

	r.Nil(PtrFileSource(FileSource("")))
}

func TestPtrTimestampFormat(t *testing.T) {
	r := require.New(t)

	r.Equal(
		func() *TimestampFormat { t := TimestampFormatEU; return &t }(),
		PtrTimestampFormat(TimestampFormatEU))

	r.Nil(PtrTimestampFormat(TimestampFormat("")))
}
