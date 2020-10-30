package model

import (
	"fmt"
	"testing"
)

func TestMovePiece(t *testing.T) {
	var samples = []struct {
		sx  int
		sy  int
		ex  int
		ey  int
		out int
	}{
		{0, 0, 1, 2, 1},
		{3, 3, 2, 5, 0},
		{2, 5, 2, 2, 2},
		{2, 5, 3, 7, 0},
	}

	cb := new(ChessBoard)
	cb.Init()
	cdr := new(ChessCoord)
	cdr.NCoord(3, 3)
	k := new(Knight)
	cb[cdr.Line()][cdr.Column()] = k

	for _, sample := range samples {
		s := new(ChessCoord)
		s.NCoord(sample.sx, sample.sy)
		e := new(ChessCoord)
		e.NCoord(sample.ex, sample.ey)
		if i, _ := cb.MovePiece(*s, *e); i != sample.out {
			t.Error("Failed: err:{}, expected:{}", i, sample.out)
		}
		cb.Print()
		fmt.Print("\n\n")

	}

}
