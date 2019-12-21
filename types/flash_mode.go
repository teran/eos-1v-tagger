package types

import (
	"strings"

	"github.com/pkg/errors"
)

// FlashMode ...
type FlashMode string

const (
	// FlashModeOn ...
	FlashModeOn FlashMode = "ON"

	// FlashModeOff ...
	FlashModeOff FlashMode = "OFF"

	// FlashModeETTL ...
	FlashModeETTL FlashMode = "E-TTL"

	// FlashModeATTL ...
	FlashModeATTL FlashMode = "A-TTL"

	// FlashModeTTLAutoflash ...
	FlashModeTTLAutoflash FlashMode = "TTL autoflash"

	// FlashModeManualFlash ...
	FlashModeManualFlash FlashMode = "Manual flash"
)

// FlashModeFromString ...
func FlashModeFromString(s string) (fm *FlashMode, err error) {
	s = strings.TrimSpace(s)

	var fmv FlashMode
	switch FlashMode(s) {
	case FlashModeOn:
		fmv = FlashModeOn
	case FlashModeOff:
		fmv = FlashModeOff
	case FlashModeETTL:
		fmv = FlashModeETTL
	case FlashModeATTL:
		fmv = FlashModeATTL
	case FlashModeTTLAutoflash:
		fmv = FlashModeTTLAutoflash
	case FlashModeManualFlash:
		fmv = FlashModeManualFlash
	default:
		err = errors.Errorf("error parsing FlashMode: unknown value `%s`", s)
	}

	return &fmv, err
}

func (fm *FlashMode) String() string {
	return string(*fm)
}
