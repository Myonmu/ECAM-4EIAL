package board

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	str := ""
	for x := 0; x < len(c[0]); x++ {
		for y := 0; y < len(c); y++ {
			crd := coord.NewCartesian(y, x)
			p, _ := c.PieceAt(crd)
			if p == nil {
				str = str + " ."
			} else {
				str = str + " " + p.String()
			}
		}
		str = str + "\n"
	}
	return str
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) (piece.Piece, error) {
	a, e1 := at.Coord(0)
	b, e2 := at.Coord(1)
	if e1 != nil {
		return nil, e1
	}
	if e2 != nil {
		return nil, e2
	}
	return c[a][b], nil
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	p1, e1 := c.PieceAt(from)
	p2, e2 := c.PieceAt(to)
	if e1 != nil {
		return e1
	}
	if e2 != nil {
		return e2
	}
	if p1 == nil {
		return fmt.Errorf("Error: no piece at starting position")
	}

	if p1.LegalMove() {
		if (p2 != nil && p1.Color() != p2.Color()) || p2 == nil {
			c.PlacePieceAt(p1, to)
		} else {
			return fmt.Errorf("Target square occupied by friendly piece")
		}

		c.PlacePieceAt(nil, from)
	} else {
		return fmt.Errorf("Illegal move")
	}
	return nil
}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	a, e1 := at.Coord(0)
	b, e2 := at.Coord(1)
	if e1 != nil {
		return e1
	}
	if e2 != nil {
		return e2
	}
	c[a][b] = p
	return nil
}
