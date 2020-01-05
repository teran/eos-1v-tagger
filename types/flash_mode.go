package types

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	_ EXIFValuer   = (*FlashMode)(nil)
	_ fmt.Stringer = (*FlashMode)(nil)
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
	if s == "" {
		return nil, ErrEmptyValue
	}

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

func (fm *FlashMode) canonFlashMode() string {
	switch *fm {
	case FlashModeOff:
		return "Off"
	case FlashModeTTLAutoflash:
		return "Auto"
	}
	return "On"
}

func (fm *FlashMode) exifFlash() string {
	switch *fm {
	case FlashModeOff:
		return "Off, Did not fire"
	case FlashModeOn:
		return "On, Fired"
	case FlashModeManualFlash:
		return "On, Fired"
	}
	return "Auto, Fired"
}

func (fm *FlashMode) canonFlashBits() string {
	switch *fm {
	case FlashModeETTL:
		return "E-TTL"
	case FlashModeATTL:
		return "A-TTL"
	case FlashModeTTLAutoflash:
		return "TTL"
	case FlashModeManualFlash:
		return "Manual"
	}
	return "(none)"
}

// EXIFValue ...
func (fm *FlashMode) EXIFValue() EXIFValue {
	return EXIFValue{
		"ExifIFD:Flash":        fm.exifFlash(),
		"Canon:FlashBits":      fm.canonFlashBits(),
		"Canon:CanonFlashMode": fm.canonFlashMode(),
	}
}
