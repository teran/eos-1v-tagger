package types

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	_ flag.Value   = (*TimestampFormat)(nil)
	_ fmt.Stringer = (*TimestampFormat)(nil)
)

// TimestampFormat ...
type TimestampFormat string

var (
	// TimestampFormatEU ...
	TimestampFormatEU TimestampFormat = "EU"

	// TimestampFormatUS ...
	TimestampFormatUS TimestampFormat = "US"
)

// NewTimestampFormat ...
func NewTimestampFormat(s string) (*TimestampFormat, error) {
	switch TimestampFormat(s) {
	case TimestampFormatEU:
		v := TimestampFormatEU
		return &v, nil
	case TimestampFormatUS:
		v := TimestampFormatUS
		return &v, nil
	}
	return nil, errors.New("Invalid value for timestamp format")
}

// Set ...
func (tf *TimestampFormat) Set(value string) error {
	value = strings.TrimSpace(value)

	switch TimestampFormat(value) {
	case TimestampFormatEU:
		*tf = TimestampFormatEU
	case TimestampFormatUS:
		*tf = TimestampFormatUS
	default:
		return errors.Errorf("Unknown value `%s` for time format", value)
	}
	return nil
}

func (tf *TimestampFormat) String() string {
	return string(*tf)
}

// TimeLayout returns layout string for time.Format()
func (tf *TimestampFormat) TimeLayout() string {
	if *tf == TimestampFormatEU {
		return "2/1/2006T15:04:05"
	}
	return "1/2/2006T15:04:05"
}
