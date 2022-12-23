package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type RichTextBlock struct {
	Type                string                     `json:"type"`
	InLines             []TextRun                  `json:"inlines"`
	HorizontalAlignment *types.HorizontalAlignment `json:"horizontalAlignment,omitempty"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c RichTextBlock) ElementType() string {
	return "RichTextBlock"
}

func (c RichTextBlock) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.InLines = make([]TextRun, 0)
	c.Type = "RichTextBlock"
	c.ID = &id

	return c
}

func (c *RichTextBlock) AddLine(item TextRun) {
	c.InLines = append(c.InLines, item)
}

func (c *RichTextBlock) SetHorizontalAlignment(value types.HorizontalAlignment) {
	c.HorizontalAlignment = &value
}

func (c *RichTextBlock) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *RichTextBlock) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *RichTextBlock) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *RichTextBlock) SetVisibility(value bool) {
	c.IsVisible = &value
}
