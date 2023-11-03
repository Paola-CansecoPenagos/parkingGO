package main

import (
	"parking/models"
	"parking/views"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "Parking",
			Bounds: pixel.R(0, 0, 2000, 600),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		e := models.NewParking(20)
		views.Run(win, e)
	})
}
