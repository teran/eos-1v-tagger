package types

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	_ flag.Value   = (*FileSource)(nil)
	_ fmt.Stringer = (*FileSource)(nil)
)

// FileSource ...
type FileSource string

const (
	// FileSourceFilmScanner ...
	FileSourceFilmScanner FileSource = "Film Scanner"

	// FileSourceReflectionPrintScanner ...
	FileSourceReflectionPrintScanner FileSource = "Reflection Print Scanner"

	// FileSourceDigitalCamera ...
	FileSourceDigitalCamera FileSource = "Digital Camera"
)

// Set is a part of flag.Value implementation
func (fs *FileSource) Set(value string) error {
	value = strings.TrimSpace(value)

	switch FileSource(value) {
	case FileSourceFilmScanner:
		*fs = FileSourceFilmScanner
	case FileSourceReflectionPrintScanner:
		*fs = FileSourceReflectionPrintScanner
	case FileSourceDigitalCamera:
		*fs = FileSourceDigitalCamera
	default:
		return errors.Errorf("Unknown value `%s` for time format", value)
	}
	return nil
}

// String is a part of fmt.Stringer and flag.Value implementation
func (fs *FileSource) String() string {
	return string(*fs)
}
