package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	flagM "github.com/teran/eos-1v-tagger/config/mocks/flags"
	"github.com/teran/eos-1v-tagger/types"
)

func TestConfigFromYAML(t *testing.T) {
	r := require.New(t)

	cfg := NewDefaultConfig()

	err := cfg.FillFromYaml("./testdata/config.yaml")
	r.NoError(err)

	r.Equal(&config{
		copyright:       types.PtrString("Test Copyright Value"),
		exiftoolBinary:  "/usr/local/bin/exiftool",
		filenamePattern: "XXX_${cameraID:02d}${filmID:03d}${frameNo:05d}.dng",
		fileSource:      func() *types.FileSource { t := types.FileSourceFilmScanner; return &t }(),
		make:            map[uint8]string{9: "Canon"},
		model:           map[uint8]string{9: "Canon EOS 1V"},
		serialNumber:    map[uint8]string{9: "XXXYYYZZZ"},
		setDigitized:    true,
		timestampFormat: types.TimestampFormatEU,
		timezone:        map[uint8]string{0: "Europe/Paris", 9: "Europe/Moscow"},
	}, cfg)
}

func TestConfigFromFlags(t *testing.T) {
	r := require.New(t)

	m := flagM.New()

	m.On("GetDisplayHelp").Return(true).Twice()
	m.On("GetDisplayVersion").Return(true).Twice()
	m.On("GetCopyright").Return("test copyright from flags").Twice()
	m.On("GetExiftoolBinary").Return("/opt/local/bin/exiftool").Twice()
	m.On("GetFileSource").Return(types.FileSourceDigitalCamera).Twice()
	m.On("GetFilenamePattern").Return("blah").Twice()
	m.On("GetGeotag").Return("blah.gpx").Twice()
	m.On("GetMake").Return("blah vendor").Twice()
	m.On("GetModel").Return("blah model").Twice()
	m.On("GetSerialNumber").Return("ZZZZZZZZZ").Twice()
	m.On("GetSetDigitized").Return(true).Twice()
	m.On("GetTimestampFormat").Return(types.TimestampFormatEU).Twice()
	m.On("GetTimezone").Return("Europe/Berlin").Twice()

	cfg := NewDefaultConfig()
	err := cfg.FillFromYaml("./testdata/config.yaml")
	r.NoError(err)

	err = cfg.FillFromFlags(m)
	r.NoError(err)

	r.Equal(&config{
		displayHelp:     true,
		displayVersion:  true,
		copyright:       types.PtrString("test copyright from flags"),
		exiftoolBinary:  "/opt/local/bin/exiftool",
		filenamePattern: "blah",
		fileSource:      types.PtrFileSource(types.FileSourceDigitalCamera),
		geotag:          types.PtrString("blah.gpx"),
		make:            map[uint8]string{0: "blah vendor"},
		model:           map[uint8]string{0: "blah model"},
		serialNumber:    map[uint8]string{0: "ZZZZZZZZZ"},
		setDigitized:    true,
		timestampFormat: types.TimestampFormatEU,
		timezone:        map[uint8]string{0: "Europe/Berlin"},
	}, cfg)
}

func TestGetters(t *testing.T) {
	r := require.New(t)

	m := flagM.New()

	m.On("GetDisplayHelp").Return(true).Twice()
	m.On("GetDisplayVersion").Return(true).Twice()
	m.On("GetCopyright").Return("test copyright").Twice()
	m.On("GetExiftoolBinary").Return("/opt/local/bin/exiftool").Twice()
	m.On("GetFileSource").Return(types.FileSourceDigitalCamera).Twice()
	m.On("GetFilenamePattern").Return("blah").Twice()
	m.On("GetGeotag").Return("blah.gpx").Twice()
	m.On("GetMake").Return("Blah Vendor").Twice()
	m.On("GetModel").Return("Blah Model").Twice()
	m.On("GetSerialNumber").Return("ZZZZZZZZZ").Twice()
	m.On("GetSetDigitized").Return(true).Twice()
	m.On("GetTimestampFormat").Return(types.TimestampFormatEU).Twice()
	m.On("GetTimezone").Return("Europe/Berlin").Twice()

	cfg := NewDefaultConfig()
	err := cfg.FillFromYaml("./testdata/config.yaml")
	r.NoError(err)

	err = cfg.FillFromFlags(m)
	r.NoError(err)

	r.Equal(true, cfg.GetDisplayHelp())
	r.Equal(true, cfg.GetDisplayVersion())
	r.Equal(types.PtrString("test copyright"), cfg.GetCopyright())
	r.Equal("/opt/local/bin/exiftool", cfg.GetExiftoolBinary())
	r.Equal(types.PtrFileSource(types.FileSourceDigitalCamera), cfg.GetFileSource())
	r.Equal("blah", cfg.GetFilenamePattern())
	r.Equal(types.PtrString("blah.gpx"), cfg.GetGeotag())
	r.Equal(types.PtrString("Blah Vendor"), cfg.GetMakeByCameraID(0))
	r.Equal(types.PtrString("Blah Model"), cfg.GetModelByCameraID(0))
	r.Equal(types.PtrString("ZZZZZZZZZ"), cfg.GetSerialNumberByCameraID(0))
	r.Equal(true, cfg.GetSetDigitized())
	r.Equal(types.PtrTimestampFormat(types.TimestampFormatEU), cfg.GetTimestampFormat())
	r.Equal("Europe/Berlin", cfg.GetTimezoneByCameraID(0))
}

func TestGetOrDefault(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name         string
		data         map[uint8]string
		key          uint8
		defaultKey   uint8
		defaultValue string
		expResult    string
	}

	tcs := []testCase{
		{
			name:         "get by key",
			data:         map[uint8]string{0: "blah", 1: "blah2", 2: "blah3"},
			key:          1,
			defaultKey:   2,
			defaultValue: "default value",
			expResult:    "blah2",
		},
		{
			name:         "get by default key",
			data:         map[uint8]string{0: "blah", 1: "blah2", 2: "blah3"},
			key:          3, // not existent
			defaultKey:   2,
			defaultValue: "default value",
			expResult:    "blah3",
		},
		{
			name:         "get default value",
			data:         map[uint8]string{},
			key:          3, // not existent
			defaultKey:   2, // not existent
			defaultValue: "default value",
			expResult:    "default value",
		},
	}

	for _, tc := range tcs {
		v := getOrDefault(tc.data, tc.key, tc.defaultKey, tc.defaultValue)
		r.Equalf(tc.expResult, v, tc.name)
	}
}
