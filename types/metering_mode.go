package types

import (
	"strings"

	"github.com/pkg/errors"
)

var _ EXIFValuer = (*MeteringMode)(nil)

// MeteringMode ...
type MeteringMode string

const (
	// MeteringModeEvaluative ...
	MeteringModeEvaluative MeteringMode = "Evaluative"

	// MeteringModePartial ...
	MeteringModePartial MeteringMode = "Partial"

	// MeteringModeSpot ...
	MeteringModeSpot MeteringMode = "Spot"

	// MeteringModeCenterAveraging ...
	MeteringModeCenterAveraging MeteringMode = "Center Averaging"
)

// MeteringModeFromString ...
func MeteringModeFromString(s string) (mm *MeteringMode, err error) {
	s = strings.TrimSpace(s)

	var mmv MeteringMode
	switch MeteringMode(s) {
	case MeteringModeEvaluative:
		mmv = MeteringModeEvaluative
	case MeteringModePartial:
		mmv = MeteringModePartial
	case MeteringModeSpot:
		mmv = MeteringModeSpot
	case MeteringModeCenterAveraging:
		mmv = MeteringModeCenterAveraging
	default:
		err = errors.Errorf("error parsing MeteringMode: unknown value `%s`", s)
	}

	return &mmv, err
}

func (mm *MeteringMode) String() string {
	return string(*mm)
}

// EXIFValue ...
func (mm *MeteringMode) EXIFValue() EXIFValue {
	return EXIFValue{
		"ExifIFD:MeteringMode":                mm.exifIFDValue(),
		"Canon:MeteringMode":                  mm.exifCanonValue(),
		"CanonCustom:PF2DisableMeteringModes": "Off",
	}
}

func (mm *MeteringMode) exifIFDValue() string {
	switch *mm {
	case MeteringModeEvaluative:
		return "Multi-Segment"
	case MeteringModePartial:
		return "Partial"
	case MeteringModeSpot:
		return "Spot"
	case MeteringModeCenterAveraging:
		return "Center-weighted average"
	}
	return "Unknown"
}

func (mm *MeteringMode) exifCanonValue() string {
	switch *mm {
	case MeteringModeEvaluative:
		return "Evaluative"
	case MeteringModePartial:
		return "Partial"
	case MeteringModeSpot:
		return "Spot"
	case MeteringModeCenterAveraging:
		return "Center-weighted average"
	}
	return "Unknown"
}
