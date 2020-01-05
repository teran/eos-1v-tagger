package tagger

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	types "github.com/teran/eos-1v-tagger/types"
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
			expCommand: `"-ISO=200" "-ISOSpeed=200" "test-file-with-iso"`,
		},
		{
			name:  "aperture specified",
			fname: "test-file-with-aperture",
			f: func(e *ExifTool) {
				e.Aperture(1.4)
			},
			expCommand: `"-FNumber=1.4" "-ApertureValue=1.4" "test-file-with-aperture"`,
		},
		{
			name:  "focal length specified",
			fname: "test-file-with-focal-length",
			f: func(e *ExifTool) {
				e.FocalLength(35)
			},
			expCommand: `"-FocalLength=35mm" "test-file-with-focal-length"`,
		},
		{
			name:  "exposure time specified",
			fname: "test-file-with-exposure",
			f: func(e *ExifTool) {
				e.Exposure("1/2566")
			},
			expCommand: `"-ExposureTime=1/2566" "-ShutterSpeedValue=1/2566" "test-file-with-exposure"`,
		},
		{
			name:  "exposure compensation specified",
			fname: "test-file-with-exposure-compensation",
			f: func(e *ExifTool) {
				e.ExposureCompensation(0.5)
			},
			expCommand: `"-ExposureCompensation=0.5" "test-file-with-exposure-compensation"`,
		},
		{
			name:  "timestamp specified",
			fname: "test-file-with-timestamp",
			f: func(e *ExifTool) {
				ts, err := time.Parse(time.RFC3339, "2019-08-21T14:06:13Z")
				r.NoError(err)

				e.Timestamp(ts)
			},
			expCommand: `"-DateTimeOriginal=2019-08-21T14:06:13Z" "-ModifyDate=2019-08-21T14:06:13Z" "test-file-with-timestamp"`,
		},
		{
			name:  "geotag specified",
			fname: "test-file-with-geotag",
			f: func(e *ExifTool) {
				e.GeoTag("blah.log")
			},
			expCommand: `"-GeoTag=blah.log" "test-file-with-geotag"`,
		},
		{
			name:  "geotime specified",
			fname: "test-file-with-geotime",
			f: func(e *ExifTool) {
				ts := time.Date(2001, 3, 14, 15, 32, 53, 0, time.UTC)
				e.GeoTime(ts)
			},
			expCommand: `"-GeoTime=2001:03:14 15:32:53Z" "test-file-with-geotime"`,
		},
		{
			name:  "DateTimeDigitized copied from CreateDate",
			fname: "test-file-with-date-time-digitized",
			f: func(e *ExifTool) {
				e.SetDateTimeDigitizedFromCreateDate()
			},
			expCommand: `"-DateTimeDigitized<CreateDate" "test-file-with-date-time-digitized"`,
		},
		{
			name:  "make specified",
			fname: "test-file-with-make",
			f: func(e *ExifTool) {
				e.Make("TestDevice/1.0")
			},
			expCommand: `"-Make=TestDevice/1.0" "test-file-with-make"`,
		},
		{
			name:  "model specified",
			fname: "test-file-with-model",
			f: func(e *ExifTool) {
				e.Model("TestModel/1.0")
			},
			expCommand: `"-Model=TestModel/1.0" "test-file-with-model"`,
		},
		{
			name:  "serial number specified",
			fname: "test-file-with-serial-number",
			f: func(e *ExifTool) {
				e.SerialNumber("1234567890!")
			},
			expCommand: `"-SerialNumber=1234567890!" "test-file-with-serial-number"`,
		},
		{
			name:  "file source specified",
			fname: "test-file-with-file-source",
			f: func(e *ExifTool) {
				e.FileSource("Film Scanner")
			},
			expCommand: `"-FileSource=Film Scanner" "test-file-with-file-source"`,
		},
	}

	for _, tc := range tcs {
		e := New("exiftool", tc.fname)
		tc.f(e)

		r.Equalf(
			fmt.Sprintf("exiftool %s %s", strings.Join(exifToolDefaultOpts, " "), tc.expCommand),
			e.Cmd(),
			tc.name)
	}
}

func TestFromFrame(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name       string
		frame      types.Frame
		expOptions []ExifToolOption
	}

	tcs := []testCase{
		{
			name: "filled frame",
			frame: types.Frame{
				Flag:                 types.PtrBool(true),
				Number:               types.PtrInt64(23),
				FocalLength:          types.PtrInt64(123),
				MaxAperture:          types.PtrAperture(1.4),
				Tv:                   types.PtrString("1/300"),
				Av:                   types.PtrAperture(1.5),
				ISO:                  types.PtrInt64(400),
				ExposureCompensation: types.PtrFloat64(3.2),
				FlashCompensation:    types.PtrFloat64(-3.2),
				FlashMode:            types.PtrFlashMode(types.FlashModeOff),
				MeteringMode:         types.PtrMeteringMode(types.MeteringModeEvaluative),
				ShootingMode:         types.PtrShootingMode(types.ShootingModeProgramAE),
				FilmAdvanceMode:      types.PtrFilmAdvanceMode(types.FilmAdvanceModeSingleFrame),
				AFMode:               types.PtrAFMode(types.AFModeOneShotAF),
				BulbExposureTime:     types.PtrString("1/300"),
				Timestamp:            types.PtrTime(time.Date(2009, 2, 12, 15, 34, 23, 0, time.UTC)),
				MultipleExposure:     types.PtrMultipleExposure(types.MultipleExposureOff),
				BatteryLoadedDate:    types.PtrTime(time.Date(2008, 1, 5, 11, 58, 14, 0, time.UTC)),
				Remarks:              types.PtrString("blah!"),
			},
			expOptions: []ExifToolOption{
				{
					key:      "Canon:FocusMode",
					value:    "One-Shot AF",
					operator: "=",
				},
				{
					key:      "ExifIFD:FNumber",
					value:    "1.5",
					operator: "=",
				},
				{
					key:      "ExifIFD:ApertureValue",
					value:    "1.5",
					operator: "=",
				},
				{
					key:      "Canon:FNumber",
					value:    "1.5",
					operator: "=",
				},
				{
					key:      "Canon:TargetAperture",
					value:    "1.5",
					operator: "=",
				},
				{
					key:      "ExposureCompensation",
					value:    "3.2",
					operator: "=",
				},
				{
					key:      "FocalLength",
					value:    "123mm",
					operator: "=",
				},
				{
					key:      "ISO",
					value:    "400",
					operator: "=",
				},
				{
					key:      "ISOSpeed",
					value:    "400",
					operator: "=",
				},
				{
					key:      "ExifIFD:Flash",
					value:    "Off, Did not fire",
					operator: "=",
				},
				{
					key:      "Canon:FlashBits",
					value:    "(none)",
					operator: "=",
				},
				{
					key:      "Canon:CanonFlashMode",
					value:    "Off",
					operator: "=",
				},
				{
					key:      "CanonCustom:PF2DisableMeteringModes",
					value:    "Off",
					operator: "=",
				},
				{
					key:      "ExifIFD:MeteringMode",
					value:    "Multi-Segment",
					operator: "=",
				},
				{
					key:      "Canon:MeteringMode",
					value:    "Evaluative",
					operator: "=",
				},
				{
					key:      "ExifIFD:ExposureProgram",
					value:    "Program AE",
					operator: "=",
				},
				{
					key:      "Canon:CanonExposureMode",
					value:    "Program AE",
					operator: "=",
				},
				{
					key:      "CanonCustom:PF1DisableShootingModes",
					value:    "Off",
					operator: "=",
				},
				{
					key:      "CanonCustom:PF6PresetShootingModes",
					value:    "Off",
					operator: "=",
				},
				{
					key:      "DateTimeOriginal",
					value:    "2009-02-12T15:34:23Z",
					operator: "=",
				},
				{
					key:      "ModifyDate",
					value:    "2009-02-12T15:34:23Z",
					operator: "=",
				},
				{
					key:      "ExposureTime",
					value:    "1/300",
					operator: "=",
				},
				{
					key:      "ShutterSpeedValue",
					value:    "1/300",
					operator: "=",
				},
			},
		},
		{
			name:       "empty frame",
			frame:      types.Frame{},
			expOptions: []ExifToolOption{},
		},
	}

	for _, tc := range tcs {
		e := NewFromFrame("bin", "fname", &tc.frame)
		r.ElementsMatchf(tc.expOptions, e.Options(), tc.name)
	}
}
