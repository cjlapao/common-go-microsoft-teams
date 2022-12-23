package types

import (
	"bytes"
	"encoding/json"
)

type FontSize int8

const (
	DefaultFontSize FontSize = iota
	SmallFontSize
	MediumFontSize
	LargeFontSize
	ExtraLargeFontSize
)

func (FontSize FontSize) String() string {
	return toFontSizeString[FontSize]
}

func (FontSize FontSize) FromString(keyType string) FontSize {
	return toFontSizeID[keyType]
}

var toFontSizeString = map[FontSize]string{
	DefaultFontSize:    "default",
	SmallFontSize:      "small",
	MediumFontSize:     "medium",
	LargeFontSize:      "large",
	ExtraLargeFontSize: "extraLarge",
}

var toFontSizeID = map[string]FontSize{
	"default":    DefaultFontSize,
	"small":      SmallFontSize,
	"medium":     MediumFontSize,
	"large":      LargeFontSize,
	"extraLarge": ExtraLargeFontSize,
}

func (FontSize FontSize) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toFontSizeString[FontSize])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (FontSize *FontSize) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*FontSize = toFontSizeID[key]
	return nil
}
