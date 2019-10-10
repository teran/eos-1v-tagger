package main

import (
	"flag"
	"fmt"
	"log"

	tagger "github.com/teran/eos-1v-tagger"
)

func main() {
	parseFlags()

	t, err := tagger.NewCSVParser(flag.Arg(0))
	if err != nil {
		log.Fatalf("error initializing CSV parser: %s", err)
	}

	film, err := t.Parse()
	if err != nil {
		log.Fatalf("error parsing CSV: %s", err)
	}

	for _, f := range film.Frames {
		et := tagger.NewExifTool(fmt.Sprintf(filenamePattern, f.Number))

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

		fmt.Println(et.Cmd())
	}
}
