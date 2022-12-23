package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type ToggleVisibilityAction struct {
	Type           string          `json:"type"`
	TargetElements []TargetElement `json:"targetElements,omitempty"`

	Title     *string            `json:"title,omitempty"`
	IconUrl   *string            `json:"iconUrl,omitempty"`
	ID        *string            `json:"id,omitempty"`
	Style     *types.ActionStyle `json:"style,omitempty"`
	Tooltip   *string            `json:"tooltip,omitempty"`
	IsEnabled *bool              `json:"isEnabled,omitempty"`
	Mode      *types.ActionMode  `json:"mode,omitempty"`
}

func (c ToggleVisibilityAction) ActionType() string {
	return "Action.ToggleVisibility"
}

func (c ToggleVisibilityAction) New() interfaces.ActionElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ActionType()
	c.TargetElements = make([]TargetElement, 0)
	c.ID = &id

	return c
}

func (c *ToggleVisibilityAction) AddTarget(target TargetElement) {
	found := false
	for _, element := range c.TargetElements {
		if element.ElementId == target.ElementId {
			element.IsVisible = target.IsVisible
			found = true
		}
	}
	if !found {
		c.TargetElements = append(c.TargetElements, target)
	}
}

func (c *ToggleVisibilityAction) SetTitle(title string) {
	c.Title = &title
}

func (c *ToggleVisibilityAction) SetIconUrl(iconUrl string) {
	c.IconUrl = &iconUrl
}

func (c *ToggleVisibilityAction) SetStyle(style types.ActionStyle) {
	c.Style = &style
}

func (c *ToggleVisibilityAction) SetTooltip(tooltip string) {
	c.Tooltip = &tooltip
}

func (c *ToggleVisibilityAction) SetIsEnabled(value bool) {
	c.IsEnabled = &value
}

func (c *ToggleVisibilityAction) SetMode(value types.ActionMode) {
	c.Mode = &value
}
