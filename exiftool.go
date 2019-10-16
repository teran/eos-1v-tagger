package tagger

import (
	"fmt"
	"strconv"
	"time"
)

// NewExifTool creates new ExifTool object
func NewExifTool(filename string) *ExifTool {
	return &ExifTool{
		filename: filename,
	}
}

// Aperture sets Aperture parameters to exiftool command
func (e *ExifTool) Aperture(v float64) *ExifTool {
	vs := strconv.FormatFloat(v, 'f', -1, 64)

	e.add("FNumber", vs)
	e.add("ApertureValue", vs)

	return e
}

// Exposure sets Exposure value to exiftool command
func (e *ExifTool) Exposure(t string) *ExifTool {
	e.add("ExposureTime", t)
	e.add("ShutterSpeedValue", t)

	return e
}

// ExposureCompensation sets ExposureCompensation value to exiftool command
func (e *ExifTool) ExposureCompensation(ec float64) *ExifTool {
	e.add("ExposureCompensation", strconv.FormatFloat(ec, 'f', -1, 64))

	return e
}

// FocalLength sets Focal length to exiftool command
func (e *ExifTool) FocalLength(fl int64) *ExifTool {
	e.add(
		"FocalLength",
		fmt.Sprintf("%dmm", fl),
	)

	return e
}

// FocusMode sets Focus mode parameters to exiftool command
func (e *ExifTool) FocusMode(m string) *ExifTool {
	e.add("FocusMode", m)

	return e
}

// GeoTag adds `-geotag` to exiftool command
func (e *ExifTool) GeoTag(filename string) *ExifTool {
	e.add("GeoTag", filename)

	return e
}

// ISO sets ISO parameters to exiftool command
func (e *ExifTool) ISO(v int64) *ExifTool {
	vs := strconv.FormatInt(v, 10)
	e.add("ISO", vs)
	e.add("ISOSpeed", vs)

	return e
}

// MeteringMode sets metering mode parameters to exiftool command
func (e *ExifTool) MeteringMode(m string) *ExifTool {
	e.add("MeteringMode", m)

	return e
}

// SetDateTimeDigitizedFromCreateDate sets SetDateTimeDigitized from CreateDate field
func (e *ExifTool) SetDateTimeDigitizedFromCreateDate() {
	e.copy("CreateDate", "DateTimeDigitized")
}

// ShootingMode sets shooting mode parameters to exiftool command
func (e *ExifTool) ShootingMode(m string) *ExifTool {
	e.add("ShootingMode", m)

	return e
}

// Timestamp sets the timestamp shot made on
func (e *ExifTool) Timestamp(t time.Time) *ExifTool {
	ts := t.Format(time.RFC3339)
	e.add("DateTimeOriginal", ts)
	e.add("ModifyDate", ts)

	return e
}

func (e *ExifTool) add(k, v string) {
	e.options = append(e.options, ExifToolOption{
		key:      k,
		value:    v,
		operator: "=",
	})
}

func (e *ExifTool) copy(from, to string) {
	e.options = append(e.options, ExifToolOption{
		key:      to,
		value:    from,
		operator: "<",
	})
}

// Cmd returns complete exiftool command
func (e *ExifTool) Cmd() string {
	cmd := exifToolDefaultCmd
	for _, o := range e.options {
		cmd += " "
		cmd += "-" + o.key + o.operator + strconv.Quote(o.value)
	}

	cmd += fmt.Sprintf(` "%s"`, e.filename)

	return cmd
}
