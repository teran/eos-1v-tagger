package tagger

import "time"

// LocationByTimeZone returns time.Location object by timezone name
func LocationByTimeZone(z string) (*time.Location, error) {
	return time.LoadLocation(z)
}
