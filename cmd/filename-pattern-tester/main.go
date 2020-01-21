package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	format "github.com/teran/eos-1v-tagger/format"
)

// LD vars
var (
	ldVersion   = "undefined"
	ldTimestamp = "0"
)

func main() {
	var (
		pattern        string
		startCameraID  int
		endCameraID    int
		startFilmID    int
		endFilmID      int
		startFrameNo   int
		endFrameNo     int
		displayVersion bool
	)

	versionString := fmt.Sprintf("Version: %s, build with %s at %s\n", ldVersion, runtime.Version(), func() string {
		tsI, err := strconv.ParseInt(ldTimestamp, 10, 64)
		if err != nil {
			panic(err)
		}
		return time.Unix(tsI, 0).Format(time.RFC3339)
	}())

	flag.Usage = func() {
		fmt.Printf("Trivial tool to test filename pattern\n")
		fmt.Printf("The tool allows to generate amount of cameraID's, filmID's and frameID's to visually check the pattern\n\n")
		flag.PrintDefaults()
		fmt.Print(versionString)
	}

	flag.StringVar(&pattern, "filename-pattern", "FILM_${cameraID:02d}${filmID:03d}${frameNo:05d}.dng", "filename pattern to render")
	flag.IntVar(&startCameraID, "start-camera-id", 9, "generate cameraID's starting this ID")
	flag.IntVar(&endCameraID, "end-camera-id", 11, "generate cameraID's to this ID")
	flag.IntVar(&startFilmID, "start-film-id", 99, "generate filmID's starting this ID")
	flag.IntVar(&endFilmID, "end-film-id", 101, "generate filmID's to this ID")
	flag.IntVar(&startFrameNo, "start-frame-no", 9, "generate frameNo's starting this No")
	flag.IntVar(&endFrameNo, "end-frame-no", 11, "generate frameNo's to this No")
	flag.BoolVar(&displayVersion, "version", false, "display version and exit")
	flag.Parse()

	if displayVersion {
		fmt.Print(versionString)
		os.Exit(1)
	}

	fmt.Printf("rendering pattern: '%s'\n\n", pattern)

	for cameraID := startCameraID; cameraID < endCameraID; cameraID++ {
		fmt.Printf("cameraID = %d\n", cameraID)
		for filmID := startFilmID; filmID < endFilmID; filmID++ {
			fmt.Printf("  filmID = %d\n", filmID)
			for frameNo := startFrameNo; frameNo < endFrameNo; frameNo++ {
				filename := format.Format(pattern, map[string]interface{}{
					"filmID":   filmID,
					"cameraID": cameraID,
					"frameNo":  frameNo,
				})
				fmt.Printf("    %d: %s\n", frameNo, filename)
			}
		}
	}
}
