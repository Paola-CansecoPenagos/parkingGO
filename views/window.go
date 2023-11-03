// views/window.go
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
	Velocidad           = 1.0
	AnchoAuto           = 80.0
	AltoAuto            = 80.0
	AltoEspacio         = 100.0
	DistanciaEntreAutos = 10.0
)

func GenerarVehiculos(e *models.Parking) {
	rand.Seed(time.Now().UnixNano())

	for {
		auto := &models.Auto{PosX: scenes.AnchoVentana - AnchoAuto, PosY: scenes.AltoVentana, Dir: 1}
		pos := e.Enter(auto)

		if pos != -1 {
			go func(p int) {
				time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)
				e.Exit(p)
			}(pos)
		}
		time.Sleep(time.Millisecond * 1500)
	}
}

func Run(win *pixelgl.Window, e *models.Parking) {
	rand.Seed(time.Now().UnixNano())

	go GenerarVehiculos(e)

	for !win.Closed() {
		win.Clear(pixel.RGB(1, 1, 1))

		im := imdraw.New(nil)
		scenes.DibujarEstacionamiento(im, e)

		e.Mu.Lock()
		for i, auto := range e.Ocupados {
			if auto != nil {
				if auto.PosY > scenes.AltoEspacio {
					auto.PosY -= Velocidad
				} else {
					targetX := scenes.TamanoEspacio * float64(i) // Posición objetivo en X
					if auto.PosX < targetX-60 {
						auto.PosX += Velocidad
					} else if auto.PosX > targetX+20 {
						auto.PosX -= Velocidad
					} else {
						auto.PosY = scenes.AltoEspacio // Ajusta la altura del auto al nivel del carril
						auto.PosX = targetX            // Se posiciona exactamente sobre el carril
					}
				}

				im.Color = pixel.RGB(0, 1, 0)
				im.Push(pixel.V(auto.PosX, auto.PosY))
				im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
				im.Rectangle(0)
			}
		}
		e.Mu.Unlock()

		im.Draw(win)
		win.Update()
	}
}
