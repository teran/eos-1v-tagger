package tagger

import "time"

// Film model to store all the data about the film itself
type Film struct {
	ID                  int64
	CameraID            int64
	Title               string
	FilmLoadedTimestamp time.Time
	FrameCount          int64
	ISO                 int64
	Remarks             string
	Frames              []Frame
}

// IsEmpty checks if Film object is empty
func (f Film) IsEmpty() bool {
	switch {
	case f.ID != 0:
		return false
	case f.CameraID != 0:
		return false
	case f.Title != "":
		return false
	case !f.FilmLoadedTimestamp.IsZero():
		return false
	case f.FrameCount != 0:
		return false
	case f.ISO != 0:
		return false
	case f.Remarks != "":
		return false
	case len(f.Frames) > 0:
		return false
	}
	return true
}
