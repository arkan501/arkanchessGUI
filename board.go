package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	b "gitlab.com/Arkan501/arkanchesslib/board"
)

var boarDir = "resources/board/"

func guiBoard(board *b.Board) *fyne.Container {
	var cells []fyne.CanvasObject
	boardImage := canvas.NewImageFromFile(boarDir + "aqua.svg")
	boardImage.FillMode = canvas.ImageFillOriginal

	for ix := 0; ix < 64; ix++ {
		boardPiece, err := board.GetPieceFrom(indexToSquare[ix])
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

func refreshBoard(pieceGrid *fyne.Container, chessBoard *b.Board) {
	for ix := range pieceGrid.Objects {
		piece, err := chessBoard.GetPieceFrom(indexToSquare[ix])
		if err != nil {
			emptySquare := emptySquare(ix)
			pieceGrid.Objects[ix] = emptySquare
		} else {
			pieceUI := NewUIPiece(piece)
			pieceUI.origin = ix
			pieceCombo := container.NewStack(pieceUI.Image, pieceUI)
			pieceGrid.Objects[ix] = pieceCombo
		}
	}

	pieceGrid.Refresh()
}
