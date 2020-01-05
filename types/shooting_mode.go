package types

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	_ EXIFValuer   = (*ShootingMode)(nil)
	_ fmt.Stringer = (*ShootingMode)(nil)
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
	if s == "" {
		return nil, ErrEmptyValue
	}

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

// EXIFValue ...
func (sm *ShootingMode) EXIFValue() EXIFValue {
	return EXIFValue{
		"ExifIFD:ExposureProgram":             sm.exifIFDValue(),
		"Canon:CanonExposureMode":             sm.exifCanonValue(),
		"CanonCustom:PF1DisableShootingModes": "Off",
		"CanonCustom:PF6PresetShootingModes":  "Off",
	}
}

func (sm *ShootingMode) exifIFDValue() string {
	switch *sm {
	case ShootingModeProgramAE:
		return "Program AE"
	case ShootingModeShutterSpeedPriorityAE:
		return "Shutter speed priority AE"
	case ShootingModeAperturePriorityAE:
		return "Aperture-priority AE"
	case ShootingModeDepthOfFieldAE:
		return "Not Defined"
	case ShootingModeManualExposure:
		return "Manual"
	case ShootingModeBulb:
		return "Bulb"
	}
	return "Not Defined"
}

func (sm *ShootingMode) exifCanonValue() string {
	switch *sm {
	case ShootingModeProgramAE:
		return "Program AE"
	case ShootingModeShutterSpeedPriorityAE:
		return "Shutter speed priority AE"
	case ShootingModeAperturePriorityAE:
		return "Aperture-priority AE"
	case ShootingModeDepthOfFieldAE:
		return "Depth-of-field AE"
	case ShootingModeManualExposure:
		return "Manual"
	case ShootingModeBulb:
		return "Bulb"
	}
	return "Easy"
}
