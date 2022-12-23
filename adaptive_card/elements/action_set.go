package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type ActionSet struct {
	Type    string                     `json:"type"`
	Actions []interfaces.ActionElement `json:"actions"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c ActionSet) ElementType() string {
	return "ActionSet"
}

func (c ActionSet) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Actions = make([]interfaces.ActionElement, 0)
	c.Type = "ActionSet"
	c.ID = &id

	return c
}

func (c *ActionSet) AddAction(item interfaces.ActionElement) {
	c.Actions = append(c.Actions, item)
}

func (c *ActionSet) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *ActionSet) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *ActionSet) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *ActionSet) SetVisibility(value bool) {
	c.IsVisible = &value
}
