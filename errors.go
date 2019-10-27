package tagger

import "fmt"

// NewErrorWithSuffix ...
func NewErrorWithSuffix(err error, suffix string) error {
	return fmt.Errorf("%s\n%s", err.Error(), suffix)
}
