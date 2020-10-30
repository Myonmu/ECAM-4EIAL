package model

import (
	"errors"
	"fmt"
)

//TODO Implement type Board

type ChessBoard [8][8]Piece

func (cb *ChessBoard) Init() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			e := new(Empty)
			cb[i][j] = e
		}

	}
}

func (cb *ChessBoard) MovePiece(start ChessCoord, end ChessCoord) (int, error) {
	if cb.IsEmptyAt(start) {
		return 1, errors.New("No gamepiece at the starting coordinate")
	}
	pc := cb.GetPieceAt(start)
	if !pc.LegalMove(start, end) {
		return 2, errors.New("Illegal")

	}
	cb[end.line][end.column] = pc
	cb[start.line][start.column] = new(Empty)

	return 0, nil
}

func (cb *ChessBoard) IsEmptyAt(c ChessCoord) bool {

	return cb[c.line][c.column].GetType() == "Empty"
}

func (cb *ChessBoard) GetPieceAt(c ChessCoord) Piece {
	return cb[c.line][c.column]
}

func (cb *ChessBoard) Print() {
	for i := range cb {
		for j := range cb[0] {
			obj := cb[i][j]
			fmt.Print(obj.Print() + " ")
		}
		fmt.Print("\n")

	}

}

type Move interface {
	MovePiece()
}

type Capture interface {
	capturePiece()
}

type PieceCheck interface {
	IsEmptyAt()
	GetPieceAt()
}

type Graphics interface {
	Print()
}
