package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

var pieceIndex = -29
var targetSquare = -29

func main() {
	myApp := app.New()
	w := myApp.NewWindow("arkanchess GUI")
	chessBoard := ac.NewBoard()
	guiBoard := guiBoard(&chessBoard, &w)

	w.SetContent(guiBoard)
	w.Resize(fyne.NewSize(800, 800))

	w.ShowAndRun()
}

func chooseMove(move *ac.Move, chessBoard *ac.BoardState, pieceGrid *fyne.Container, parentWin *fyne.Window) {
	move.FromSquare = index_Square[pieceIndex]
	move.ToSquare = index_Square[targetSquare]
	if piece, err := chessBoard.GetPieceFrom(move.FromSquare); err == nil {
		switch piece.Type() {
		case ac.Pawn:
			if chessBoard.CanPromote(piece.Colour(), move.ToSquare) {
				selectPromotion(move, chessBoard, pieceGrid, *parentWin)
                pieceIndex = -29
                targetSquare = -29
                return
			}
			fallthrough
		default:
			chessBoard.MakeMove(*move)
		}
	}
	refreshBoard(pieceGrid, chessBoard, parentWin)
	pieceIndex = -29
	targetSquare = -29
}
