package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	b "gitlab.com/Arkan501/arkanchesslib/board"
	p "gitlab.com/Arkan501/arkanchesslib/pieces"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("arkanchess GUI")
	chessBoard := b.NewBoard()

	board := guiBoard(&chessBoard)
	w.SetContent(board)
	w.Resize(fyne.NewSize(800, 800))

	go func() {
		turn := p.White
		for chessBoard.GameState(turn) == 0 {
			movePiece(turn, &chessBoard, board)
			switch turn {
			case p.White:
				turn = p.Black
			case p.Black:
				turn = p.White
			}
		}
	}()

	w.ShowAndRun()
}

func movePiece(colour p.Colour, board *b.Board, guiBoard *fyne.Container) {
	var from int
	var to int
	board.MakeMove(colour, from, to)
	refreshBoard(guiBoard)
}
