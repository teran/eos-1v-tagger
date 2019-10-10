package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	usageMessage = "Usage: tagger [OPTIONS] file.csv\n\nOptions:\n"
)

var (
	displayHelp     bool   = false
	filenamePattern string = "FILM_%05d.dng"
	timezone        string = "UTC"
)

func parseFlags() {
	flag.Usage = func() {
		fmt.Print(usageMessage)

		flag.PrintDefaults()
	}

	flag.BoolVar(&displayHelp, "help", displayHelp, "display help message")
	flag.StringVar(&filenamePattern, "filename-pattern", filenamePattern, "filename pattern for generate exiftool command. %d means frame number on the film")
	flag.StringVar(&timezone, "timezone", timezone, "location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow')")

	flag.Parse()

	if displayHelp || flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
}
