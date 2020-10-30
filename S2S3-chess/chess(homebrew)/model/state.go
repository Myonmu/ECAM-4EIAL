package model

import (
	"errors"
	"math"
	"strings"
)

//TODO Implement game state related elements here

type ChessCoord struct {
	line   int
	column int
}

func (cc *ChessCoord) NCoord(line, column int) {
	cc.line = line
	cc.column = column
}

func (cc *ChessCoord) ReadFrom(raw string) (string, error) {
	var formalized = strings.Replace(raw, " ", "", -1)
	if len(formalized) > 2 {
		return "", errors.New("Coordinate syntax error")
	}
	cc.column = int(rune(formalized[0]) - 65)
	cc.line = 8 - int(rune(formalized[1])-48)
	return formalized, nil
}

func (cc *ChessCoord) Print() string {
	c := rune(cc.column + 65)
	l := rune(8 - cc.line + 48)
	return string(c) + string(l)
}

func (cc *ChessCoord) ColumnDelta(other ChessCoord) int {
	return int(math.Abs(float64(cc.column - other.column)))
}

func (cc *ChessCoord) LineDelta(other ChessCoord) int {
	return int(math.Abs(float64(cc.line - other.line)))
}

func (cc *ChessCoord) Line() int {
	return cc.line
}

func (cc *ChessCoord) Column() int {
	return cc.column
}

type CoordinateIO interface {
	ReadFrom()
	NCoord()
	Line()
	Column()
	Print()
}

type CoordinateMath interface {
	ColumnDelta()
	LineDelta()
}
