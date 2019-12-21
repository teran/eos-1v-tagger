package types

import (
	"strings"

	"github.com/pkg/errors"
)

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
