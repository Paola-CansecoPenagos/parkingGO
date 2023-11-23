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

func CalcularTiempo(auto *models.Auto) {
	time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)
	auto.State = models.StateExiting
}

func GenerarVehiculos(e *models.Parking) {
	rand.Seed(time.Now().UnixNano())

	for {
		auto := &models.Auto{PosX: -models.AnchoAuto - models.DistanciaEntreAutos, PosY: models.AltoAuto + models.AltoEspacio, Dir: 1, State: models.StateEntering}
		pos := e.Enter(auto)

		if pos != -1 {
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
					auto.MoveAndDraw(im, i)
				} else if auto.State == models.StateParked {
					auto.DrawParked(im)
				} else if auto.State == models.StateExiting {
					auto.DrawExiting(im, e, i)
				}
			}
		}
		e.Mu.Unlock()

		im.Draw(win)
		win.Update()
	}
}
