package tagger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// CSVParser type
type CSVParser struct {
	rc io.ReadCloser
}

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
			log.Printf("error parsing frame data: %s: `%s`", err, frameStr)
			continue
		}

		film.Frames = append(film.Frames, frame)
	}

	j, err := json.MarshalIndent(film, "", "  ")
	if err != nil {
		return film, err
	}

	fmt.Println(string(j))

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

	f := Frame{
		Flag:        flag,
		Number:      frameID,
		FocalLength: focalLength,
		MaxAperture: maxAperture,
		Tv:          ss[4],
		Av:          av,
		ISO:         iso,
	}
	return f, nil
}
