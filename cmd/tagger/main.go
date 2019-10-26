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
				"filmID":   film.ID,
				"cameraID": film.CameraID,
				"frameNo":  f.Number,
			})

			et := tagger.NewExifTool(exiftoolBinary, filename)

			if f.AFMode != "" {
				et.FocusMode(f.AFMode)
			}

			if f.Av > 0 {
				et.Aperture(f.Av)
			}

			if f.ExposureCompensation > 0 {
				et.ExposureCompensation(f.ExposureCompensation)
			}

			if f.FocalLength > 0 {
				et.FocalLength(f.FocalLength)
			}

			if f.ISO > 0 {
				et.ISO(f.ISO)
			}

			if f.MeteringMode != "" {
				et.MeteringMode(f.MeteringMode)
			}

			if f.ShootingMode != "" {
				et.ShootingMode(f.ShootingMode)
			}

			if !f.Timestamp.IsZero() {
				et.Timestamp(f.Timestamp)
			}

			if f.Tv != "" {
				et.Exposure(f.Tv)
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
