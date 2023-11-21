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
	Velocidad           = 0.5
	AnchoAuto           = 80.0
	AltoAuto            = 80.0
	AltoEspacio         = 100.0
	DistanciaEntreAutos = 10.0
)

func CalcularTiempo(auto *models.Auto) {
	time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)
	auto.State = models.StateExiting
}

func GenerarVehiculos(e *models.Parking) {
	rand.Seed(time.Now().UnixNano())

	for {
		auto := &models.Auto{PosX: -AnchoAuto - DistanciaEntreAutos, PosY: AltoAuto + AltoEspacio, Dir: 1, State: models.StateEntering}
		pos := e.Enter(auto)

		if pos != -1 {
			// Pasamos el auto como argumento a la función CalcularTiempo
			go CalcularTiempo(auto)
		}
		time.Sleep(time.Millisecond * 1500)
	}
}

func Run(win *pixelgl.Window, e *models.Parking) {

	go GenerarVehiculos(e)

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0))

		im := imdraw.New(nil)
		scenes.DibujarEstacionamiento(im, e)

		e.Mu.Lock()
		for i, auto := range e.Ocupados {
			if auto != nil {
				if auto.State == models.StateEntering {
					auto.PosX += Velocidad * auto.Dir
					auto.PosY = AltoAuto + AltoEspacio*2.2
					if auto.PosX >= 90*float64(i) {
						auto.State = models.StateParked
					}
					im.Color = pixel.RGB(133.0/255, 193.0/255, 233.0/255)
					im.Push(pixel.V(auto.PosX, auto.PosY))
					im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
					im.Rectangle(0)
				} else if auto.State == models.StateParked {
					auto.PosX = 90 * float64(i)
					auto.PosY = AltoAuto + AltoEspacio*0.2
					im.Color = pixel.RGB(247.0/255, 220.0/255, 111.0/255)
					im.Push(pixel.V(auto.PosX, auto.PosY))
					im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
					im.Rectangle(0)
				} else if auto.State == models.StateExiting {
					auto.PosX -= Velocidad * auto.Dir
					auto.PosY = AltoAuto + AltoEspacio*2.2
					im.Color = pixel.RGB(0xF5/255.0, 0xB7/255.0, 0xB1/255.0)
					im.Push(pixel.V(auto.PosX, auto.PosY))
					im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
					im.Rectangle(0)

					if auto.PosX <= -AnchoAuto-DistanciaEntreAutos {
						e.Ocupados[i] = nil
						e.Cajones[auto.Cajon] = false
					}
				}
			}
		}
		e.Mu.Unlock()

		im.Draw(win)
		win.Update()
	}
}
