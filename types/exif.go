package types

// EXIFValue ...
type EXIFValue map[string]string

// EXIFValuer ...
type EXIFValuer interface {
	EXIFValue() EXIFValue
}
