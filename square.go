package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

var targetSquare int
type square struct {
	*canvas.Rectangle
	origin   int
}

func emptySquare(origin int) *square {
	return &square{
		Rectangle: canvas.NewRectangle(color.Transparent),
		origin:    origin,
	}
}

func (sq *square) Tapped(ev *fyne.PointEvent) {
	log.Println("Tapped square", sq.origin)
    if ac.WithinBounds(indexToSquare[pieceIndex]) {
        targetSquare = sq.origin
    }
}

// not sure if the draggable interface is necessary or not.
func (sq *square) Dragged(ev *fyne.DragEvent) {}

func (sq *square) DragEnd() {
}
