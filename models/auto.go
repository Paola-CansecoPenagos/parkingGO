package models

const (
	AnchoAuto = 50.0
	AltoAuto  = 50.0
)

type AutoState int

const (
	StateEntering AutoState = iota
	StateParked
	StateExiting
)

type Auto struct {
	PosX       float64
	PosY       float64
	Dir        float64
	Cajon      int
	EnTransito bool
	State      AutoState
}
