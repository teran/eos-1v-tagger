package tagger

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	types "github.com/teran/eos-1v-tagger/types"
)

// ExifToolOption type
type ExifToolOption struct {
	key      string
	value    string
	operator string
}

// ExifTool type
type ExifTool struct {
	binary   string
	filename string
	options  []ExifToolOption
}

var (
	exifToolDefaultOpts = []string{"-overwrite_original"}
)

// New creates new ExifTool object
func New(binary, filename string) *ExifTool {
	return &ExifTool{
		binary:   binary,
		filename: filename,
	}
}

// NewFromFrame creates exiftool command right from frame object
func NewFromFrame(binary, filename string, f *types.Frame) *ExifTool {
	et := New(binary, filename)

	if f.AFMode != nil {
		et.FocusMode(f.AFMode.String())
	}

	if f.Av != nil {
		et.Aperture(*f.Av)
	}

	if f.ExposureCompensation != nil {
		et.ExposureCompensation(*f.ExposureCompensation)
	}

	if f.FocalLength != nil {
		et.FocalLength(*f.FocalLength)
	}

	if f.ISO != nil {
		et.ISO(*f.ISO)
	}

	if f.MeteringMode != nil {
		et.MeteringMode(f.MeteringMode.String())
	}

	if f.ShootingMode != nil {
		et.ShootingMode(f.ShootingMode.String())
	}

	if !f.Timestamp.IsZero() {
		et.Timestamp(*f.Timestamp)
	}

	if f.Tv != nil {
		et.Exposure(*f.Tv)
	}

	return et
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

// FileSource sets File Source field to exiftool command
func (e *ExifTool) FileSource(fsource string) *ExifTool {
	e.add("FileSource", fsource)
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

// Make sets Make parameters to exiftool command
func (e *ExifTool) Make(m string) *ExifTool {
	e.add("Make", m)

	return e
}

// Model sets Model parameters to exiftool command
func (e *ExifTool) Model(m string) *ExifTool {
	e.add("Model", m)

	return e
}

// MeteringMode sets metering mode parameters to exiftool command
func (e *ExifTool) MeteringMode(m string) *ExifTool {
	e.add("MeteringMode", m)

	return e
}

// SerialNumber sets SerialNumber parameters to exiftool command
func (e *ExifTool) SerialNumber(sn string) *ExifTool {
	e.add("SerialNumber", sn)

	return e
}

// SetDateTimeDigitizedFromCreateDate sets SetDateTimeDigitized from CreateDate field
func (e *ExifTool) SetDateTimeDigitizedFromCreateDate() *ExifTool {
	e.copy("CreateDate", "DateTimeDigitized")

	return e
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
	cmd := e.binary
	cmd += " " + strings.Join(exifToolDefaultOpts, " ")

	for _, o := range e.options {
		cmd += " "
		cmd += strconv.Quote("-" + o.key + o.operator + o.value)
	}

	cmd += fmt.Sprintf(` "%s"`, e.filename)

	return cmd
}
