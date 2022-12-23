package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type TextBlock struct {
	Type                string                     `json:"type"`
	Text                string                     `json:"text"`
	Color               *types.Colors              `json:"color,omitempty"`
	FontType            *types.FontType            `json:"fontType,omitempty"`
	HorizontalAlignment *types.HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	IsSubtle            *bool                      `json:"isSubtle,omitempty"`
	MaxLines            *int                       `json:"maxLines,omitempty"`
	Size                *types.FontSize            `json:"size,omitempty"`
	Weight              *types.FontWeight          `json:"weight,omitempty"`
	Wrap                *bool                      `json:"wrap,omitempty"`
	Style               *types.TextBlockStyle      `json:"style,omitempty"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c TextBlock) ElementType() string {
	return "TextBlock"
}

func (c TextBlock) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Type = c.ElementType()
	c.ID = &id

	return c
}

func (c *TextBlock) SetColor(value types.Colors) {
	c.Color = &value
}

func (c *TextBlock) SetFontType(value types.FontType) {
	c.FontType = &value
}

func (c *TextBlock) SetHorizontalAlignment(value types.HorizontalAlignment) {
	c.HorizontalAlignment = &value
}

func (c *TextBlock) SetIsSubtle(value bool) {
	c.IsSubtle = &value
}

func (c *TextBlock) SetMaxLines(value int) {
	c.MaxLines = &value
}

func (c *TextBlock) SetSize(value types.FontSize) {
	c.Size = &value
}

func (c *TextBlock) SetWeight(value types.FontWeight) {
	c.Weight = &value
}

func (c *TextBlock) SetWrap(value bool) {
	c.Wrap = &value
}

func (c *TextBlock) SetStyle(value types.TextBlockStyle) {
	c.Style = &value
}

func (c *TextBlock) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *TextBlock) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *TextBlock) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *TextBlock) SetVisibility(value bool) {
	c.IsVisible = &value
}
