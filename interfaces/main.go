package interfaces

type BodyElement interface {
	New() BodyElement
	ElementType() string
}

type ActionElement interface {
	New() ActionElement
	ActionType() string
}
