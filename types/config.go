package types

// Config repository interface
type Config interface {
	GetDisplayHelp() bool
	GetDisplayVersion() bool
	GetCopyright() *string
	GetExiftoolBinary() string
	GetFilenamePattern() string
	GetFileSource() *FileSource
	GetGeotag() *string
	GetMakeByCameraID(cameraID uint8) *string
	GetModelByCameraID(cameraID uint8) *string
	GetSerialNumberByCameraID(cameraID uint8) *string
	GetSetDigitized() bool
	GetTimestampFormat() *TimestampFormat
	GetTimezoneByCameraID(cameraID uint8) Timezone

	FillFromFlags(f Flags) error
	FillFromYaml(path string) error
}
