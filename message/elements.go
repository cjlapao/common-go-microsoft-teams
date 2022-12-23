package message

import "github.com/cjlapao/common-go-microsoft-teams/interfaces"

func NewElement[T interfaces.BodyElement]() T {
	t := *new(T)
	r := t.New()

	return r.(T)
}

func NewAction[T interfaces.ActionElement]() T {
	t := *new(T)
	r := t.New()

	return r.(T)
}

type BodyElements struct {
	Elements []interfaces.BodyElement
}

func NewBodyElements() BodyElements {
	body := BodyElements{
		Elements: make([]interfaces.BodyElement, 0),
	}

	return body
}

func (e *BodyElements) AddElement(item interfaces.BodyElement) {
	e.Elements = append(e.Elements, item)
}

func (e *BodyElements) Get() []interfaces.BodyElement {
	return e.Elements
}
