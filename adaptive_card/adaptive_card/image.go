package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type Image struct {
	Type                string                     `json:"type"`
	Url                 string                     `json:"url"`
	AlternativeText     *string                    `json:"altText,omitempty"`
	BackgroundColor     *string                    `json:"backgroundColor,omitempty"`
	Height              *types.BlockElementHeight  `json:"height,omitempty"`
	HorizontalAlignment *types.HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	Action              *interfaces.ActionElement  `json:"selectAction"`
	Size                *types.ImageSize           `json:"size,omitempty"`
	Style               *types.ImageStyle          `json:"style,omitempty"`
	Width               *string                    `json:"width,omitempty"`

	Separator *bool          `json:"separator,omitempty"`
	Spacing   *types.Spacing `json:"spacing,omitempty"`
	ID        *string        `json:"id,omitempty"`
	IsVisible *bool          `json:"isVisible,omitempty"`
}

func (c Image) ElementType() string {
	return "Image"
}

func (c Image) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ElementType()
	c.ID = &id

	return c
}

func (c *Image) SetAlternativeText(value string) {
	c.AlternativeText = &value
}

func (c *Image) SetBackgroundColor(value string) {
	c.BackgroundColor = &value
}

func (c *Image) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *Image) SetHorizontalAlignment(value types.HorizontalAlignment) {
	c.HorizontalAlignment = &value
}

func (c *Image) SetAction(value interfaces.ActionElement) {
	c.Action = &value
}

func (c *Image) SetSize(value types.ImageSize) {
	c.Size = &value
}

func (c *Image) SetStyle(value types.ImageStyle) {
	c.Style = &value
}

func (c *Image) SetWidth(value string) {
	c.Width = &value
}

func (c *Image) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *Image) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *Image) SetVisibility(value bool) {
	c.IsVisible = &value
}
