package types

import (
	"strings"

	"github.com/pkg/errors"
)

// AFMode ...
type AFMode string

const (
	// AFModeOneShotAF ...
	AFModeOneShotAF AFMode = "One-Shot AF"

	// AFModeAIServoAF ...
	AFModeAIServoAF AFMode = "AI Servo AF"

	// AFModeManualFocus ...
	AFModeManualFocus AFMode = "Manual focus"
)

// AFModeFromString ...
func AFModeFromString(s string) (am *AFMode, err error) {
	s = strings.TrimSpace(s)

	var amv AFMode
	switch AFMode(s) {
	case AFModeOneShotAF:
		amv = AFModeOneShotAF
	case AFModeAIServoAF:
		amv = AFModeAIServoAF
	case AFModeManualFocus:
		amv = AFModeManualFocus
	default:
		return nil, errors.Errorf("error parsing AFMode: unknown value `%s`", s)
	}

	return &amv, nil
}

func (am *AFMode) String() string {
	return string(*am)
}