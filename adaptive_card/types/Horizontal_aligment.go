package types

import (
	"bytes"
	"encoding/json"
)

type HorizontalAlignment int8

const (
	LeftHorizontalAlignment HorizontalAlignment = iota
	CenterHorizontalAlignment
	RightHorizontalAlignment
)

func (HorizontalAlignment HorizontalAlignment) String() string {
	return toHorizontalAlignmentString[HorizontalAlignment]
}

func (HorizontalAlignment HorizontalAlignment) FromString(keyType string) HorizontalAlignment {
	return toHorizontalAlignmentID[keyType]
}

var toHorizontalAlignmentString = map[HorizontalAlignment]string{
	LeftHorizontalAlignment:   "left",
	CenterHorizontalAlignment: "center",
	RightHorizontalAlignment:  "right",
}

var toHorizontalAlignmentID = map[string]HorizontalAlignment{
	"left":   LeftHorizontalAlignment,
	"center": CenterHorizontalAlignment,
	"right":  RightHorizontalAlignment,
}

func (HorizontalAlignment HorizontalAlignment) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toHorizontalAlignmentString[HorizontalAlignment])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (HorizontalAlignment *HorizontalAlignment) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*HorizontalAlignment = toHorizontalAlignmentID[key]
	return nil
}
