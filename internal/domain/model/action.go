package model

type ActionType int

const (
	ActionGet ActionType = iota
	ActionAdd
	ActionRemove
)

func (a ActionType) Int() int {
	return int(a)
}

func (a ActionType) String() string {
	return [...]string{
		"Get",
		"Add",
		"Remove",
	}[a.Int()]
}

type Action[T FigureType] struct {
	Type   ActionType
	Figure Figure[T]
	Args   []int
}
