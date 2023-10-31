package views

import (
	"math/rand"
	"time"

	"parking/models"
	"parking/scenes"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Velocidad           = 0.2
	AnchoAuto           = 80.0
	AltoAuto            = 80.0
	AltoEspacio         = 100.0
	DistanciaEntreAutos = 10.0
)

func GenerarVehiculos(e *models.Estacionamiento) {
	rand.Seed(time.Now().UnixNano())

	for {
		auto := &models.Auto{PosX: -models.AnchoAuto - DistanciaEntreAutos, PosY: scenes.AltoVentana/2 + models.AltoAuto/2, Dir: 1}
		pos := e.Entrar(auto)

		if pos != -1 {
			go func(p int) {
				time.Sleep(time.Duration(5) * time.Second)
				e.Salir(p)
			}(pos)
		}
		time.Sleep(time.Millisecond * 1500)
	}
}

func Run(win *pixelgl.Window, e *models.Estacionamiento) {
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			auto := &models.Auto{PosX: -AnchoAuto - DistanciaEntreAutos, PosY: scenes.AltoVentana/2 + AltoAuto/2, Dir: 1}
			pos := e.Entrar(auto)

			if pos != -1 {
				go func(p int) {
					time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)
					e.Salir(p)
				}(pos)
			}
			time.Sleep(time.Millisecond * 1500)
		}
	}()

	for !win.Closed() {
		win.Clear(pixel.RGB(1, 1, 1))

		im := imdraw.New(nil)
		scenes.DibujarEstacionamiento(im, e)

		e.Mu.Lock()
		for i, auto := range e.Ocupados {
			if auto != nil {
				if i%2 == 0 {
					auto.PosY = scenes.AltoVentana/2 + AltoAuto/2
				} else {
					auto.PosY = scenes.AltoVentana/2 - AltoAuto/2 - AltoEspacio
				}

				im.Color = pixel.RGB(0, 1, 0)
				im.Push(pixel.V(auto.PosX, auto.PosY))
				im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
				im.Rectangle(0)

				if auto.PosX < scenes.TamanoEspacio*float64(i/2)+scenes.TamanoEspacio/2 || auto.Dir == -1 {
					auto.PosX += Velocidad * auto.Dir
				}
			}
		}
		e.Mu.Unlock()

		im.Draw(win)
		win.Update()
	}
}
