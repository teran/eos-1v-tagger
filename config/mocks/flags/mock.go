package flags

import (
	"github.com/stretchr/testify/mock"

	"github.com/teran/eos-1v-tagger/types"
)

var (
	_ types.Flags = (*Mock)(nil)
)

// Mock ...
type Mock struct {
	mock.Mock
}

// New ...
func New() *Mock {
	return &Mock{}
}

func (m *Mock) PrintUsageString() {
	_ = m.Called()
	return
}

func (m *Mock) PrintVersionString() {
	_ = m.Called()
	return
}

func (m *Mock) GetDisplayHelp() bool {
	args := m.Called()
	return args.Get(0).(bool)
}

func (m *Mock) GetDisplayVersion() bool {
	args := m.Called()
	return args.Get(0).(bool)
}

func (m *Mock) GetCopyright() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetExiftoolBinary() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetFilenamePattern() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetFileSource() types.FileSource {
	args := m.Called()
	return args.Get(0).(types.FileSource)
}

func (m *Mock) GetGeotag() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetMake() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetModel() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetSerialNumber() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *Mock) GetSetDigitized() bool {
	args := m.Called()
	return args.Get(0).(bool)
}

func (m *Mock) GetTimestampFormat() types.TimestampFormat {
	args := m.Called()
	return args.Get(0).(types.TimestampFormat)
}

func (m *Mock) GetTimezone() types.Timezone {
	args := m.Called()
	return args.Get(0).(types.Timezone)
}

func (m *Mock) GetCSVPath() string {
	args := m.Called()
	return args.Get(0).(string)
}
