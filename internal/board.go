package internal

// Board represents the state of the board and keeps track of
// pieces on it.  It coordinates interactions between pieces and the player.
type Board struct{}

// Stub method responsible for placing a piece on the board at a particular
// position.  Returns an error if the position was already occupied or pos is
// not valid chess notation
func (b *Board) PlacePiece(piece ChessPiece, pos string) error {
	_, err := NewPosition(pos)
	if err != nil {
		return err
	}
	return nil
}

// Stub method responsible for moving a piece from one position on the board
// to another.
func (b *Board) MovePiece(piece ChessPiece, pos Position) error {
	return nil
}
