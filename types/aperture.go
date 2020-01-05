package types

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	_ EXIFValuer   = (*Aperture)(nil)
	_ fmt.Stringer = (*Aperture)(nil)
)

// Aperture ...
type Aperture float64

// ApertureFromString ...
func ApertureFromString(s string) (*Aperture, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, ErrEmptyValue
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	av := Aperture(f)

	return &av, nil
}

// String ...
func (a *Aperture) String() string {
	return fmt.Sprintf("%.1f", *a)
}

// EXIFValue ...
func (a *Aperture) EXIFValue() EXIFValue {
	return EXIFValue{
		"ExifIFD:FNumber":       a.String(),
		"ExifIFD:ApertureValue": a.String(),
		"Canon:FNumber":         a.String(),
		"Canon:TargetAperture":  a.String(),
	}
}
