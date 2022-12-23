package types

import (
	"bytes"
	"encoding/json"
)

type ContainerStyle int8

const (
	DefaultContainerStyle ContainerStyle = iota
	EmphasisContainerStyle
	GoodContainerStyle
	AttentionContainerStyle
	WarningContainerStyle
	AccentContainerStyle
)

func (ContainerStyle ContainerStyle) String() string {
	return toContainerStyleString[ContainerStyle]
}

func (ContainerStyle ContainerStyle) FromString(keyType string) ContainerStyle {
	return toContainerStyleID[keyType]
}

var toContainerStyleString = map[ContainerStyle]string{
	DefaultContainerStyle:   "default",
	EmphasisContainerStyle:  "emphasis",
	GoodContainerStyle:      "good",
	AttentionContainerStyle: "attention",
	WarningContainerStyle:   "warning",
	AccentContainerStyle:    "accent",
}

var toContainerStyleID = map[string]ContainerStyle{
	"default":   DefaultContainerStyle,
	"emphasis":  EmphasisContainerStyle,
	"good":      GoodContainerStyle,
	"attention": AttentionContainerStyle,
	"warning":   WarningContainerStyle,
	"accent":    AccentContainerStyle,
}

func (ContainerStyle ContainerStyle) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toContainerStyleString[ContainerStyle])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (ContainerStyle *ContainerStyle) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*ContainerStyle = toContainerStyleID[key]
	return nil
}
