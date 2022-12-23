package types

import (
	"bytes"
	"encoding/json"
)

type Colors int8

const (
	DefaultColor Colors = iota
	DarkColor
	LightColor
	AccentColor
	GoodColor
	WarningColor
	AttentionColor
)

func (Colors Colors) String() string {
	return toColorsString[Colors]
}

func (Colors Colors) FromString(keyType string) Colors {
	return toColorsID[keyType]
}

var toColorsString = map[Colors]string{
	DefaultColor:   "default",
	DarkColor:      "dark",
	LightColor:     "light",
	AccentColor:    "accent",
	GoodColor:      "good",
	WarningColor:   "warning",
	AttentionColor: "attention",
}

var toColorsID = map[string]Colors{
	"default":   DefaultColor,
	"dark":      DarkColor,
	"light":     LightColor,
	"accent":    AccentColor,
	"good":      GoodColor,
	"warning":   WarningColor,
	"attention": AttentionColor,
}

func (Colors Colors) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toColorsString[Colors])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (Colors *Colors) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*Colors = toColorsID[key]
	return nil
}
