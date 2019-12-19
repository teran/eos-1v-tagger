package tagger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// CSVParser type
type CSVParser struct {
	rc              io.ReadCloser
	tz              *time.Location
	timestampFormat string
}

var (
	// ErrEmptyFrame ...
	ErrEmptyFrame = errors.New("frame line contains no data")
)

// NewCSVParser creates new CSVParser object
func NewCSVParser(fn string, tz *time.Location, timestampFormat string) (*CSVParser, error) {
	fp, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	return &CSVParser{
		rc:              fp,
		tz:              tz,
		timestampFormat: timestampFormat,
	}, nil
}

// Close ...
func (p *CSVParser) Close() error {
	return p.rc.Close()
}

// Parse ...
func (p *CSVParser) Parse() ([]Film, error) {
	rd := bufio.NewReader(p.rc)

	films := []Film{}
	f := Film{}
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}

		switch {
		case isFilmHeader(str):
			if !f.IsEmpty() {
				films = append(films, f)
			}

			f, err = parseFilmData(str, p.tz, p.timestampFormat)
			if err != nil {
				return nil, err
			}
			break
		case isFilmRemarksHeader(str):
			f.Remarks = parseFilmRemarks(str)
		case isFrameHeader(str):
			continue
		default:
			fr, err := parseFrameData(str, p.tz, p.timestampFormat)
			if err != nil {
				return nil, err
			}

			if fr.ISO == 0 {
				fr.ISO = f.ISO
			}

			f.Frames = append(f.Frames, fr)
		}
	}

	if !f.IsEmpty() {
		films = append(films, f)
	}

	return films, nil
}

func parseFilmData(s string, tz *time.Location, timestmapFormat string) (Film, error) {
	ss := strings.Split(s, ",")

	ts := fmt.Sprintf("%sT%s", ss[6], ss[7])
	tt, err := time.ParseInLocation(timestmapFormat, ts, tz)
	if err != nil {
		return Film{}, err
	}

	fc, err := strconv.ParseInt(strings.TrimSpace(ss[9]), 10, 64)
	if err != nil {
		return Film{}, err
	}

	iso, err := strconv.ParseInt(strings.TrimSpace(ss[11]), 10, 64)
	if err != nil {
		return Film{}, err
	}

	ids := strings.Split(ss[2], "-")
	if len(ids) != 2 {
		return Film{}, errors.New("improper film ID data")
	}

	fID, err := strconv.ParseInt(ids[1], 10, 64)
	if err != nil {
		return Film{}, err
	}

	cID, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		return Film{}, err
	}

	f := Film{
		ID:                  fID,
		CameraID:            cID,
		Title:               ss[4],
		FilmLoadedTimestamp: tt,
		FrameCount:          fc,
		ISO:                 iso,
	}

	return f, nil
}

func parseFrameData(s string, tz *time.Location, timestampFormat string) (*Frame, error) {
	ss := strings.Split(s, ",")
	if len(ss) != 21 {
		return nil, fmt.Errorf("wrong amount of columns for frame: %d: `%s`", len(ss), s)
	}

	// ss[2:] is everything except flag and number fields
	if isEmptySliceOfStrings(ss[2:]) {
		return nil, ErrEmptyFrame
	}

	frameID, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing frameID value")
	}

	flag := func() bool {
		if ss[0] == "*" {
			return true
		}
		return false
	}()

	focalLength, err := func() (int64, error) {
		l := strings.TrimRight(ss[2], "mm")
		return strconv.ParseInt(l, 10, 64)
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing focal length value; frameNo=%d", frameID)
	}

	maxAperture, err := strconv.ParseFloat(ss[3], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing max aperture value; frameNo=%d", frameID)
	}

	tv, err := func() (string, error) {
		tv := ss[4]
		tv = strings.Replace(tv, `"`, "", -1)
		tv = strings.Replace(tv, "=", "", -1)
		return tv, nil
	}()

	av, err := strconv.ParseFloat(ss[5], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing AV value; frameNo=%d", frameID)
	}

	iso, err := func() (int64, error) {
		if ss[6] == "" {
			return 0, nil
		}
		return strconv.ParseInt(ss[6], 10, 64)
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing ISO value; frameNo=%d", frameID)
	}

	expcomp, err := strconv.ParseFloat(ss[7], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing exposure compensation value; frameNo=%d", frameID)
	}

	flashcomp, err := strconv.ParseFloat(ss[8], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing flash compensation value; frameNo=%d", frameID)
	}

	timestamp, err := parseTimestamp(ss[15], ss[16], tz, timestampFormat)
	if err != nil {
		return nil, NewErrorWithSuffix(err, "Possible solution: consider using `-timestamp-format` to specify proper format for timestamps")
	}

	batteryTimestamp, err := parseTimestamp(ss[18], ss[19], tz, timestampFormat)
	if err != nil {
		batteryTimestamp = time.Time{}
	}

	f := &Frame{
		Flag:                 flag,
		Number:               frameID,
		FocalLength:          focalLength,
		MaxAperture:          maxAperture,
		Tv:                   tv,
		Av:                   av,
		ISO:                  iso,
		ExposureCompensation: expcomp,
		FlashCompensation:    flashcomp,
		FlashMode:            ss[9],
		MeteringMode:         ss[10],
		ShootingMode:         ss[11],
		FilmAdvanceMode:      ss[12],
		AFMode:               ss[13],
		BulbExposureTime:     ss[14],
		Timestamp:            timestamp,
		MultipleExposure:     ss[17],
		BatteryLoadedDate:    batteryTimestamp,
		Remarks:              ss[20],
	}
	return f, nil
}

func isEmptySliceOfStrings(ss []string) bool {
	for _, s := range ss {
		if strings.TrimSpace(s) != "" {
			return false
		}
	}
	return true
}

func parseTimestamp(d, t string, tz *time.Location, timestampFormat string) (time.Time, error) {
	if d == "" || t == "" {
		return time.Time{}, nil
	}

	return time.ParseInLocation(timestampFormat, fmt.Sprintf("%vT%v", d, t), tz)
}

func isFilmHeader(s string) bool {
	return strings.HasPrefix(strings.TrimLeft(s, "*"), ",Film ID,")
}

func isFilmRemarksHeader(s string) bool {
	return strings.HasPrefix(s, ",Remarks,")
}

func isFrameHeader(s string) bool {
	return strings.HasPrefix(strings.TrimLeft(s, "*"), ",Frame No.,")
}

func parseFilmRemarks(s string) string {
	ss := strings.Split(s, ",")
	return strings.Join(ss[2:], ",")
}
