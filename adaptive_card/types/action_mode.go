package types

import (
	"bytes"
	"encoding/json"
)

type ActionMode int8

const (
	PrimaryActionMode ActionMode = iota
	SecondaryActionMode
)

func (ActionMode ActionMode) String() string {
	return toActionModeString[ActionMode]
}

func (ActionMode ActionMode) FromString(keyType string) ActionMode {
	return toActionModeID[keyType]
}

var toActionModeString = map[ActionMode]string{
	PrimaryActionMode:   "primary",
	SecondaryActionMode: "secondary",
}

var toActionModeID = map[string]ActionMode{
	"primary":   PrimaryActionMode,
	"secondary": SecondaryActionMode,
}

func (ActionMode ActionMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toActionModeString[ActionMode])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (ActionMode *ActionMode) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*ActionMode = toActionModeID[key]
	return nil
}
