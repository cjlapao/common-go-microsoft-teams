package adaptive_card

import (
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type TextRun struct {
	Type          string                    `json:"type"`
	Text          string                    `json:"text"`
	Color         *types.Colors             `json:"color,omitempty"`
	FontType      *types.FontType           `json:"fontType,omitempty"`
	Highlight     *bool                     `json:"highlight,omitempty"`
	IsSubtle      *bool                     `json:"isSubtle,omitempty"`
	Italic        *bool                     `json:"italic,omitempty"`
	Action        *interfaces.ActionElement `json:"selectAction,omitempty"`
	Size          *types.FontSize           `json:"size,omitempty"`
	StrikeThrough *bool                     `json:"strikethrough,omitempty"`
	Underline     *bool                     `json:"underline,omitempty"`
	Weight        *types.FontWeight         `json:"weight,omitempty"`
}

func (c TextRun) ElementType() string {
	return "TextRun"
}

func (c TextRun) New() interfaces.BodyElement {
	c.Type = c.ElementType()

	return c
}

func (c *TextRun) SetColor(value types.Colors) {
	c.Color = &value
}

func (c *TextRun) SetFontType(value types.FontType) {
	c.FontType = &value
}

func (c *TextRun) SetHighlight(value bool) {
	c.Highlight = &value
}

func (c *TextRun) SetIsSubtle(value bool) {
	c.IsSubtle = &value
}

func (c *TextRun) SetItalic(value bool) {
	c.Italic = &value
}

func (c *TextRun) SetAction(value interfaces.ActionElement) {
	c.Action = &value
}

func (c *TextRun) SetSize(value types.FontSize) {
	c.Size = &value
}

func (c *TextRun) SetStrikeThrough(value bool) {
	c.StrikeThrough = &value
}

func (c *TextRun) SetUnderline(value bool) {
	c.Underline = &value
}

func (c *TextRun) SetWeight(value types.FontWeight) {
	c.Weight = &value
}
