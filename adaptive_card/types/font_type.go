package types

import (
	"bytes"
	"encoding/json"
)

type FontType int8

const (
	DefaultFontType FontType = iota
	MonospaceFontType
)

func (FontType FontType) String() string {
	return toFontTypeString[FontType]
}

func (FontType FontType) FromString(keyType string) FontType {
	return toFontTypeID[keyType]
}

var toFontTypeString = map[FontType]string{
	DefaultFontType:   "default",
	MonospaceFontType: "monospace",
}

var toFontTypeID = map[string]FontType{
	"default":   DefaultFontType,
	"monospace": MonospaceFontType,
}

func (FontType FontType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toFontTypeString[FontType])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (FontType *FontType) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*FontType = toFontTypeID[key]
	return nil
}
