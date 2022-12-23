package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type FactSet struct {
	Type  string `json:"type"`
	Facts []Fact `json:"facts"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c FactSet) ElementType() string {
	return "FactSet"
}

func (c FactSet) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Facts = make([]Fact, 0)
	c.Type = "FactSet"
	c.ID = &id

	return c
}

func (c *FactSet) AddFact(item Fact) {
	c.Facts = append(c.Facts, item)
}

func (c *FactSet) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *FactSet) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *FactSet) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *FactSet) SetVisibility(value bool) {
	c.IsVisible = &value
}
