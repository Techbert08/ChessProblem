package internal

import (
	"fmt"
)

// Color labels pieces WHITE or BLACK
type Color int

const (
	EMPTY Color = iota
	WHITE
	BLACK
)

// ChessPiece is the interface implemented by each piece allowing them
// to be moved around the board legally.
// Unexported functions are used in partnership with Board.
type ChessPiece interface {
	// Pulls in String() to emit a human readable representation.
	fmt.Stringer

	// IsLegalMove returns true is dest is a legal move for this piece
	// and false otherwise.
	IsLegalMove(dest Position) bool

	// GetPosition returns the current position of this piece on the
	// board.  Returns nil if not on the board.
	GetPosition() *Position

	// GetColor returns the Color of this piece.
	GetColor() Color

	// place puts this ChessPiece on the given Board at Position p
	place(b *Board, p Position)

	// remove clears this ChessPiece off of a board.
	remove()
}

// Rook is a normal chess Rook
type Rook struct {
	basicPiece
}

// NewRook builds a new Rook off-board.
// It can be added to the board by PlacePiece.
func NewRook(c Color) *Rook {
	return &Rook{
		basicPiece{
			name:  "Rook",
			color: c,
		},
	}
}

// Bishop is a normal chess Bishop.
type Bishop struct {
	basicPiece
}

// NewBishop builds a new Bishop off-board.
// It can be added to the board by PlacePiece.
func NewBishop(c Color) *Bishop {
	return &Bishop{
		basicPiece{
			name:  "Bishop",
			color: c,
		},
	}
}

// basicPiece collects behaviors every piece needs to represent position
// on the board and do basic movements.
type basicPiece struct {
	// name is the name of this piece for printing.
	name string

	// color is the color of the given piece.
	color Color

	// position is the pieces current position on the board.
	position Position

	// board is a pointer to the board this piece is on.
	board *Board
}

func (bP *basicPiece) GetColor() Color {
	return bP.color
}

func (bP *basicPiece) place(b *Board, p Position) {
	bP.board = b
	bP.position = p
}

func (bP *basicPiece) remove() {
	bP.board = nil
}

func (bP *basicPiece) GetPosition() *Position {
	if bP.board != nil {
		return &bP.position
	}
	return nil
}

func (bP *basicPiece) String() string {
	color := "Unknown"
	if bP.color == WHITE {
		color = "White"
	} else if bP.color == BLACK {
		color = "Black"
	} else if bP.color == EMPTY {
		color = "Empty"
	}
	if bP.board == nil {
		return fmt.Sprintf("Off board %v %v", color, bP.name)
	}
	return fmt.Sprintf("%v %v at %v", color, bP.name, bP.position)
}

// checkClearSpaces returns true if all Positions in positions are clear
// on the provided Board b.
func checkClearSpaces(positions []Position, b *Board) bool {
	for _, p := range positions {
		if b.GetPieceAtPosition(p) != nil {
			return false
		}
	}
	return true
}

func (r *Rook) IsLegalMove(dest Position) bool {
	if r.board == nil {
		// Not on the board
		return false
	}
	if (dest.rank == r.position.rank) && (dest.file == r.position.file) {
		// Can stay put.
		return true
	}
	if (dest.rank != r.position.rank) && (dest.file != r.position.file) {
		// If not in at least the same rank or same file, this is impossible.
		return false
	}
	destPiece := r.board.GetPieceAtPosition(dest)
	if destPiece != nil && destPiece.GetColor() == r.GetColor() {
		// This is running into a piece on this Rook's side.
		return false
	}
	// Now check for collisions in relevant directions.
	if dest.rank == r.position.rank {
		rightSpaces := make([]Position, 0)
		search := r.position.Move(1, 0)
		for search != dest && search != r.position {
			rightSpaces = append(rightSpaces, search)
			search = search.Move(1, 0)
		}
		if search == dest && checkClearSpaces(rightSpaces, r.board) {
			// Clear path to right.
			return true
		}
		leftSpaces := make([]Position, 0)
		search = r.position.Move(-1, 0)
		for search != dest && search != r.position {
			leftSpaces = append(leftSpaces, search)
			search = search.Move(-1, 0)
		}
		// At this point either left is clear or there is no path.
		return search == dest && checkClearSpaces(leftSpaces, r.board)
	} else if dest.file == r.position.file {
		upSpaces := make([]Position, 0)
		search := r.position.Move(0, 1)
		for search != dest && search != r.position {
			upSpaces = append(upSpaces, search)
			search = search.Move(0, 1)
		}
		if search == dest && checkClearSpaces(upSpaces, r.board) {
			// Clear path up.
			return true
		}
		downSpaces := make([]Position, 0)
		search = r.position.Move(0, -1)
		for search != dest && search != r.position {
			downSpaces = append(downSpaces, search)
			search = search.Move(0, -1)
		}
		// At this point either down is clear or there is no path
		return search == dest && checkClearSpaces(downSpaces, r.board)
	}
	// This should never happen, as it indicates a logic error in this function
	// The earlier traps ensuring equal rank or file should have headed this off.
	panic(fmt.Sprintf("error computing move of %v to %v", r, dest))
}

func (b *Bishop) IsLegalMove(dest Position) bool {
	if b.board == nil {
		// Not on the board.
		return false
	}
	if (dest.rank == b.position.rank) && (dest.file == b.position.file) {
		// Can stay put
		return true
	}
	destPiece := b.board.GetPieceAtPosition(dest)
	if destPiece != nil && destPiece.GetColor() == b.GetColor() {
		// This is running into a piece on this piece's side.
		return false
	}
	// Now check for collisions in relevant directions.
	upRightSpaces := make([]Position, 0)
	search := b.position.Move(1, 1)
	for search != dest && search != b.position {
		upRightSpaces = append(upRightSpaces, search)
		search = search.Move(1, 1)
	}
	if search == dest && checkClearSpaces(upRightSpaces, b.board) {
		// Clear path to up right.
		return true
	}
	upLeftSpaces := make([]Position, 0)
	search = b.position.Move(-1, 1)
	for search != dest && search != b.position {
		upLeftSpaces = append(upLeftSpaces, search)
		search = search.Move(-1, 1)
	}
	if search == dest && checkClearSpaces(upLeftSpaces, b.board) {
		// Clear path to up left.
		return true
	}
	downRightSpaces := make([]Position, 0)
	search = b.position.Move(1, -1)
	for search != dest && search != b.position {
		downRightSpaces = append(downRightSpaces, search)
		search = search.Move(1, -1)
	}
	if search == dest && checkClearSpaces(downRightSpaces, b.board) {
		// Clear path to down right.
		return true
	}
	downLeftSpaces := make([]Position, 0)
	search = b.position.Move(-1, -1)
	for search != dest && search != b.position {
		downLeftSpaces = append(downLeftSpaces, search)
		search = search.Move(-1, -1)
	}
	return search == dest && checkClearSpaces(downLeftSpaces, b.board)
}
