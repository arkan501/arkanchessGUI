package main

import (
	"fyne.io/fyne/v2"
    // "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type chessLayout struct {
}

// TODO: figure out what this is doing exactly
func (c *chessLayout) Layout(cells []fyne.CanvasObject, size fyne.Size) {
	squareSize := size.Width / 8

	for ix, piece := range cells {
		x := float32(ix%8) * squareSize
		y := float32(ix/8) * squareSize
		piece.Resize(fyne.NewSize(squareSize, squareSize))
		piece.Move(fyne.NewPos(x, y))
	}
}

func (b *chessLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	edge := theme.IconInlineSize() * 8
	return fyne.NewSize(edge, edge)
}

// func (b *chessLayout) MaxSize([]fyne.CanvasObject) fyne.Size {
//	return b.boardImage.Size()
// }
