package tagger

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// CSVParser type
type CSVParser struct {
	rc io.ReadCloser
}

var (
	// ErrEmptyFrame ...
	ErrEmptyFrame = errors.New("frame line contains no data")
)

// NewCSVParser creates new CSVParser object
func NewCSVParser(fn string) (*CSVParser, error) {
	fp, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	return &CSVParser{
		rc: fp,
	}, nil
}

// NewCSVParserFromReadCloser creats new CSVParser object from ReadCloser
func NewCSVParserFromReadCloser(fp io.ReadCloser) *CSVParser {
	return &CSVParser{
		rc: fp,
	}
}

// Close ...
func (p *CSVParser) Close() error {
	return p.rc.Close()
}

// Parse ...
func (p *CSVParser) Parse() (Film, error) {
	rd := bufio.NewReader(p.rc)

	filmDataStr, err := rd.ReadString('\n')
	if err != nil {
		return Film{}, err
	}

	remarks, err := rd.ReadString('\n')
	if err != nil {
		return Film{}, err
	}
	remarks = strings.TrimSpace(strings.Split(remarks, ",")[2])

	film, err := parseFilmData(filmDataStr)
	if err != nil {
		return film, err
	}
	film.Remarks = remarks

	for {
		frameStr, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return film, err
		}
		frameStr = strings.TrimSpace(frameStr)
		if frameStr == "" || frameStr == frameHeader {
			continue
		}

		frame, err := parseFrameData(frameStr)
		if err != nil {
			if err == ErrEmptyFrame {
				continue
			}
			return Film{}, err
		}
		frame.ISO = film.ISO

		film.Frames = append(film.Frames, frame)
	}

	return film, err
}

func parseFilmData(s string) (Film, error) {
	ss := strings.Split(s, ",")

	ts := fmt.Sprintf("%sT%s", ss[6], ss[7])
	tt, err := time.Parse(TimestampFormat, ts)
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

	f := Film{
		ID:                  ss[2],
		Title:               ss[4],
		FilmLoadedTimestamp: tt,
		FrameCount:          fc,
		ISO:                 iso,
	}

	return f, nil
}

func parseFrameData(s string) (Frame, error) {
	ss := strings.Split(s, ",")
	if len(ss) != 21 {
		return Frame{}, fmt.Errorf("wrong amount of columns for frame: %d: `%s`", len(ss), s)
	}

	// ss[2:] is everything except flag and number fields
	if isEmptySliceOfStrings(ss[2:]) {
		return Frame{}, ErrEmptyFrame
	}

	frameID, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		return Frame{}, err
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
		return Frame{}, err
	}

	maxAperture, err := strconv.ParseFloat(ss[3], 64)
	if err != nil {
		return Frame{}, err
	}

	av, err := strconv.ParseFloat(ss[5], 64)
	if err != nil {
		return Frame{}, err
	}

	iso, err := func() (int64, error) {
		if ss[6] == "" {
			return 0, nil
		}
		return strconv.ParseInt(ss[6], 10, 64)
	}()
	if err != nil {
		return Frame{}, err
	}

	tv, err := func() (string, error) {
		tv := ss[4]
		tv = strings.Replace(tv, `"`, "", -1)
		tv = strings.Replace(tv, "=", "", -1)
		return tv, nil
	}()

	f := Frame{
		Flag:        flag,
		Number:      frameID,
		FocalLength: focalLength,
		MaxAperture: maxAperture,
		Tv:          tv,
		Av:          av,
		ISO:         iso,
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
