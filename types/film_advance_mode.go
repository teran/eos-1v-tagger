package types

import (
	"strings"

	"github.com/pkg/errors"
)

// FilmAdvanceMode ...
type FilmAdvanceMode string

const (
	// FilmAdvanceModeSingleFrame ...
	FilmAdvanceModeSingleFrame FilmAdvanceMode = "Single-frame"

	// FilmAdvanceModeContinuousBodyOnly ...
	FilmAdvanceModeContinuousBodyOnly FilmAdvanceMode = "Continuous (body only)"

	// FilmAdvanceModeLowSpeedContinuous ...
	FilmAdvanceModeLowSpeedContinuous FilmAdvanceMode = "Low-speed continuous"

	// FilmAdvanceModeHighSpeedContinuous ...
	FilmAdvanceModeHighSpeedContinuous FilmAdvanceMode = "High-speed continuous"

	// FilmAdvanceModeUltraHighSpeedContinuous ...
	FilmAdvanceModeUltraHighSpeedContinuous FilmAdvanceMode = "Ultra-high-speed continuous"

	// FilmAdvanceMode2secSelfTimer ...
	FilmAdvanceMode2secSelfTimer FilmAdvanceMode = "2-sec. self-timer"

	// FilmAdvanceMode10secSelfTimer ...
	FilmAdvanceMode10secSelfTimer FilmAdvanceMode = "10-sec. self-timer"
)

// FilmAdvanceModeFromString ...
func FilmAdvanceModeFromString(s string) (fam *FilmAdvanceMode, err error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, ErrEmptyValue
	}

	var famv FilmAdvanceMode
	switch FilmAdvanceMode(s) {
	case FilmAdvanceMode10secSelfTimer:
		famv = FilmAdvanceMode10secSelfTimer
	case FilmAdvanceMode2secSelfTimer:
		famv = FilmAdvanceMode2secSelfTimer
	case FilmAdvanceModeContinuousBodyOnly:
		famv = FilmAdvanceModeContinuousBodyOnly
	case FilmAdvanceModeHighSpeedContinuous:
		famv = FilmAdvanceModeHighSpeedContinuous
	case FilmAdvanceModeLowSpeedContinuous:
		famv = FilmAdvanceModeLowSpeedContinuous
	case FilmAdvanceModeSingleFrame:
		famv = FilmAdvanceModeSingleFrame
	case FilmAdvanceModeUltraHighSpeedContinuous:
		famv = FilmAdvanceModeUltraHighSpeedContinuous
	default:
		err = errors.Errorf("error parsing FilmAdvanceMode: unknown value `%s`", s)
	}

	return &famv, err
}

func (famv *FilmAdvanceMode) String() string {
	return string(*famv)
}
