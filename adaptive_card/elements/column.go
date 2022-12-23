package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type Column struct {
	Type                     string                          `json:"type"`
	Items                    []interfaces.BodyElement        `json:"items"`
	BackgroundImage          *string                         `json:"backgroundImage,omitempty"`
	Bleed                    *bool                           `json:"bleed,omitempty"`
	MinHeight                *string                         `json:"minHeight,omitempty"`
	RightToLeft              *bool                           `json:"rtl,omitempty"`
	Separator                *bool                           `json:"separator,omitempty"`
	Spacing                  *types.Spacing                  `json:"spacing,omitempty"`
	Action                   *interfaces.ActionElement       `json:"selectAction,omitempty"`
	Style                    *types.ContainerStyle           `json:"style,omitempty"`
	VerticalContentAlignment *types.VerticalContentAlignment `json:"VerticalContentAlignment,omitempty"`
	Width                    *string                         `json:"width,omitempty"`

	ID        *string `json:"id,omitempty"`
	IsVisible *bool   `json:"isVisible,omitempty"`
}

func (c Column) ElementType() string {
	return "Column"
}

func (c Column) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Items = make([]interfaces.BodyElement, 0)
	c.Type = "Column"
	c.ID = &id

	return c
}

func (c *Column) AddItem(item interfaces.BodyElement) {
	c.Items = append(c.Items, item)
}

func (c *Column) SetBackgroundImage(value string) {
	c.BackgroundImage = &value
}

func (c *Column) SetBleed(value bool) {
	c.Bleed = &value
}

func (c *Column) SetMinHeight(value string) {
	c.MinHeight = &value
}

func (c *Column) SetRightToLeft(value bool) {
	c.RightToLeft = &value
}

func (c *Column) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *Column) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *Column) SetAction(value interfaces.ActionElement) {
	c.Action = &value
}

func (c *Column) SetStyle(value types.ContainerStyle) {
	c.Style = &value
}

func (c *Column) SetVerticalContentAlignment(value types.VerticalContentAlignment) {
	c.VerticalContentAlignment = &value
}

func (c *Column) SetWidth(value string) {
	c.MinHeight = &value
}

func (c *Column) SetVisibility(value bool) {
	c.IsVisible = &value
}
