package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type Container struct {
	Type                     string                          `json:"type"`
	Items                    []interfaces.BodyElement        `json:"items"`
	Action                   *interfaces.ActionElement       `json:"selectAction,omitempty"`
	Style                    *types.ContainerStyle           `json:"style,omitempty"`
	VerticalContentAlignment *types.VerticalContentAlignment `json:"VerticalContentAlignment,omitempty"`
	Bleed                    *bool                           `json:"bleed,omitempty"`
	BackgroundImage          *string                         `json:"backgroundImage,omitempty"`
	MinHeight                *string                         `json:"minHeight,omitempty"`
	RightToLeft              *bool                           `json:"rtl,omitempty"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c Container) ElementType() string {
	return "Container"
}

func (c Container) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Items = make([]interfaces.BodyElement, 0)
	c.Type = "Container"
	c.ID = &id

	return c
}

func (c *Container) AddItem(item interfaces.BodyElement) {
	c.Items = append(c.Items, item)
}

func (c *Container) SetAction(value interfaces.ActionElement) {
	c.Action = &value
}

func (c *Container) SetStyle(value types.ContainerStyle) {
	c.Style = &value
}

func (c *Container) SetVerticalContentAlignment(value types.VerticalContentAlignment) {
	c.VerticalContentAlignment = &value
}

func (c *Container) SetBleed(value bool) {
	c.Bleed = &value
}

func (c *Container) SetBackgroundImage(value string) {
	c.BackgroundImage = &value
}

func (c *Container) SetMinHeight(value string) {
	c.MinHeight = &value
}

func (c *Container) SetRightToLeft(value bool) {
	c.RightToLeft = &value
}

func (c *Container) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *Container) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *Container) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *Container) SetVisibility(value bool) {
	c.IsVisible = &value
}
