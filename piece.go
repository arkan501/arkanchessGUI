package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

    ac "gitlab.com/Arkan501/arkanchesslib"
)

var (
	pieceDir = "resources/piece_sets/pieces-basic-svg/"

	pieceToImage = map[ac.Colour]map[ac.PieceType]string{
		ac.White: {
			ac.Pawn:   pieceDir + "pawn-w.svg",
			ac.Knight: pieceDir + "knight-w.svg",
			ac.Bishop: pieceDir + "bishop-w.svg",
			ac.Rook:   pieceDir + "rook-w.svg",
			ac.Queen:  pieceDir + "queen-w.svg",
			ac.King:   pieceDir + "king-w.svg",
		},
		ac.Black: {
			ac.Pawn:   pieceDir + "pawn-b.svg",
			ac.Knight: pieceDir + "knight-b.svg",
			ac.Bishop: pieceDir + "bishop-b.svg",
			ac.Rook:   pieceDir + "rook-b.svg",
			ac.Queen:  pieceDir + "queen-b.svg",
			ac.King:   pieceDir + "king-b.svg",
		},
	}
)

var pieceIndex = -29

type UIPiece struct {
	*canvas.Image
	// *widget.Icon
	boardState *ac.BoardState
	piece      ac.IPiece
	origin     int
}

func NewUIPiece(chessBoard *ac.BoardState, piece ac.IPiece) *UIPiece {
	pieceImage :=
		canvas.NewImageFromFile(pieceToImage[piece.Colour()][piece.Type()])
	pieceImage.FillMode = canvas.ImageFillContain

	UIpiece := &UIPiece{
		Image: pieceImage,
        boardState: chessBoard,
		piece: piece,
	}

	return UIpiece
}

func (uiPiece *UIPiece) Tapped(ev *fyne.PointEvent) {
	log.Println("Tapped piece", uiPiece.origin)
	if !ac.WithinBounds(indexToSquare[pieceIndex]) {
		pieceIndex = uiPiece.origin
		return
	}

	myPiece, _ := uiPiece.boardState.GetPieceFrom(indexToSquare[pieceIndex])
	otherPiece, _ := uiPiece.boardState.GetPieceFrom(indexToSquare[uiPiece.origin])
    if myPiece.Colour() == otherPiece.Colour() {
        pieceIndex = uiPiece.origin
    } else {
        targetSquare = uiPiece.origin
    }
}

// func (piece *UIPiece) Dragged(ev *fyne.DragEvent) {
// 	log.Println("Dragged square", piece.origin)
// 	piece.Move(ev.Position)
// }
//
// func (piece *UIPiece) DragEnd() {
// 	log.Println("DragEnd square")
//
// }
