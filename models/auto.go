package models

const (
	StateEntering       AutoState = iota
	StateParked
	StateExiting
)

type AutoState int

type Auto struct {
	PosX       float64
	PosY       float64
	Dir        float64
	Cajon      int
	EnTransito bool
	State      AutoState
}
