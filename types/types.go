package types

import "time"

// Frame model to store all the data about particular frame
type Frame struct {
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
