package adaptive_card

import (
	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type AdaptiveCard struct {
	Type string                   `json:"type"`
	Body []interfaces.BodyElement `json:"body,omitempty"`
}

func (c AdaptiveCard) New() AdaptiveCard {
	c.Type = "AdaptiveCard"
	c.Body = make([]interfaces.BodyElement, 0)
	return c
}

func (c *AdaptiveCard) AddElement(item interfaces.BodyElement) {
	c.Body = append(c.Body, item)
}
