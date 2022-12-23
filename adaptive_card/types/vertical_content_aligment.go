package types

import (
	"bytes"
	"encoding/json"
)

type VerticalContentAlignment int8

const (
	TopVerticalContentAlignment VerticalContentAlignment = iota
	CenterVerticalContentAlignment
	BottomVerticalContentAlignment
)

func (VerticalContentAlignment VerticalContentAlignment) String() string {
	return toVerticalContentAlignmentString[VerticalContentAlignment]
}

func (VerticalContentAlignment VerticalContentAlignment) FromString(keyType string) VerticalContentAlignment {
	return toVerticalContentAlignmentID[keyType]
}

var toVerticalContentAlignmentString = map[VerticalContentAlignment]string{
	TopVerticalContentAlignment:    "top",
	CenterVerticalContentAlignment: "center",
	BottomVerticalContentAlignment: "bottom",
}

var toVerticalContentAlignmentID = map[string]VerticalContentAlignment{
	"top":    TopVerticalContentAlignment,
	"center": CenterVerticalContentAlignment,
	"bottom": BottomVerticalContentAlignment,
}

func (VerticalContentAlignment VerticalContentAlignment) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toVerticalContentAlignmentString[VerticalContentAlignment])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (VerticalContentAlignment *VerticalContentAlignment) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*VerticalContentAlignment = toVerticalContentAlignmentID[key]
	return nil
}
