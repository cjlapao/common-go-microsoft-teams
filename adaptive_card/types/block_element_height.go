package types

import (
	"bytes"
	"encoding/json"
)

type BlockElementHeight int8

const (
	AutoBlockElementHeight BlockElementHeight = iota
	StretchBlockElementHeight
)

func (BlockElementHeight BlockElementHeight) String() string {
	return toBlockElementHeightString[BlockElementHeight]
}

func (BlockElementHeight BlockElementHeight) FromString(keyType string) BlockElementHeight {
	return toBlockElementHeightID[keyType]
}

var toBlockElementHeightString = map[BlockElementHeight]string{
	AutoBlockElementHeight:    "auto",
	StretchBlockElementHeight: "stretch",
}

var toBlockElementHeightID = map[string]BlockElementHeight{
	"auto":    AutoBlockElementHeight,
	"stretch": StretchBlockElementHeight,
}

func (BlockElementHeight BlockElementHeight) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toBlockElementHeightString[BlockElementHeight])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (BlockElementHeight *BlockElementHeight) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*BlockElementHeight = toBlockElementHeightID[key]
	return nil
}
