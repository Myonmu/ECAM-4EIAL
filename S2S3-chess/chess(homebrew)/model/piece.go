package model

type Piece interface {
	Print() string
	GetType() string
	LegalMove(start ChessCoord, end ChessCoord) bool
	GetFaction() string
}

type NKnight interface {
	NewKnight(faction string)
}

type NEmpty interface {
	NewEmpty()
}

type Knight struct {
	faction string
}

func (k *Knight) NewKnight(faction string) {
	k.faction = faction
}

func (k *Knight) LegalMove(start ChessCoord, end ChessCoord) bool {
	if start.LineDelta(end) == 1 && start.ColumnDelta(end) == 2 {
		return true
	}
	if start.LineDelta(end) == 2 && start.ColumnDelta(end) == 1 {
		return true
	}
	return false
}

func (k *Knight) GetFaction() string {
	return k.faction
}

func (k *Knight) GetType() string {
	return "Knight"
}

func (k *Knight) Print() string {
	return "K"
}

type Empty struct{}

func (e *Empty) Print() string {
	return "."
}

func (e *Empty) GetType() string {
	return "Empty"
}

func (e *Empty) LegalMove(start ChessCoord, end ChessCoord) bool {
	return false
}

func (e *Empty) GetFaction() string {
	return "None"
}

func (e *Empty) NewEmpty() {

}
