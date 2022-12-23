package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type ExecuteAction struct {
	Type             string                  `json:"type"`
	Verb             *string                 `json:"verb,omitempty"`
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

func (c ExecuteAction) ActionType() string {
	return "Action.Submit"
}

func (c ExecuteAction) New() interfaces.ActionElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ActionType()
	c.ID = &id

	return c
}

func (c *ExecuteAction) SetVerb(value string) {
	c.Verb = &value
}

func (c *ExecuteAction) AddData(key string, value interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}

	c.Data[key] = value
}

func (c *ExecuteAction) SetAssociatedInputs(value types.AssociatedInputs) {
	c.AssociatedInputs = &value
}

func (c *ExecuteAction) SetTitle(title string) {
	c.Title = &title
}

func (c *ExecuteAction) SetIconUrl(iconUrl string) {
	c.IconUrl = &iconUrl
}

func (c *ExecuteAction) SetStyle(style types.ActionStyle) {
	c.Style = &style
}

func (c *ExecuteAction) SetTooltip(tooltip string) {
	c.Tooltip = &tooltip
}

func (c *ExecuteAction) SetIsEnabled(value bool) {
	c.IsEnabled = &value
}

func (c *ExecuteAction) SetMode(value types.ActionMode) {
	c.Mode = &value
}
