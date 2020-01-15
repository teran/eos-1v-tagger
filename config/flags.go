package config

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/teran/eos-1v-tagger/types"
)

type flags struct {
	displayHelp     bool
	copyright       string
	exiftoolBinary  string
	filenamePattern string
	fileSource      types.FileSource
	geotag          string
	make            string
	model           string
	serialNumber    string
	setDigitized    bool
	timestampFormat types.TimestampFormat
	timezone        string
	displayVersion  bool

	usagePrefix string
	usageSuffix string
}

// NewFlags ...
func NewFlags(binary, version, timestamp string) types.Flags {
	f := flags{
		usagePrefix: fmt.Sprintf("Usage: %v [OPTIONS] file.csv\n\nOptions:\n", binary),
		usageSuffix: fmt.Sprintf("Version: %s, build with %s at %s\n", version, runtime.Version(), func() string {
			tsI, err := strconv.ParseInt(timestamp, 10, 64)
			if err != nil {
				panic(err)
			}
			return time.Unix(tsI, 0).Format(time.RFC3339)
		}()),
	}

	flag.Usage = func() {
		fmt.Print(f.usagePrefix)
		flag.PrintDefaults()
		fmt.Print("\n")
		fmt.Print(f.usageSuffix)
	}

	flag.BoolVar(&f.displayHelp, "help", false, "display help message")
	flag.StringVar(&f.copyright, "copyright", "", "copyright notice for images")
	flag.StringVar(&f.exiftoolBinary, "exiftool-binary", "exiftool", "path to exiftool binary")
	flag.StringVar(&f.filenamePattern, "filename-pattern", `FILM_${cameraID:02d}${filmID:03d}${frameNo:05d}.dng`, "filename pattern for generate exiftool command. Available variables: frameNo, cameraID, filmID. More details are available in README.")
	flag.Var(&f.fileSource, "file-source", "adds file source EXIF tag. Available options: 'Film Scanner', 'Reflection Print Scanner', 'Digital Camera'")
	flag.StringVar(&f.geotag, "geotag", "", "GPS track log file to set location data, supported formats are the ones supported by exiftool. Please refer to exiftool docs for details.")
	flag.StringVar(&f.make, "make", "", "Make tag value. NOTE: it will overwrite the value set by your film scanner software")
	flag.StringVar(&f.model, "model", "", "Model tag value. NOTE: it will overwrite the value set by your film scanner software")
	flag.StringVar(&f.serialNumber, "serial-number", "", "SerialNumber tag value. NOTE: it will overwrite the value set by your film scanner software")
	flag.BoolVar(&f.setDigitized, "set-digitized", false, "set DateTimeDigitized from CreateDate field")
	flag.Var(&f.timestampFormat, "timestamp-format", "the timestamp format in the locale your're using on the system with ES-E1 software. Allowed values: 'US', 'EU'")
	flag.StringVar(&f.timezone, "timezone", "", "location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow'; default: 'UTC')")
	flag.BoolVar(&f.displayVersion, "version", false, "show program version")

	flag.Parse()

	return &f
}

func (f *flags) PrintUsageString() {
	flag.Usage()
}

func (f *flags) PrintVersionString() {
	fmt.Println(f.usageSuffix)
}

func (f *flags) GetDisplayHelp() bool {
	return f.displayHelp
}

func (f *flags) GetDisplayVersion() bool {
	return f.displayVersion
}

func (f *flags) GetCopyright() string {
	return f.copyright
}

func (f *flags) GetExiftoolBinary() string {
	return f.exiftoolBinary
}

func (f *flags) GetFilenamePattern() string {
	return f.filenamePattern
}

func (f *flags) GetFileSource() types.FileSource {
	return f.fileSource
}

func (f *flags) GetGeotag() string {
	return f.geotag
}

func (f *flags) GetMake() string {
	return f.make
}

func (f *flags) GetModel() string {
	return f.model
}

func (f *flags) GetSerialNumber() string {
	return f.serialNumber
}

func (f *flags) GetSetDigitized() bool {
	return f.setDigitized
}

func (f *flags) GetTimestampFormat() types.TimestampFormat {
	return f.timestampFormat
}

func (f *flags) GetTimezone() types.Timezone {
	return f.timezone
}

func (f *flags) GetCSVPath() string {
	return flag.Arg(0)
}
