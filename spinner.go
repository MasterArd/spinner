package spinner

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	dots   = 8
	radius = 20
	size   = 8
)

type spinnerLayout struct{}

func (d *spinnerLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	diameter := float32((radius + size) * 2)
	return fyne.NewSize(diameter, diameter)
}

func (d *spinnerLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) == 0 {
		return
	}

	centerX := containerSize.Width / 2
	centerY := containerSize.Height / 2

	for i, obj := range objects {
		angle := (float64(i) / float64(len(objects))) * 2 * math.Pi

		x := centerX + float32(math.Cos(angle)*radius) - (size / 2)
		y := centerY + float32(math.Sin(angle)*radius) - (size / 2)

		obj.Move(fyne.NewPos(x, y))
		obj.Resize(fyne.NewSize(size, size))
	}
}

func New() fyne.CanvasObject {
	circles := make([]*canvas.Circle, dots)
	objects := make([]fyne.CanvasObject, dots)

	for i := range dots {
		c := canvas.NewCircle(color.NRGBA{255, 255, 255, 50})
		circles[i] = c
		objects[i] = c
	}

	phase := 0
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			for i := 0; i < dots; i++ {
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
				circles[i].Refresh()
			}
			phase = (phase + 1) % dots
		}
	}()

	return container.New(&spinnerLayout{}, objects...)
}
