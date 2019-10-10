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
	filenamePattern string = "FILM_%05d.dng"
	displayHelp     bool   = false
)

func parseFlags() {
	flag.Usage = func() {
		fmt.Print(usageMessage)

		flag.PrintDefaults()
	}

	flag.StringVar(&filenamePattern, "filename-pattern", filenamePattern, "filename pattern for generate exiftool command. %d means frame number on the film.")
	flag.BoolVar(&displayHelp, "help", displayHelp, "display help message")

	flag.Parse()

	if displayHelp || flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
}
