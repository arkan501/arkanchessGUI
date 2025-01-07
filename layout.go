package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type chessLayout struct{}
type promotionLayout struct{}

func (c *chessLayout) Layout(cells []fyne.CanvasObject, size fyne.Size) {
	squareSize := fyne.Min(size.Width, size.Height) / 8
	totalSize := squareSize * 8
	offsetx := (size.Width - totalSize) / 2
	offsety := (size.Height - totalSize) / 2

	for ix, piece := range cells {
		x := float32(ix%8)*squareSize + offsetx
		y := float32(ix/8)*squareSize + offsety
		piece.Resize(fyne.NewSize(squareSize, squareSize))
		piece.Move(fyne.NewPos(x, y))
	}
}

// Should I even have these?
func (c *chessLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	edge := theme.IconInlineSize() * 8
	return fyne.NewSize(edge, edge)
}

func (p *promotionLayout) Layout(cells []fyne.CanvasObject, size fyne.Size) {
	squareSize := fyne.Min(size.Width, size.Height) / 8
	totalSize := squareSize * 8
	offsetx := (size.Width - totalSize) / 2
	offsety := (size.Height - totalSize) / 2

    for ix := 0; ix < 4; ix++ {
        x := float32(ix)*squareSize + offsetx
        y := float32(ix)*squareSize + offsety
        cells[ix].Resize(fyne.NewSize(squareSize, squareSize))
        cells[ix].Move(fyne.NewPos(x, y))
    }
}
