package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type OpenUrlAction struct {
	Type string `json:"type"`
	Url  string `json:"url"`

	Title     *string            `json:"title,omitempty"`
	IconUrl   *string            `json:"iconUrl,omitempty"`
	ID        *string            `json:"id,omitempty"`
	Style     *types.ActionStyle `json:"style,omitempty"`
	Tooltip   *string            `json:"tooltip,omitempty"`
	IsEnabled *bool              `json:"isEnabled,omitempty"`
	Mode      *types.ActionMode  `json:"mode,omitempty"`
}

func (c OpenUrlAction) ActionType() string {
	return "Action.OpenUrl"
}

func (c OpenUrlAction) New() interfaces.ActionElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ActionType()
	c.ID = &id

	return c
}

func (c *OpenUrlAction) SetTitle(title string) {
	c.Title = &title
}

func (c *OpenUrlAction) SetIconUrl(iconUrl string) {
	c.IconUrl = &iconUrl
}

func (c *OpenUrlAction) SetStyle(style types.ActionStyle) {
	c.Style = &style
}

func (c *OpenUrlAction) SetTooltip(tooltip string) {
	c.Tooltip = &tooltip
}

func (c *OpenUrlAction) SetIsEnabled(value bool) {
	c.IsEnabled = &value
}

func (c *OpenUrlAction) SetMode(value types.ActionMode) {
	c.Mode = &value
}
