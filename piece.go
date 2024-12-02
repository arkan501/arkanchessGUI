package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	p "gitlab.com/Arkan501/arkanchesslib/pieces"
)

var (
	pieceDir = "resources/piece_sets/pieces-basic-svg/"

	pieceToImage = map[p.Colour]map[p.PieceType]string{
		p.White: {
			p.Pawn:   pieceDir + "pawn-w.svg",
			p.Knight: pieceDir + "knight-w.svg",
			p.Bishop: pieceDir + "bishop-w.svg",
			p.Rook:   pieceDir + "rook-w.svg",
			p.Queen:  pieceDir + "queen-w.svg",
			p.King:   pieceDir + "king-w.svg",
		},
		p.Black: {
			p.Pawn:   pieceDir + "pawn-b.svg",
			p.Knight: pieceDir + "knight-b.svg",
			p.Bishop: pieceDir + "bishop-b.svg",
			p.Rook:   pieceDir + "rook-b.svg",
			p.Queen:  pieceDir + "queen-b.svg",
			p.King:   pieceDir + "king-b.svg",
		},
	}
)

var pieceIndex = -29

type UIPiece struct {
	*canvas.Image
	// *widget.Icon
	piece    p.IPiece
	origin   int
}

func NewUIPiece(piece p.IPiece) *UIPiece {
	pieceImage :=
		canvas.NewImageFromFile(pieceToImage[piece.Colour()][piece.Type()])
	pieceImage.FillMode = canvas.ImageFillContain

	UIpiece := &UIPiece{
		Image: pieceImage,
		piece: piece,
	}

	return UIpiece
}

func (piece *UIPiece) Tapped(ev *fyne.PointEvent) {
	log.Println("Tapped piece", piece.origin)
    pieceIndex = piece.origin
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
