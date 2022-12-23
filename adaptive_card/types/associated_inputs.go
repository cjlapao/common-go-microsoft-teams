package types

import (
	"bytes"
	"encoding/json"
)

type AssociatedInputs int8

const (
	AutoAssociatedInputs AssociatedInputs = iota
	NoneAssociatedInputs
)

func (AssociatedInputs AssociatedInputs) String() string {
	return toAssociatedInputsString[AssociatedInputs]
}

func (AssociatedInputs AssociatedInputs) FromString(keyType string) AssociatedInputs {
	return toAssociatedInputsID[keyType]
}

var toAssociatedInputsString = map[AssociatedInputs]string{
	AutoAssociatedInputs: "auto",
	NoneAssociatedInputs: "none",
}

var toAssociatedInputsID = map[string]AssociatedInputs{
	"auto": AutoAssociatedInputs,
	"none": NoneAssociatedInputs,
}

func (AssociatedInputs AssociatedInputs) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toAssociatedInputsString[AssociatedInputs])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (AssociatedInputs *AssociatedInputs) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*AssociatedInputs = toAssociatedInputsID[key]
	return nil
}
