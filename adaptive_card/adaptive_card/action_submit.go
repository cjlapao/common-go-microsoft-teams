package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type SubmitAction struct {
	Type             string                  `json:"type"`
	Data             map[string]interface{}  `json:"data,omitempty"`
	AssociatedInputs *types.AssociatedInputs `json:"associatedInputs,omitempty"`

	Title     *string            `json:"title,omitempty"`
	IconUrl   *string            `json:"iconUrl,omitempty"`
	ID        *string            `json:"id,omitempty"`
	Style     *types.ActionStyle `json:"style,omitempty"`
	Tooltip   *string            `json:"tooltip,omitempty"`
	IsEnabled *bool              `json:"isEnabled,omitempty"`
	Mode      *types.ActionMode  `json:"mode,omitempty"`
}

func (c SubmitAction) ActionType() string {
	return "Action.Submit"
}

func (c SubmitAction) New() interfaces.ActionElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ActionType()
	c.ID = &id

	return c
}

func (c *SubmitAction) AddData(key string, value interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}

	c.Data[key] = value
}

func (c *SubmitAction) SetAssociatedInputs(value types.AssociatedInputs) {
	c.AssociatedInputs = &value
}

func (c *SubmitAction) SetTitle(title string) {
	c.Title = &title
}

func (c *SubmitAction) SetIconUrl(iconUrl string) {
	c.IconUrl = &iconUrl
}

func (c *SubmitAction) SetStyle(style types.ActionStyle) {
	c.Style = &style
}

func (c *SubmitAction) SetTooltip(tooltip string) {
	c.Tooltip = &tooltip
}

func (c *SubmitAction) SetIsEnabled(value bool) {
	c.IsEnabled = &value
}

func (c *SubmitAction) SetMode(value types.ActionMode) {
	c.Mode = &value
}
