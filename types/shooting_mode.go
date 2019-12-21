package types

import (
	"strings"

	"github.com/pkg/errors"
)

// ShootingMode ...
type ShootingMode string

const (
	// ShootingModeProgramAE ...
	ShootingModeProgramAE ShootingMode = "Program AE"

	// ShootingModeShutterSpeedPriorityAE ...
	ShootingModeShutterSpeedPriorityAE ShootingMode = "Shutter-speed-priority AE"

	// ShootingModeAperturePriorityAE ...
	ShootingModeAperturePriorityAE ShootingMode = "Aperture-priority AE"

	// ShootingModeDepthOfFieldAE ...
	ShootingModeDepthOfFieldAE ShootingMode = "Depth-of-field AE"

	// ShootingModeManualExposure ...
	ShootingModeManualExposure ShootingMode = "Manual exposure"

	// ShootingModeBulb ...
	ShootingModeBulb ShootingMode = "Bulb"
)

// ShootingModeFromString ...
func ShootingModeFromString(s string) (sm *ShootingMode, err error) {
	s = strings.TrimSpace(s)

	var smv ShootingMode
	switch ShootingMode(s) {
	case ShootingModeProgramAE:
		smv = ShootingModeProgramAE
	case ShootingModeShutterSpeedPriorityAE:
		smv = ShootingModeShutterSpeedPriorityAE
	case ShootingModeAperturePriorityAE:
		smv = ShootingModeAperturePriorityAE
	case ShootingModeDepthOfFieldAE:
		smv = ShootingModeDepthOfFieldAE
	case ShootingModeManualExposure:
		smv = ShootingModeManualExposure
	case ShootingModeBulb:
		smv = ShootingModeBulb
	default:
		return nil, errors.Errorf("error parsing ShootingMode: unknown value `%s`", s)
	}

	return &smv, nil
}

func (sm *ShootingMode) String() string {
	return string(*sm)
}
