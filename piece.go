package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	ac "gitlab.com/Arkan501/arkanchesslib"
)

var (
	pieceDir = "resources/piece_sets/pieces-basic-svg/"

	piece_Image = map[ac.Colour]map[ac.PieceType]string{
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

type UIPiece struct {
	*canvas.Image
	// *widget.Icon
	boardState *ac.BoardState
	grid       *fyne.Container
	parentWin  *fyne.Window
	piece      ac.ChessPiece
	origin     int
}

func NewUIPiece(piece ac.ChessPiece, chessBoard *ac.BoardState, grid *fyne.Container, parentWin *fyne.Window, index int) *UIPiece {
	pieceImage := canvas.NewImageFromFile(piece_Image[piece.Colour()][piece.Type()])
	pieceImage.FillMode = canvas.ImageFillContain

	UIpiece := &UIPiece{
		Image:      pieceImage,
		boardState: chessBoard,
		grid:       grid,
		parentWin:  parentWin,
		piece:      piece,
		origin:     index,
	}

	return UIpiece
}

func (uiPiece *UIPiece) Tapped(ev *fyne.PointEvent) {
	switch ac.WithinBounds(index_Square[pieceIndex]) {
	case true:
		myPiece, _ := uiPiece.boardState.GetPieceFrom(index_Square[pieceIndex])
		otherPiece, _ := uiPiece.boardState.GetPieceFrom(index_Square[uiPiece.origin])
		switch myPiece.Colour() {
		case otherPiece.Colour():
			pieceIndex = uiPiece.origin
			return
		default:
			targetSquare = uiPiece.origin
			myMove := &ac.Move{}
			chooseMove(myMove, uiPiece.boardState, uiPiece.grid, uiPiece.parentWin)
		}
	case false:
		piece, _ := uiPiece.boardState.GetPieceFrom(index_Square[uiPiece.origin])
		switch piece.Colour() {
		case uiPiece.boardState.SideToMove:
			pieceIndex = uiPiece.origin
			return
		default:
			return
		}
	}
}

// func (uiPiece *UIPiece) Dragged(ev *fyne.DragEvent) {
//     if !ac.WithinBounds(index_Square[pieceIndex]) {
//         pieceIndex = uiPiece.origin
//         offset := squareToOffset(uiPiece.origin)
//         cell := uiPiece.grid.Objects[offset].(*fyne.Container)
//
//         position := cell.Position().Add(ev.Position)
//         uiPiece.Image.Move(ev.Position)
//     }
// }
//
// func (uiPiece *UIPiece) DragEnd() {
// }
