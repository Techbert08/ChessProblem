package internal

// Holds White or Black
type Color int

const (
	WHITE Color = iota
	BLACK
)

// ChessPiece is the interface implemented by each piece allowing them
// to be moved around the board legally.
type ChessPiece interface {
	IsLegalMove(pos Position) bool
}

// Rook is a normal chess Rook, identified by color C
type Rook struct {
	C Color
}

// Stub method saying any move is legal for a Rook.
func (r *Rook) IsLegalMove(pos Position) bool {
	return true
}

// Bishop is a normal chess Bishop, identified by color C
type Bishop struct {
	C Color
}

// Stub method saying any move is illegal for a Bishop.
func (b *Bishop) IsLegalMove(pos Position) bool {
	return false
}
