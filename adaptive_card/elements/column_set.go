package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type ColumnSet struct {
	ID                  *string                    `json:"id,omitempty"`
	Type                string                     `json:"type"`
	Columns             []Column                   `json:"columns"`
	Action              *interfaces.ActionElement  `json:"selectAction,omitempty"`
	Style               *types.ContainerStyle      `json:"style,omitempty"`
	Bleed               *bool                      `json:"bleed,omitempty"`
	MinHeight           *string                    `json:"minHeight,omitempty"`
	HorizontalAlignment *types.HorizontalAlignment `json:"horizontalAlignment,omitempty"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c ColumnSet) ElementType() string {
	return "ColumnSet"
}

func (c ColumnSet) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Columns = make([]Column, 0)
	c.Type = "ColumnSet"
	c.ID = &id

	return c
}

func (c *ColumnSet) SetAction(value interfaces.ActionElement) {
	c.Action = &value
}

func (c *ColumnSet) SetStyle(value types.ContainerStyle) {
	c.Style = &value
}

func (c *ColumnSet) SetBleed(value bool) {
	c.Bleed = &value
}

func (c *ColumnSet) SetMinHeight(value string) {
	c.MinHeight = &value
}

func (c *ColumnSet) SetHorizontalAlignment(value types.HorizontalAlignment) {
	c.HorizontalAlignment = &value
}

func (c *ColumnSet) AddColumn(item Column) {
	c.Columns = append(c.Columns, item)
}

func (c *ColumnSet) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *ColumnSet) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *ColumnSet) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *ColumnSet) SetVisibility(value bool) {
	c.IsVisible = &value
}
