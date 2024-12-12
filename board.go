package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

var boarDir = "resources/board/"

func guiBoard(chessBoard *ac.BoardState) *fyne.Container {
	var cells []fyne.CanvasObject
	boardImage := canvas.NewImageFromFile(boarDir + "aqua.svg")
	boardImage.FillMode = canvas.ImageFillOriginal

	for ix := 0; ix < 64; ix++ {
		boardPiece, err := chessBoard.GetPieceFrom(indexToSquare[ix])
		if err != nil {
			emptySquare := emptySquare(ix)
			cells = append(cells, emptySquare)
		} else {
			piece := NewUIPiece(chessBoard, boardPiece)
			piece.origin = ix
			pieceCombo := container.NewStack(piece.Image, piece)
			cells = append(cells, pieceCombo)
		}
	}

	return container.NewStack(
		boardImage,
		container.New(&chessLayout{}, cells...),
	)
}

func refreshBoard(pieceGrid *fyne.Container, chessBoard *ac.BoardState) {
	for ix := range pieceGrid.Objects {
		piece, err := chessBoard.GetPieceFrom(indexToSquare[ix])
		if err != nil {
			emptySquare := emptySquare(ix)
			pieceGrid.Objects[ix] = emptySquare
		} else {
			pieceUI := NewUIPiece(chessBoard, piece)
			pieceUI.origin = ix
			pieceCombo := container.NewStack(pieceUI.Image, pieceUI)
			pieceGrid.Objects[ix] = pieceCombo
		}
	}

	pieceGrid.Refresh()
}
