package types

import (
	"bytes"
	"encoding/json"
)

type ActionStyle int8

const (
	DefaultActionStyle ActionStyle = iota
	PositiveActionStyle
	DestructiveActionStyle
)

func (ActionStyle ActionStyle) String() string {
	return toActionStyleString[ActionStyle]
}

func (ActionStyle ActionStyle) FromString(keyType string) ActionStyle {
	return toActionStyleID[keyType]
}

var toActionStyleString = map[ActionStyle]string{
	DefaultActionStyle:     "default",
	PositiveActionStyle:    "positive",
	DestructiveActionStyle: "destructive",
}

var toActionStyleID = map[string]ActionStyle{
	"default":     DefaultActionStyle,
	"positive":    PositiveActionStyle,
	"destructive": DestructiveActionStyle,
}

func (ActionStyle ActionStyle) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toActionStyleString[ActionStyle])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (ActionStyle *ActionStyle) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*ActionStyle = toActionStyleID[key]
	return nil
}
