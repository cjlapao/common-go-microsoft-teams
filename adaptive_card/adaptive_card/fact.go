package adaptive_card

import (
	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type Fact struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

func (c Fact) ElementType() string {
	return "Fact"
}

func (c Fact) New() interfaces.BodyElement {
	return c
}
