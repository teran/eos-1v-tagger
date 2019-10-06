package tagger

import "time"

const (
	// TimestampFormat ...
	TimestampFormat = "1/2/2006T15:04:05"

	frameHeader = `,Frame No.,Focal length,Max. aperture,Tv,Av,ISO (M)` +
		`,Exposure compensation,Flash exposure compensation,Flash mode,` +
		`Metering mode,Shooting mode,Film advance mode,AF mode,Bulb exposure ` +
		`time,Date,Time,Multiple exposure,Battery-loaded date,Battery-loaded time,Remarks`
)

type (
	// Film model to store all the data about the film itself
	Film struct {
		ID                  string
		Title               string
		FilmLoadedTimestamp time.Time
		FrameCount          int64
		ISO                 int64
		Remarks             string
		Frames              []Frame
	}

	// Frame model to store all the data about particular frame
	Frame struct {
		Flag                 bool
		Number               int64
		FocalLength          int64
		MaxAperture          float64
		Tv                   string
		Av                   float64
		ISO                  int64
		ExposureCompensation interface{}
		FlashCompensation    interface{}
		FlashMode            interface{}
		MeteringMode         interface{}
		ShoothingMode        interface{}
		FilmAdvanceMode      interface{}
		AFMode               interface{}
		BulbExposureTime     interface{}
		Timestamp            time.Time
		MultipleExposure     interface{}
		BatteryLoadedDate    time.Time
		Remarks              string
	}
)
