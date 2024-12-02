package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	b "gitlab.com/Arkan501/arkanchesslib/board"
	p "gitlab.com/Arkan501/arkanchesslib/pieces"
)


func main() {
	myApp := app.New()
	w := myApp.NewWindow("arkanchess GUI")
	chessBoard := b.NewBoard()
	guiBoard := guiBoard(&chessBoard)

	w.SetContent(guiBoard)
	w.Resize(fyne.NewSize(800, 800))

	moveChan := make(chan bool)

	// This is to constantly check if these values have been set or not
	go func() {
		for {
			if pieceIndex != targetSquare &&
				b.WithinBounds(indexToSquare[pieceIndex]) &&
				b.WithinBounds(indexToSquare[targetSquare]) {
				moveChan <- true
			}
		}
	}()

	// the game loop needs to run concurrently with the gui
	go func() {
		turn := p.White
		for {
			log.Println("turn:", turn)
			log.Println("executing movePiece function")
			<-moveChan
			movePiece(turn, &chessBoard, guiBoard.Objects[1].(*fyne.Container))
			log.Println("piece was moved on board")

			switch turn {
			case p.White:
				turn = p.Black
			case p.Black:
				turn = p.White
			}
			pieceIndex = -29
			targetSquare = -29
		}
	}()

	w.ShowAndRun()
}

func movePiece(colour p.Colour, board *b.Board, pieceGrid *fyne.Container) {
	fromSquare := indexToSquare[pieceIndex]
	toSquare := indexToSquare[targetSquare]
    board.MakeMove(colour, fromSquare, toSquare)
	refreshBoard(pieceGrid, board)
}
