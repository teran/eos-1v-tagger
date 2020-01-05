package types

import (
	"strings"

	"github.com/pkg/errors"
)

// MultipleExposure ...
type MultipleExposure string

const (
	// MultipleExposureOn ...
	MultipleExposureOn MultipleExposure = "ON"
	// MultipleExposureOff ...
	MultipleExposureOff MultipleExposure = "OFF"
)

// MultipleExposureFromString ...
func MultipleExposureFromString(s string) (me *MultipleExposure, err error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, ErrEmptyValue
	}

	var mev MultipleExposure
	switch MultipleExposure(s) {
	case MultipleExposureOn:
		mev = MultipleExposureOn
	case MultipleExposureOff:
		mev = MultipleExposureOff
	default:
		return nil, errors.Errorf("error parsing MultipleExposure: unknown value `%s`", s)
	}

	return &mev, nil
}

func (me *MultipleExposure) String() string {
	return string(*me)
}
