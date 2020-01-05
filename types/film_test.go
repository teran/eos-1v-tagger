package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFilmIsEmpty(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name       string
		filmSample Film
		expResult  bool
	}

	tcs := []testCase{
		{
			name:       "id filled",
			filmSample: Film{ID: ptrInt64(1234)},
			expResult:  false,
		},
		{
			name:       "camera id filled",
			filmSample: Film{CameraID: ptrInt64(1234)},
			expResult:  false,
		},
		{
			name:       "title filled",
			filmSample: Film{Title: ptrString("test titme")},
			expResult:  false,
		},
		{
			name:       "timestamp filled",
			filmSample: Film{FilmLoadedTimestamp: ptrTime(time.Now())},
			expResult:  false,
		},
		{
			name:       "frame count filled",
			filmSample: Film{FrameCount: ptrInt64(1234)},
			expResult:  false,
		},
		{
			name:       "iso filled",
			filmSample: Film{ISO: ptrInt64(1234)},
			expResult:  false,
		},
		{
			name:       "string filled",
			filmSample: Film{Remarks: ptrString("blah")},
			expResult:  false,
		},
		{
			name:       "frames are present",
			filmSample: Film{Frames: []*Frame{{}}},
			expResult:  false,
		},
		{
			name:       "empty film",
			filmSample: Film{},
			expResult:  true,
		},
	}

	for _, tc := range tcs {
		r.Equalf(tc.expResult, tc.filmSample.IsEmpty(), tc.name)
	}
}

func ptrInt64(t int64) *int64        { return &t }
func ptrString(t string) *string     { return &t }
func ptrTime(t time.Time) *time.Time { return &t }
