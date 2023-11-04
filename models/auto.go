package models

import (
	"sync"
)

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

type Parking struct {
	Espacios int
	Mu       sync.Mutex
	Ocupados []*Auto
	EnEspera []*Auto
	Cajones  []bool
}

func NewParking(capacidad int) *Parking {
	return &Parking{
		Espacios: capacidad,
		Ocupados: make([]*Auto, capacidad),
		Cajones:  make([]bool, capacidad),
	}
}

func (e *Parking) Enter(auto *Auto) int {
	e.Mu.Lock()
	defer e.Mu.Unlock()
	for i, lugar := range e.Ocupados {
		if lugar == nil {
			e.Ocupados[i] = auto
			auto.EnTransito = true
			for j := 0; j < e.Espacios; j++ {
				if !e.Cajones[j] {
					e.Cajones[j] = true
					auto.Cajon = j
					return i
				}
			}
			return -1
		}
	}
	return -1
}
