package types

import (
	"bytes"
	"encoding/json"
)

type ImageSize int8

const (
	AutoImageSize ImageSize = iota
	SmallImageSize
	MediumImageSize
	LargeImageSize
	StretchImageSize
)

func (ImageSize ImageSize) String() string {
	return toImageSizeString[ImageSize]
}

func (ImageSize ImageSize) FromString(keyType string) ImageSize {
	return toImageSizeID[keyType]
}

var toImageSizeString = map[ImageSize]string{
	AutoImageSize:    "auto",
	SmallImageSize:   "small",
	MediumImageSize:  "medium",
	LargeImageSize:   "large",
	StretchImageSize: "stretch",
}

var toImageSizeID = map[string]ImageSize{
	"auto":    AutoImageSize,
	"small":   SmallImageSize,
	"medium":  MediumImageSize,
	"large":   LargeImageSize,
	"stretch": StretchImageSize,
}

func (ImageSize ImageSize) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toImageSizeString[ImageSize])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (ImageSize *ImageSize) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*ImageSize = toImageSizeID[key]
	return nil
}
