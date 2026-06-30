package spinner

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func New() fyne.CanvasObject {
	const (
		dots   = 8
		radius = 20
		size   = 8
	)

	circles := make([]*canvas.Circle, dots)
	objects := make([]fyne.CanvasObject, dots)

	for i := range dots {
		c := canvas.NewCircle(color.NRGBA{255, 255, 255, 50})
		c.Resize(fyne.NewSize(size, size))
		circles[i] = c
		objects[i] = c
	}

	phase := 0
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			for i := 0; i < dots; i++ {
				angle := (float64(i) / dots) * 2 * math.Pi

				x := float32(math.Cos(angle)*radius + radius + size)
				y := float32(math.Sin(angle)*radius + radius + size)

				circles[i].Move(fyne.NewPos(x, y))

				if i == phase {
					circles[i].FillColor = color.White
				} else {
					alpha := uint8(30 + ((i+dots-phase)%dots)*20)
					circles[i].FillColor = color.NRGBA{
						R: 255,
						G: 255,
						B: 255,
						A: alpha,
					}
				}
				fyne.Do(func() {
					circles[i].Refresh()
				})
			}

			phase = (phase + 1) % dots
		}
	}()

	return container.NewWithoutLayout(objects...)
}
