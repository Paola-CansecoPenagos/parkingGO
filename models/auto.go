package models

import "sync"

const (
	AnchoAuto = 50.0
	AltoAuto  = 50.0
)

type Auto struct {
	PosX  float64
	PosY  float64
	Dir   float64
	Cajon int
}

type Estacionamiento struct {
	Espacios int
	Mu       sync.Mutex
	Ocupados []*Auto
	EnEspera []*Auto
	Cajones  []bool
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
	return &Estacionamiento{
		Espacios: capacidad,
		Ocupados: make([]*Auto, capacidad),
		Cajones:  make([]bool, capacidad),
	}
}

func (e *Estacionamiento) Entrar(auto *Auto) int {
	e.Mu.Lock()
	defer e.Mu.Unlock()
	for i, lugar := range e.Ocupados {
		if lugar == nil {
			e.Ocupados[i] = auto
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
	e.EnEspera = append(e.EnEspera, auto)
	return -1
}

func (e *Estacionamiento) Salir(i int) {
	e.Mu.Lock()
	defer e.Mu.Unlock()
	if e.Ocupados[i] != nil {
		e.Cajones[e.Ocupados[i].Cajon] = false
	}
	e.Ocupados[i] = nil
	if len(e.EnEspera) > 0 {
		e.Ocupados[i] = e.EnEspera[0]
		e.EnEspera = e.EnEspera[1:]
	}
}
