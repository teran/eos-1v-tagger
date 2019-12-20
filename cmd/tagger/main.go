package main

import (
	"flag"
	"fmt"
	"log"

	tagger "github.com/teran/eos-1v-tagger"
)

func main() {
	parseFlags()

	tz, err := tagger.LocationByTimeZone(timezone)
	if err != nil {
		log.Fatalf("error looking up timezone: %s", err)
	}

	tf, err := TimestampFormatFactory(timestampFormat)
	if err != nil {
		log.Fatalf("error loading timestamp format: %s", err)
	}

	t, err := tagger.NewCSVParser(flag.Arg(0), tz, tf)
	if err != nil {
		log.Fatalf("error initializing CSV parser: %s", err)
	}

	films, err := t.Parse()
	if err != nil {
		log.Fatalf("error parsing CSV: %s", err)
	}

	for _, film := range films {
		for _, f := range film.Frames {
			filename := tagger.Format(filenamePattern, map[string]interface{}{
				"filmID":   *film.ID,
				"cameraID": *film.CameraID,
				"frameNo":  *f.Number,
			})

			et := tagger.NewExifTool(exiftoolBinary, filename)

			if f.AFMode != nil {
				et.FocusMode(*f.AFMode)
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
				et.MeteringMode(*f.MeteringMode)
			}

			if f.ShootingMode != nil {
				et.ShootingMode(*f.ShootingMode)
			}

			if !f.Timestamp.IsZero() {
				et.Timestamp(*f.Timestamp)
			}

			if f.Tv != nil {
				et.Exposure(*f.Tv)
			}

			if setDigitized {
				et.SetDateTimeDigitizedFromCreateDate()
			}

			if make != "" {
				et.Make(make)
			}

			if model != "" {
				et.Model(model)
			}

			if serialNumber != "" {
				et.SerialNumber(serialNumber)
			}

			if fileSource != "" {
				if !validateFileSource(fileSource) {
					log.Fatalf("Bad `file-source` value. Available options: 'Film Scanner', 'Reflection Print Scanner', 'Digital Camera'")
				}
				et.FileSource(fileSource)
			}

			fmt.Println(et.Cmd())

			if geotag != "" {
				gt := tagger.NewExifTool(exiftoolBinary, filename)
				gt.GeoTag(geotag)
				fmt.Println(gt.Cmd())
			}
		}
	}
}

// TimestampFormatFactory provides timestamp format depending on settings
func TimestampFormatFactory(tf string) (string, error) {
	switch tf {
	case "US":
		return tagger.TimestampFormatUS, nil
	case "EU":
		return tagger.TimestampFormatEU, nil
	}
	return "", fmt.Errorf("unknown timestamp format: %s", tf)
}

func validateFileSource(fs string) bool {
	switch fs {
	case "Film Scanner":
		return true
	case "Reflection Print Scanner":
		return true
	case "Digital Camera":
		return true
	}
	return false
}
