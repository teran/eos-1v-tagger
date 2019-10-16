package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

// LD vars
var (
	ldVersion   = "undefined"
	ldTimestamp = "0"
)

// help message prefix/suffix
var (
	usagePrefix = "Usage: tagger [OPTIONS] file.csv\n\nOptions:\n"
	usageSuffix = fmt.Sprintf("\nVersion: %s, build with %s at %s\n", ldVersion, runtime.Version(), func() string {
		tsI, err := strconv.ParseInt(ldTimestamp, 10, 64)
		if err != nil {
			panic(err)
		}
		return time.Unix(tsI, 0).Format(time.RFC3339)
	}())
)

// Flags
var (
	displayHelp     bool   = false
	filenamePattern string = "FILM_%05d.dng"
	geotag          string = ""
	make            string = ""
	model           string = ""
	serialNumber    string = ""
	setDigitized    bool   = false
	timezone        string = "UTC"
)

func parseFlags() {
	flag.Usage = func() {
		fmt.Print(usagePrefix)
		flag.PrintDefaults()
		fmt.Print(usageSuffix)
	}

	flag.BoolVar(&displayHelp, "help", displayHelp, "display help message")
	flag.StringVar(&filenamePattern, "filename-pattern", filenamePattern, "filename pattern for generate exiftool command. %d means frame number on the film")
	flag.StringVar(&geotag, "geotag", geotag, "GPS track log file to set location data, supported formats are the ones supported by exiftool. Please refer to exiftool docs for details.")
	flag.StringVar(&make, "make", make, "Make tag value. NOTE: it will overwrite the value set by your film scanner software")
	flag.StringVar(&model, "model", model, "Model tag value. NOTE: it will overwrite the value set by your film scanner software")
	flag.StringVar(&serialNumber, "serial-number", serialNumber, "SerialNumber tag value. NOTE: it will overwrite the value set by your film scanner software")
	flag.BoolVar(&setDigitized, "set-digitized", setDigitized, "set DateTimeDigitized from CreateDate field")
	flag.StringVar(&timezone, "timezone", timezone, "location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow')")

	flag.Parse()

	if displayHelp || flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
}
