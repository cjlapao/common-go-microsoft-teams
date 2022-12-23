package types

import (
	"bytes"
	"encoding/json"
)

type ImageStyle int8

const (
	DefaultImageStyle ImageStyle = iota
	PersonImageStyle
)

func (ImageStyle ImageStyle) String() string {
	return toImageStyleString[ImageStyle]
}

func (ImageStyle ImageStyle) FromString(keyType string) ImageStyle {
	return toImageStyleID[keyType]
}

var toImageStyleString = map[ImageStyle]string{
	DefaultImageStyle: "default",
	PersonImageStyle:  "person",
}

var toImageStyleID = map[string]ImageStyle{
	"default": DefaultImageStyle,
	"person":  PersonImageStyle,
}

func (ImageStyle ImageStyle) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toImageStyleString[ImageStyle])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (ImageStyle *ImageStyle) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*ImageStyle = toImageStyleID[key]
	return nil
}
