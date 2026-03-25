package model

const (
	NoWinner = -1
	Draw     = -2
)

type Enum interface {
	Int() int
	String() string
}

// TODO: Maybe change indexing of players so default FigureType values represent no player also

// Implementation of FigureType should be defined something like:
//
//	type GameNameType int
//
// Possible values of the type should be defined as:
//
//	const (
//		GameNameNoFigure GameNameType = iota
//		GameNameSomeFigure
//		GameNameAnotherFigure
//	)
//
// Zero value of FigureType should represent the absence of figure
type FigureType interface {
	Enum
}

type Figure[T FigureType] struct {
	Player int
	Type   T
}
