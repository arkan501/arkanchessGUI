package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

type square struct {
	*canvas.Rectangle
	boardState *ac.BoardState
	grid       *fyne.Container
	parentWin  *fyne.Window
	origin     int
}

func emptySquare(boardState *ac.BoardState, grid *fyne.Container, parentWin *fyne.Window, origin int) *square {
	return &square{
		Rectangle:  canvas.NewRectangle(color.Transparent),
		boardState: boardState,
		grid:       grid,
        parentWin: parentWin,
		origin:     origin,
	}
}

func (sq *square) Tapped(ev *fyne.PointEvent) {
	if ac.WithinBounds(index_Square[pieceIndex]) {
		targetSquare = sq.origin
		myMove := &ac.Move{}
		chooseMove(myMove, sq.boardState, sq.grid, sq.parentWin)
	}
}
