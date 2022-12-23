package types

import (
	"bytes"
	"encoding/json"
)

type TextBlockStyle int8

const (
	DefaultTextBlockStyle TextBlockStyle = iota
	HeadingTextBlockStyle
)

func (TextBlockStyle TextBlockStyle) String() string {
	return toTextBlockStyleString[TextBlockStyle]
}

func (TextBlockStyle TextBlockStyle) FromString(keyType string) TextBlockStyle {
	return toTextBlockStyleID[keyType]
}

var toTextBlockStyleString = map[TextBlockStyle]string{
	DefaultTextBlockStyle: "default",
	HeadingTextBlockStyle: "heading",
}

var toTextBlockStyleID = map[string]TextBlockStyle{
	"default": DefaultTextBlockStyle,
	"heading": HeadingTextBlockStyle,
}

func (TextBlockStyle TextBlockStyle) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toTextBlockStyleString[TextBlockStyle])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (TextBlockStyle *TextBlockStyle) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*TextBlockStyle = toTextBlockStyleID[key]
	return nil
}
