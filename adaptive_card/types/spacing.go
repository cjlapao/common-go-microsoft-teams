package types

import (
	"bytes"
	"encoding/json"
)

type Spacing int8

const (
	DefaultSpacing Spacing = iota
	NoneSpacing
	SmallSpacing
	MediumSpacing
	LargeSpacing
	ExtraLargeSpacing
	PaddingSpacing
)

func (Spacing Spacing) String() string {
	return toSpacingString[Spacing]
}

func (Spacing Spacing) FromString(keyType string) Spacing {
	return toSpacingID[keyType]
}

var toSpacingString = map[Spacing]string{
	DefaultSpacing:    "default",
	NoneSpacing:       "none",
	SmallSpacing:      "small",
	MediumSpacing:     "medium",
	LargeSpacing:      "large",
	ExtraLargeSpacing: "extraLarge",
	PaddingSpacing:    "padding",
}

var toSpacingID = map[string]Spacing{
	"default":    DefaultSpacing,
	"none":       NoneSpacing,
	"small":      SmallSpacing,
	"medium":     MediumSpacing,
	"large":      LargeSpacing,
	"extraLarge": ExtraLargeSpacing,
	"padding":    PaddingSpacing,
}

func (Spacing Spacing) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toSpacingString[Spacing])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (Spacing *Spacing) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*Spacing = toSpacingID[key]
	return nil
}
