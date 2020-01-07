package main

import (
	"fmt"
	"log"
	"os"
	"time"

	config "github.com/teran/eos-1v-tagger/config"
	exiftool "github.com/teran/eos-1v-tagger/exiftool"
	format "github.com/teran/eos-1v-tagger/format"
	parser "github.com/teran/eos-1v-tagger/parser"
)

// LD vars
var (
	ldVersion   = "undefined"
	ldTimestamp = "0"
)

func main() {
	cfg := config.NewDefaultConfig()

	err := cfg.FillFromYaml("~/.tagger/config.yaml")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("error reading config file: %s", err)
	}

	if len(os.Args) < 1 {
		log.Fatalf("looks like programmer error: os.Args is less than 1 in length")
	}

	f := config.NewFlags(os.Args[0], ldVersion, ldTimestamp)
	err = cfg.FillFromFlags(f)
	if err != nil {
		log.Fatalf("error handling CLI flags: %s", err)
	}

	if cfg.GetDisplayHelp() {
		f.PrintUsageString()
		os.Exit(1)
	}

	if cfg.GetDisplayVersion() {
		f.PrintVersionString()
		os.Exit(1)
	}

	lookupTzFn := func(cID uint8) *time.Location {
		tzname := cfg.GetTimezoneByCameraID(cID)
		location, err := time.LoadLocation(tzname)
		if err != nil {
			log.Printf("ERROR: error looking up timezone: %s; Switching to default: UTC", err)
			return time.UTC
		}

		return location
	}

	if f.GetCSVPath() == "" {
		log.Fatal("CSV file is required to specify")
	}

	t, err := parser.New(f.GetCSVPath(), cfg.GetTimestampFormat().TimeLayout(), lookupTzFn)
	if err != nil {
		log.Fatalf("error initializing CSV parser: %s", err)
	}

	films, err := t.Parse()
	if err != nil {
		log.Fatalf("error parsing CSV: %s", err)
	}

	for _, film := range films {
		for _, f := range film.Frames {
			filename := format.Format(cfg.GetFilenamePattern(), map[string]interface{}{
				"filmID":   *film.ID,
				"cameraID": *film.CameraID,
				"frameNo":  *f.Number,
			})

			et := exiftool.NewFromFrame(cfg.GetExiftoolBinary(), filename, f)

			if cfg.GetSetDigitized() {
				et.SetDateTimeDigitizedFromCreateDate()
			}

			if cfg.GetMakeByCameraID(*film.CameraID) != nil {
				v := cfg.GetMakeByCameraID(*film.CameraID)
				et.Make(*v)
			}

			if cfg.GetModelByCameraID(*film.CameraID) != nil {
				v := cfg.GetModelByCameraID(*film.CameraID)
				et.Model(*v)
			}

			if cfg.GetSerialNumberByCameraID(*film.CameraID) != nil {
				v := cfg.GetSerialNumberByCameraID(*film.CameraID)
				et.SerialNumber(*v)
			}

			if cfg.GetFileSource() != nil {
				v := cfg.GetFileSource()
				et.FileSource(v.String())
			}

			if cfg.GetCopyright() != nil {
				v := cfg.GetCopyright()
				et.Copyright(*v)
			}

			if cfg.GetGeotag() != nil {
				v := cfg.GetGeotag()

				et.GeoTime(*f.Timestamp)
				et.GeoTag(*v)
			}

			fmt.Println(et.Cmd())
		}
	}
}
