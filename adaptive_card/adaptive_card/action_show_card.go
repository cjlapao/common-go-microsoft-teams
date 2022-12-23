package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type ShowCardAction struct {
	Type string       `json:"type"`
	Card AdaptiveCard `json:"card,omitempty"`

	Title     *string            `json:"title,omitempty"`
	IconUrl   *string            `json:"iconUrl,omitempty"`
	ID        *string            `json:"id,omitempty"`
	Style     *types.ActionStyle `json:"style,omitempty"`
	Tooltip   *string            `json:"tooltip,omitempty"`
	IsEnabled *bool              `json:"isEnabled,omitempty"`
	Mode      *types.ActionMode  `json:"mode,omitempty"`
}

func (c ShowCardAction) ActionType() string {
	return "Action.ShowCard"
}

func (c ShowCardAction) New() interfaces.ActionElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ActionType()
	c.ID = &id

	return c
}

func (c *ShowCardAction) SetTitle(title string) {
	c.Title = &title
}

func (c *ShowCardAction) SetIconUrl(iconUrl string) {
	c.IconUrl = &iconUrl
}

func (c *ShowCardAction) SetStyle(style types.ActionStyle) {
	c.Style = &style
}

func (c *ShowCardAction) SetTooltip(tooltip string) {
	c.Tooltip = &tooltip
}

func (c *ShowCardAction) SetIsEnabled(value bool) {
	c.IsEnabled = &value
}

func (c *ShowCardAction) SetMode(value types.ActionMode) {
	c.Mode = &value
}
