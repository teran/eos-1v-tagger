package tagger

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestExiftoolOptions(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name       string
		fname      string
		f          func(*ExifTool)
		expCommand string
	}

	tcs := []testCase{
		{
			name:       "no options, filename only specified",
			fname:      "test-file",
			f:          func(e *ExifTool) {},
			expCommand: `"test-file"`,
		},
		{
			name:  "iso specified",
			fname: "test-file-with-iso",
			f: func(e *ExifTool) {
				e.ISO(200)
			},
			expCommand: `-ISO="200" -ISOSpeed="200" "test-file-with-iso"`,
		},
		{
			name:  "aperture specified",
			fname: "test-file-with-aperture",
			f: func(e *ExifTool) {
				e.Aperture(1.4)
			},
			expCommand: `-FNumber="1.4" -ApertureValue="1.4" "test-file-with-aperture"`,
		},
		{
			name:  "focal length specified",
			fname: "test-file-with-focal-length",
			f: func(e *ExifTool) {
				e.FocalLength(35)
			},
			expCommand: `-FocalLength="35mm" "test-file-with-focal-length"`,
		},
		{
			name:  "exposure time specified",
			fname: "test-file-with-exposure",
			f: func(e *ExifTool) {
				e.Exposure("1/2566")
			},
			expCommand: `-ExposureTime="1/2566" -ShutterSpeedValue="1/2566" "test-file-with-exposure"`,
		},
		{
			name:  "exposure compensation specified",
			fname: "test-file-with-exposure-compensation",
			f: func(e *ExifTool) {
				e.ExposureCompensation(0.5)
			},
			expCommand: `-ExposureCompensation="0.5" "test-file-with-exposure-compensation"`,
		},
		{
			name:  "metering mode specified",
			fname: "test-file-with-metering-mode",
			f: func(e *ExifTool) {
				e.MeteringMode("Evaluative")
			},
			expCommand: `-MeteringMode="Evaluative" "test-file-with-metering-mode"`,
		},
		{
			name:  "shooting mode specified",
			fname: "test-file-with-shooting-mode",
			f: func(e *ExifTool) {
				e.ShootingMode("Aperture-priority AE")
			},
			expCommand: `-ShootingMode="Aperture-priority AE" "test-file-with-shooting-mode"`,
		},
		{
			name:  "focus mode specified",
			fname: "test-file-with-focus-mode",
			f: func(e *ExifTool) {
				e.FocusMode("One-Shot AF")
			},
			expCommand: `-FocusMode="One-Shot AF" "test-file-with-focus-mode"`,
		},
		{
			name:  "timestamp specified",
			fname: "test-file-with-timestamp",
			f: func(e *ExifTool) {
				ts, err := time.Parse(time.RFC3339, "2019-08-21T14:06:13Z")
				r.NoError(err)

				e.Timestamp(ts)
			},
			expCommand: `-DateTimeOriginal="2019-08-21T14:06:13Z" -ModifyDate="2019-08-21T14:06:13Z" "test-file-with-timestamp"`,
		},
		{
			name:  "geotag specified",
			fname: "test-file-with-geotag",
			f: func(e *ExifTool) {
				e.GeoTag("blah.log")
			},
			expCommand: `-GeoTag="blah.log" "test-file-with-geotag"`,
		},
		{
			name:  "DateTimeDigitized copied from CreateDate",
			fname: "test-file-with-date-time-digitized",
			f: func(e *ExifTool) {
				e.SetDateTimeDigitizedFromCreateDate()
			},
			expCommand: `-DateTimeDigitized<"CreateDate" "test-file-with-date-time-digitized"`,
		},
	}

	for _, tc := range tcs {
		e := NewExifTool(tc.fname)
		tc.f(e)

		r.Equalf(
			fmt.Sprintf("%s %s", exifToolDefaultCmd, tc.expCommand),
			e.Cmd(),
			tc.name)
	}
}
