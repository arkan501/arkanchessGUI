package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	b "gitlab.com/Arkan501/arkanchesslib/board"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("arkanchess GUI")
    chessBoard := b.NewBoard()

	board := createBoard(&chessBoard)

	w.SetContent(board)
	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()
}

func createBoard(board *b.Board) *fyne.Container {
	var cells []fyne.CanvasObject
    boardImage := canvas.NewImageFromFile("resources/board/aqua.svg")
    boardImage.FillMode = canvas.ImageFillOriginal

	for ix := 0; ix < 64; ix++ {
        boardPiece, err := board.GetPieceAt(indexToSquare[ix])
        if err != nil {
            cells = append(cells, canvas.NewRectangle(color.Transparent))
        } else {
            piece :=
            canvas.NewImageFromFile(
            pieceToImage[boardPiece.Colour()][boardPiece.Type()],
            )
            piece.FillMode = canvas.ImageFillContain
            cells = append(cells, piece)
        }
	}

	return container.NewStack(boardImage, container.New(&chessLayout{}, cells...))
}
