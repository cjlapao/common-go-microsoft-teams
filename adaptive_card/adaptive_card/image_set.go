package adaptive_card

import (
	cryptorand "github.com/cjlapao/common-go-cryptorand"
	"github.com/cjlapao/common-go-microsoft-teams/adaptive_card/types"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type ImageSet struct {
	Type      string           `json:"type"`
	Images    []Image          `json:"images"`
	ImageSize *types.ImageSize `json:"imageSize,omitempty"`

	Height    *types.BlockElementHeight `json:"height,omitempty"`
	Separator *bool                     `json:"separator,omitempty"`
	Spacing   *types.Spacing            `json:"spacing,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	IsVisible *bool                     `json:"isVisible,omitempty"`
}

func (c ImageSet) ElementType() string {
	return "ImageSet"
}

func (c ImageSet) New() interfaces.BodyElement {
	id := cryptorand.GetRandomString(40)
	c.Images = make([]Image, 0)
	c.Type = "ImageSet"
	c.ID = &id

	return c
}

func (c *ImageSet) AddImage(item Image) {
	c.Images = append(c.Images, item)
}

func (c *ImageSet) SetImageSize(value types.ImageSize) {
	c.ImageSize = &value
}

func (c *ImageSet) SetHeight(value types.BlockElementHeight) {
	c.Height = &value
}

func (c *ImageSet) SetSeparator(value bool) {
	c.Separator = &value
}

func (c *ImageSet) SetSpacing(value types.Spacing) {
	c.Spacing = &value
}

func (c *ImageSet) SetVisibility(value bool) {
	c.IsVisible = &value
}
