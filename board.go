package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	b "gitlab.com/Arkan501/arkanchesslib/board"
)

func guiBoard(board *b.Board) *fyne.Container {
	var cells []fyne.CanvasObject
	boardImage := canvas.NewImageFromFile("resources/board/aqua.svg")
	boardImage.FillMode = canvas.ImageFillOriginal

	for ix := 0; ix < 64; ix++ {
		boardPiece, err := board.GetPieceAt(indexToSquare[ix])
		if err != nil {
			emptySquare := emptySquare(ix)
			cells = append(cells, emptySquare)
			// cells = append(cells, canvas.NewRectangle(color.Transparent))
		} else {
			piece := NewUIPiece(boardPiece)
			piece.origin = ix
			// cells = append(cells, piece)
			pieceCombo := container.NewStack(piece.Image, piece)
			cells = append(cells, pieceCombo)
		}
	}

	return container.NewStack(
		boardImage,
		container.New(&chessLayout{}, cells...),
	)
}

func refreshBoard(pieceGrid *fyne.Container) {
	for _, cell := range pieceGrid.Objects {
		image := cell
		image.Refresh()
	}
}
