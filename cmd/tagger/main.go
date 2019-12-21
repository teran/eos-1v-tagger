package main

import (
	"flag"
	"fmt"
	"log"

	tagger "github.com/teran/eos-1v-tagger"
	exiftool "github.com/teran/eos-1v-tagger/exiftool"
	format "github.com/teran/eos-1v-tagger/format"
	parser "github.com/teran/eos-1v-tagger/parser"
	types "github.com/teran/eos-1v-tagger/types"
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

	t, err := parser.New(flag.Arg(0), tz, tf)
	if err != nil {
		log.Fatalf("error initializing CSV parser: %s", err)
	}

	films, err := t.Parse()
	if err != nil {
		log.Fatalf("error parsing CSV: %s", err)
	}

	for _, film := range films {
		for _, f := range film.Frames {
			filename := format.Format(filenamePattern, map[string]interface{}{
				"filmID":   *film.ID,
				"cameraID": *film.CameraID,
				"frameNo":  *f.Number,
			})

			et := exiftool.NewFromFrame(exiftoolBinary, filename, f)

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
				gt := exiftool.New(exiftoolBinary, filename)
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
		return types.TimestampFormatUS, nil
	case "EU":
		return types.TimestampFormatEU, nil
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
