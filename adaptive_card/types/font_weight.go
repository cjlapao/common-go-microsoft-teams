package types

import (
	"bytes"
	"encoding/json"
)

type FontWeight int8

const (
	DefaultFontWeight FontWeight = iota
	LighterFontWeight
	BolderFontWeight
)

func (FontWeight FontWeight) String() string {
	return toFontWeightString[FontWeight]
}

func (FontWeight FontWeight) FromString(keyType string) FontWeight {
	return toFontWeightID[keyType]
}

var toFontWeightString = map[FontWeight]string{
	DefaultFontWeight: "default",
	LighterFontWeight: "lighter",
	BolderFontWeight:  "bolder",
}

var toFontWeightID = map[string]FontWeight{
	"default": DefaultFontWeight,
	"lighter": LighterFontWeight,
	"bolder":  BolderFontWeight,
}

func (FontWeight FontWeight) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toFontWeightString[FontWeight])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (FontWeight *FontWeight) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*FontWeight = toFontWeightID[key]
	return nil
}
