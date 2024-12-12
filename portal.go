package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
    ac "gitlab.com/Arkan501/arkanchesslib"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("arkanchess GUI")
	chessBoard := ac.NewBoard()
	guiBoard := guiBoard(&chessBoard)

	w.SetContent(guiBoard)
	w.Resize(fyne.NewSize(800, 800))

	moveBoolChan := make(chan bool)
    // moveChan := make(chan ac.Move)
    
	// This is to constantly check if these values have been set or not
	 go func() {
	 	for {
	 		if pieceIndex != targetSquare &&
	 			ac.WithinBounds(indexToSquare[pieceIndex]) &&
	 			ac.WithinBounds(indexToSquare[targetSquare]) {
	 			moveBoolChan <- true
			}
	 	}
	 }()

	// the game loop needs to run concurrently with the gui
	go func() {
		for chessBoard.GameState(chessBoard.SideToMove) == 0 {
			log.Printf("turn: %v\n", chessBoard.SideToMove)
			log.Println("executing movePiece function")
			<-moveBoolChan
			movePiece(&chessBoard, guiBoard.Objects[1].(*fyne.Container))
			log.Println("piece was moved on board")

			pieceIndex = -29
			targetSquare = -29
		}
	}()

	w.ShowAndRun()
}

func movePiece(board *ac.BoardState, pieceGrid *fyne.Container) {
    move := ac.Move{
        FromSquare: indexToSquare[pieceIndex],
        ToSquare: indexToSquare[targetSquare],
    }
    board.MakeMove(move)
	refreshBoard(pieceGrid, board)
}
