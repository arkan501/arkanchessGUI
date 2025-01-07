package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

var boardDir = "resources/board/"

func guiBoard(chessBoard *ac.BoardState, parentWin *fyne.Window) *fyne.Container {
	var cells []fyne.CanvasObject
	pieceGrid := container.New(&chessLayout{})
	boardImage := canvas.NewImageFromFile(boardDir + "aqua.svg")
	boardImage.FillMode = canvas.ImageFillOriginal

	for ix := 0; ix < 64; ix++ {
		boardPiece, err := chessBoard.GetPieceFrom(index_Square[ix])
		if err != nil {
			emptySquare := emptySquare(chessBoard, pieceGrid, parentWin, ix)
			cells = append(cells, emptySquare)
		} else {
			piece := NewUIPiece(boardPiece, chessBoard, pieceGrid, parentWin, ix)
			pieceCombo := container.NewStack(piece.Image, piece)
			cells = append(cells, pieceCombo)
		}
	}

	pieceGrid.Objects = cells

	return container.NewStack(
		boardImage,
		pieceGrid,
	)
}

func refreshBoard(pieceGrid *fyne.Container, chessBoard *ac.BoardState, parentWin *fyne.Window) {
	go func() {
		for ix := range pieceGrid.Objects {
			piece, err := chessBoard.GetPieceFrom(index_Square[ix])
			if err != nil {
				emptySquare := emptySquare(chessBoard, pieceGrid, parentWin, ix)
				pieceGrid.Objects[ix] = emptySquare
			} else {
				uiPiece := NewUIPiece(piece, chessBoard, pieceGrid, parentWin, ix)
				pieceCombo := container.NewStack(uiPiece.Image, uiPiece)
				pieceGrid.Objects[ix] = pieceCombo
			}
		}
	}()

	pieceGrid.Refresh()
}
