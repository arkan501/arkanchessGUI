package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	ac "gitlab.com/Arkan501/arkanchesslib"
)

var pieceType_String = map[ac.PieceType]string{
	ac.King:   "king",
	ac.Queen:  "queen",
	ac.Bishop: "bishop",
	ac.Knight: "knight",
	ac.Rook:   "rook",
	ac.Pawn:   "pawn",
}

func selectPromotion(move *ac.Move, chessBoard *ac.BoardState, pieceGrid *fyne.Container, parent fyne.Window) {
	choices := container.NewVBox()
	cells := []fyne.CanvasObject{}
	myPopUp := widget.NewPopUp(choices, parent.Canvas())

	for _, pieceType := range []ac.PieceType{ac.Queen, ac.Rook, ac.Bishop, ac.Knight} {
		resource, _ := fyne.LoadResourceFromPath(piece_Image[chessBoard.SideToMove][pieceType])
		cells = append(cells, widget.NewButtonWithIcon(
			"",
			resource,
			func() {
				move.Promotion = pieceType
                chessBoard.MakeMove(*move)
                refreshBoard(pieceGrid, chessBoard, &parent)
				myPopUp.Hide()
			}),
		)
	}
	choices.Objects = cells

	myPopUp.Show()
}
