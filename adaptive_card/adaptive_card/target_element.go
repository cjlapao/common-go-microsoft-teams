package adaptive_card

import (
	"encoding/json"
	"strconv"
)

type TargetElement struct {
	ElementId string `json:"elementId"`
	IsVisible *bool  `json:"isVisible,omitempty"`
}

func (c TargetElement) Toggle(elementId string) TargetElement {
	c.ElementId = elementId
	return c
}

func (c TargetElement) Show(elementId string) TargetElement {
	value := true
	c.ElementId = elementId
	c.IsVisible = &value
	return c
}

func (c TargetElement) Hide(elementId string) TargetElement {
	value := false
	c.ElementId = elementId
	c.IsVisible = &value
	return c
}

func (TargetElement TargetElement) MarshalJSON() ([]byte, error) {
	if TargetElement.IsVisible != nil {
		value := `{
  "elementId": "` + TargetElement.ElementId + `",
  "isVisible": ` + strconv.FormatBool(*TargetElement.IsVisible) + `
}`
		return []byte(value), nil
	} else {
		value := `"` + TargetElement.ElementId + `"`
		return []byte(value), nil
	}
}

func (t *TargetElement) UnmarshalJSON(b []byte) error {
	var key string
	objByte := make([]byte, len(b))
	copy(objByte, b)
	stringErr := json.Unmarshal(b, &key)
	if stringErr == nil {
		t.ElementId = key
		return nil
	}

	var target map[string]interface{}
	objErr := json.Unmarshal(objByte, &target)

	if objErr != nil {
		return objErr
	}

	id := target["elementId"].(string)
	visible := target["isVisible"].(bool)
	t.ElementId = id
	t.IsVisible = &visible

	return nil
}
