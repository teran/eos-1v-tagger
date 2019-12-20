package tagger

import "time"

const (
	// TimestampFormatUS ...
	TimestampFormatUS = "1/2/2006T15:04:05"

	// TimestampFormatEU ...
	TimestampFormatEU = "2/1/2006T15:04:05"
)

type (
	// Frame model to store all the data about particular frame
	Frame struct {
		Flag                 *bool
		Number               *int64
		FocalLength          *int64
		MaxAperture          *float64
		Tv                   *string
		Av                   *float64
		ISO                  *int64
		ExposureCompensation *float64
		FlashCompensation    *float64
		FlashMode            *string
		MeteringMode         *string
		ShootingMode         *string
		FilmAdvanceMode      *string
		AFMode               *string
		BulbExposureTime     *string
		Timestamp            *time.Time
		MultipleExposure     *string
		BatteryLoadedDate    *time.Time
		Remarks              *string
	}
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
