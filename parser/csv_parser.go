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

	tagger "github.com/teran/eos-1v-tagger"
	types "github.com/teran/eos-1v-tagger/types"
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

	// ErrNotProvided ...
	ErrNotProvided = errors.New("value is not provided")
)

// New creates new CSVParser object
func New(fn string, tz *time.Location, timestampFormat string) (*CSVParser, error) {
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
func (p *CSVParser) Parse() ([]*types.Film, error) {
	rd := bufio.NewReader(p.rc)

	films := []*types.Film{}
	var f *types.Film
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
			if f != nil && !f.IsEmpty() {
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

			if fr.ISO == nil {
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

func parseFilmData(s string, tz *time.Location, timestmapFormat string) (*types.Film, error) {
	ss := strings.Split(s, ",")

	tt, err := parseTimestamp(ss[6], ss[7], tz, timestmapFormat)
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrap(err, "error parsing film timestamp value")
	}

	fc, err := parseInt(strings.TrimSpace(ss[9]))
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrap(err, "error parsing film frame count value")
	}

	iso, err := parseInt(strings.TrimSpace(ss[11]))
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrap(err, "error parsing film ISO value")
	}

	ids := strings.Split(ss[2], "-")
	if err != nil && err != ErrNotProvided {
		return nil, errors.New("improper film ID data")
	}

	fID, err := parseInt(ids[1])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrap(err, "error parsing film ID value")
	}

	cID, err := parseInt(ids[0])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrap(err, "error parsing camera ID value")
	}

	title, err := parseString(ss[4])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrap(err, "error parsing film title")
	}

	return &types.Film{
		ID:                  fID,
		CameraID:            cID,
		Title:               title,
		FilmLoadedTimestamp: tt,
		FrameCount:          fc,
		ISO:                 iso,
	}, nil
}

func parseFrameData(s string, tz *time.Location, timestampFormat string) (*types.Frame, error) {
	ss := strings.Split(s, ",")
	if len(ss) != 21 {
		return nil, fmt.Errorf("wrong amount of columns for frame: %d: `%s`", len(ss), s)
	}

	// ss[2:] is everything except flag and number fields
	if isEmptySliceOfStrings(ss[2:]) {
		return nil, ErrEmptyFrame
	}

	frameID, err := parseFrameID(ss[1])
	if err != nil {
		return nil, errors.Wrap(err, "error parsing frameID value")
	}

	flag, err := parseFlag(ss[0])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing flag value; frameNo=%d", *frameID)
	}

	focalLength, err := parseFocalLength(ss[2])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing focal length value; frameNo=%d", *frameID)
	}

	maxAperture, err := parseAperture(ss[3])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing max aperture value; frameNo=%d", *frameID)
	}

	tv, err := parseExposure(ss[4])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing exposure value; frameNo=%d", *frameID)
	}

	av, err := parseAperture(ss[5])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing AV value; frameNo=%d", *frameID)
	}

	iso, err := parseISO(ss[6])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing ISO value; frameNo=%d", *frameID)
	}

	expcomp, err := parseCompensation(ss[7])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing exposure compensation value; frameNo=%d", *frameID)
	}

	flashcomp, err := parseCompensation(ss[8])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing flash compensation value; frameNo=%d", *frameID)
	}

	timestamp, err := parseTimestamp(ss[15], ss[16], tz, timestampFormat)
	if err != nil && err != ErrNotProvided {
		return nil, tagger.NewErrorWithSuffix(
			err, "Possible solution: consider using `-timestamp-format` to specify proper format for timestamps")
	}

	batteryTimestamp, err := parseTimestamp(ss[18], ss[19], tz, timestampFormat)
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing timestamp value; frameNo=%d", *frameID)
	}

	flashMode, err := parseFlashMode(ss[9])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing flash mode value; frameNo=%d", *frameID)
	}

	meteringMode, err := parseMeteringMode(ss[10])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing metering mode value; frameNo=%d", *frameID)
	}

	shootingMode, err := parseShootingMode(ss[11])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing shooting mode value; frameNo=%d", *frameID)
	}

	filmAdvanceMode, err := parseFilmAdvanceMode(ss[12])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing film advance mode value; frameNo=%d", *frameID)
	}

	afMode, err := parseAFMode(ss[13])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing AF mode value; frameNo=%d", *frameID)
	}

	bulbExposureTime, err := parseBulbExposureTime(ss[14])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing AF mode value; frameNo=%d", *frameID)
	}

	multipleExposure, err := parseMultipleExposure(ss[17])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing multiple exposure value; frameNo=%d", *frameID)
	}

	remarks, err := parseRemarks(ss[20])
	if err != nil && err != ErrNotProvided {
		return nil, errors.Wrapf(err, "error parsing remarks value; frameNo=%d", *frameID)
	}

	f := &types.Frame{
		Flag:                 flag,
		Number:               frameID,
		FocalLength:          focalLength,
		MaxAperture:          maxAperture,
		Tv:                   tv,
		Av:                   av,
		ISO:                  iso,
		ExposureCompensation: expcomp,
		FlashCompensation:    flashcomp,
		FlashMode:            flashMode,
		MeteringMode:         meteringMode,
		ShootingMode:         shootingMode,
		FilmAdvanceMode:      filmAdvanceMode,
		AFMode:               afMode,
		BulbExposureTime:     bulbExposureTime,
		Timestamp:            timestamp,
		MultipleExposure:     multipleExposure,
		BatteryLoadedDate:    batteryTimestamp,
		Remarks:              remarks,
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

func parseTimestamp(d, t string, tz *time.Location, timestampFormat string) (*time.Time, error) {
	if strings.TrimSpace(d) == "" || strings.TrimSpace(t) == "" {
		return nil, ErrNotProvided
	}
	ts, err := time.ParseInLocation(timestampFormat, fmt.Sprintf("%vT%v", d, t), tz)
	if err != nil {
		return nil, err
	}
	return &ts, nil
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

func parseFilmRemarks(s string) *string {
	ss := strings.Split(s, ",")
	sss := strings.Join(ss[2:], ",")
	return &sss
}

func parseString(s string) (*string, error) {
	if strings.TrimSpace(s) == "" {
		return nil, ErrNotProvided
	}

	return &s, nil
}

func parseFloat(s string) (*float64, error) {
	if strings.TrimSpace(s) == "" {
		return nil, ErrNotProvided
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func parseInt(s string) (*int64, error) {
	if strings.TrimSpace(s) == "" {
		return nil, ErrNotProvided
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &i, err
}

func parseFrameID(s string) (*int64, error) {
	return parseInt(s)
}

func parseFlag(s string) (*bool, error) {
	var t bool
	if s == "*" {
		t = true
		return &t, nil
	}
	return &t, nil
}

func parseFocalLength(s string) (*int64, error) {
	l := strings.TrimRight(s, "mm")
	return parseInt(l)
}

func parseAperture(s string) (*float64, error) {
	return parseFloat(s)
}

func parseExposure(s string) (*string, error) {
	if strings.TrimSpace(s) == "" {
		return nil, ErrNotProvided
	}
	tv := s
	tv = strings.Replace(tv, `"`, "", -1)
	tv = strings.Replace(tv, "=", "", -1)
	return &tv, nil
}

func parseISO(s string) (*int64, error) {
	return parseInt(s)
}

func parseCompensation(s string) (*float64, error) {
	return parseFloat(s)
}

func parseFlashMode(s string) (*string, error) {
	return parseString(s)
}

func parseMeteringMode(s string) (*string, error) {
	return parseString(s)
}

func parseShootingMode(s string) (*string, error) {
	return parseString(s)
}

func parseFilmAdvanceMode(s string) (*string, error) {
	return parseString(s)
}

func parseAFMode(s string) (*string, error) {
	return parseString(s)
}

func parseBulbExposureTime(s string) (*string, error) {
	return parseString(s)
}

func parseMultipleExposure(s string) (*string, error) {
	return parseString(s)
}

func parseRemarks(s string) (*string, error) {
	return parseString(s)
}
