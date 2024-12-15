package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

var pieceIndex = -29
var targetSquare = -29
var moveReady = make(chan bool)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("arkanchess GUI")
	chessBoard := ac.NewBoard()
	guiBoard := guiBoard(&chessBoard)

	w.SetContent(guiBoard)
	w.Resize(fyne.NewSize(800, 800))

	// the game loop needs to run concurrently with the gui
	go gameEngine(&chessBoard, guiBoard)

	w.ShowAndRun()
}

func gameEngine(chessBoard *ac.BoardState, guiBoard *fyne.Container) {
	for chessBoard.GameState() == 0 {
		ready := make(chan bool)
		move := ac.Move{}
		<-moveReady
		go chooseMove(ready, &move)
        <-ready
        // if chessBoard.CanPromote(chessBoard.SideToMove, move.ToSquare) {
        //     move.Promotion = selectPromotion(chessBoard, guiBoard)
        // }
		chessBoard.MakeMove(move)
        refreshBoard(guiBoard.Objects[1].(*fyne.Container), chessBoard)
        pieceIndex = -29
        targetSquare = -29
	}
}

func chooseMove(ready chan bool, move *ac.Move) {
	move.FromSquare = indexToSquare[pieceIndex]
	move.ToSquare = indexToSquare[targetSquare]
	ready <- true
}

// func selectPromotion(chessBoard *ac.BoardState, guiBoard *fyne.Container) ac.PieceType {
//     var choice ac.PieceType
//     selection := container.NewVBox()
// 
//     return choice
// }
