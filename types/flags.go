package types

// Flags repository interface
type Flags interface {
	PrintUsageString()
	PrintVersionString()

	GetCSVPath() string

	GetDisplayHelp() bool
	GetDisplayVersion() bool
	GetCopyright() string
	GetExiftoolBinary() string
	GetFilenamePattern() string
	GetFileSource() FileSource
	GetGeotag() string
	GetMake() string
	GetModel() string
	GetSerialNumber() string
	GetSetDigitized() bool
	GetTimestampFormat() TimestampFormat
	GetTimezone() Timezone
}
