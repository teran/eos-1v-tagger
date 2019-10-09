package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tagger "github.com/teran/eos-1v-tagger"
)

func main() {
	t, err := tagger.NewCSVParser(os.Args[1])
	if err != nil {
		log.Fatalf("error initializing CSV parser: %s", err)
	}

	film, err := t.Parse()
	if err != nil {
		log.Fatalf("error parsing CSV: %s", err)
	}

	for _, f := range film.Frames {
		fmt.Printf(
			strings.Join([]string{
				`exiftool`,
				`-overwrite_original`,
				`-FocalLength="%dmm"`,
				`-ExposureTime="%s"`,
				`-ShutterSpeedValue="%s"`,
				`-FNumber="%v"`,
				`-ApertureValue="%v"`,
				`-ISO="%d"`,
				`-ISOSpeed="%d"`,
				`-ExposureCompensation="%v"`,
				`-MeteringMode="%v"`,
				`-ShootingMode="%v"`,
				`-FocusMode="%v"`,
				`-DateTimeOriginal="%s"`,
				`-ModifyDate="%s"`,
				`FILM_%05d.dng`,
			}, " ")+"\n",
			f.FocalLength,
			f.Tv,
			f.Tv,
			f.Av,
			f.Av,
			f.ISO,
			f.ISO,
			f.ExposureCompensation,
			f.MeteringMode,
			f.ShoothingMode,
			f.AFMode,
			f.Timestamp.Format(time.RFC3339),
			f.Timestamp.Format(time.RFC3339),
			f.Number,
		)
	}
}
