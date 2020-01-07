package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/teran/eos-1v-tagger/types"
)

// Config ...
type config struct {
	displayHelp     bool
	displayVersion  bool
	copyright       *string
	exiftoolBinary  string
	filenamePattern string
	fileSource      *types.FileSource
	geotag          *string
	make            map[uint8]string
	model           map[uint8]string
	serialNumber    map[uint8]string
	setDigitized    bool
	timestampFormat types.TimestampFormat
	timezone        map[uint8]types.Timezone
}

// YamlConfig ...
type YamlConfig struct {
	Copyright       *string                  `yaml:"copyright"`
	ExiftoolBinary  *string                  `yaml:"exiftool-binary"`
	FilenamePattern *string                  `yaml:"filename-pattern"`
	FileSource      *types.FileSource        `yaml:"file-source"`
	Make            map[uint8]string         `yaml:"make"`
	Model           map[uint8]string         `yaml:"model"`
	SerialNumber    map[uint8]string         `yaml:"serial-number"`
	SetDigitized    *bool                    `yaml:"set-digitized"`
	TimestampFormat *types.TimestampFormat   `yaml:"timestamp-format"`
	Timezone        map[uint8]types.Timezone `yaml:"timezone"`
}

// NewDefaultConfig ...
func NewDefaultConfig() types.Config {
	c := &config{
		exiftoolBinary:  "exiftool",
		filenamePattern: `FILM_${cameraID:02d}${filmID:03d}${frameNo:05d}.dng`,
		timestampFormat: types.TimestampFormatUS,
		timezone:        map[uint8]string{0: "UTC"},
	}

	return c
}

// FillFromYAMLFile ...
func (c *config) FillFromYaml(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	var ycfg YamlConfig
	err = yaml.NewDecoder(fp).Decode(&ycfg)
	if err != nil {
		return err
	}
	return c.fillFromYamlConfig(ycfg)
}

func (c *config) fillFromYamlConfig(ycfg YamlConfig) error {
	if ycfg.Copyright != nil {
		c.copyright = ycfg.Copyright
	}

	if ycfg.ExiftoolBinary != nil {
		c.exiftoolBinary = *ycfg.ExiftoolBinary
	}

	if ycfg.FilenamePattern != nil {
		c.filenamePattern = *ycfg.FilenamePattern
	}

	if ycfg.FileSource != nil {
		c.fileSource = ycfg.FileSource
	}

	if ycfg.Make != nil {
		c.make = ycfg.Make
	}

	if ycfg.Model != nil {
		c.model = ycfg.Model
	}

	if ycfg.SerialNumber != nil {
		c.serialNumber = ycfg.SerialNumber
	}

	if ycfg.SetDigitized != nil {
		c.setDigitized = *ycfg.SetDigitized
	}

	if ycfg.TimestampFormat != nil {
		c.timestampFormat = *ycfg.TimestampFormat
	}

	if ycfg.Timezone != nil {
		c.timezone = ycfg.Timezone
	}

	return nil
}

func (c *config) FillFromFlags(f types.Flags) error {
	if f.GetDisplayHelp() {
		c.displayHelp = f.GetDisplayHelp()
	}

	if f.GetDisplayVersion() {
		c.displayVersion = f.GetDisplayVersion()
	}

	if f.GetCopyright() != "" {
		v := f.GetCopyright()
		c.copyright = &v
	}

	if f.GetExiftoolBinary() != "" {
		c.exiftoolBinary = f.GetExiftoolBinary()
	}

	if f.GetFileSource() != "" {
		v := f.GetFileSource()
		c.fileSource = &v
	}

	if f.GetFilenamePattern() != "" {
		c.filenamePattern = f.GetFilenamePattern()
	}

	if f.GetGeotag() != "" {
		v := f.GetGeotag()
		c.geotag = &v
	}

	if f.GetMake() != "" {
		c.make = map[uint8]string{
			0: f.GetMake(),
		}
	}

	if f.GetModel() != "" {
		c.model = map[uint8]string{
			0: f.GetModel(),
		}
	}

	if f.GetSerialNumber() != "" {
		c.serialNumber = map[uint8]string{
			0: f.GetSerialNumber(),
		}
	}

	if f.GetSetDigitized() {
		c.setDigitized = f.GetSetDigitized()
	}

	if f.GetTimestampFormat() != "" {
		c.timestampFormat = f.GetTimestampFormat()
	}

	if f.GetTimezone() != "" {
		c.timezone = map[uint8]string{
			0: f.GetTimezone(),
		}
	}

	return nil
}

func (c *config) GetDisplayHelp() bool {
	return c.displayHelp
}
func (c *config) GetDisplayVersion() bool {
	return c.displayVersion
}

func (c *config) GetCopyright() *string {
	return c.copyright
}

func (c *config) GetExiftoolBinary() string {
	return c.exiftoolBinary
}

func (c *config) GetFilenamePattern() string {
	return c.filenamePattern
}

func (c *config) GetFileSource() *types.FileSource {
	return c.fileSource
}

func (c *config) GetGeotag() *string {
	return c.geotag
}

func (c *config) GetMakeByCameraID(cameraID uint8) *string {
	v := getOrDefault(c.make, cameraID, 0, "")
	if v == "" {
		return nil
	}
	return &v
}

func (c *config) GetModelByCameraID(cameraID uint8) *string {
	v := getOrDefault(c.model, cameraID, 0, "")
	if v == "" {
		return nil
	}
	return &v
}
func (c *config) GetSerialNumberByCameraID(cameraID uint8) *string {
	v := getOrDefault(c.serialNumber, cameraID, 0, "")
	if v == "" {
		return nil
	}
	return &v
}

func (c *config) GetSetDigitized() bool {
	return c.setDigitized
}

func (c *config) GetTimestampFormat() *types.TimestampFormat {
	return &c.timestampFormat
}
func (c *config) GetTimezoneByCameraID(cameraID uint8) types.Timezone {
	return getOrDefault(c.timezone, cameraID, 0, "UTC")
}

func getOrDefault(d map[uint8]string, key, defaultKey uint8, defaultValue string) string {
	if k, ok := d[key]; ok {
		return k
	}

	if dk, ok := d[defaultKey]; ok {
		return dk
	}

	return defaultValue
}
