package internal

import (
	"fmt"
)

// Board represents the state of the board and keeps track of
// pieces on it.  It coordinates interactions between pieces and the player.
type Board struct {
	positions map[Position]ChessPiece
}

func NewBoard() *Board {
	return &Board{
		positions: make(map[Position]ChessPiece),
	}
}

// PlacePiece places a piece on the board at a particular
// position.  Returns an error if the position was already occupied or pos is
// not valid chess notation.
func (b *Board) PlacePiece(piece ChessPiece, pos string) error {
	p, err := NewPosition(pos)
	if err != nil {
		return err
	}
	if current := b.positions[*p]; current != nil {
		return fmt.Errorf("PlacePiece: cannot place %v on top of %v", piece, current)
	}
	b.positions[*p] = piece
	piece.place(b, *p)
	return nil
}

// Moves a piece from one position on the board to another.
func (b *Board) MovePiece(piece ChessPiece, pos Position) error {
	current := piece.GetPosition()
	if current == nil {
		return fmt.Errorf("MovePiece: %v is not on the board", piece)
	}
	if !piece.IsLegalMove(pos) {
		return fmt.Errorf("MovePiece: %v cannot move to %v", piece, pos)
	}
	destPiece := b.GetPieceAtPosition(pos)
	if destPiece != nil {
		destPiece.remove()
	}
	delete(b.positions, *current)
	b.positions[pos] = piece
	piece.place(b, pos)
	return nil
}

// Gets the piece at a given position, or nil if the space is empty.
func (b *Board) GetPieceAtPosition(pos Position) ChessPiece {
	return b.positions[pos]
}
