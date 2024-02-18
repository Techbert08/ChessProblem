package internal

import (
	"testing"
)

// mustPlace forces a piece onto the board, failing on error.
func mustPlace(t *testing.T, b *Board, p ChessPiece, pos string) {
	t.Helper()
	if err := b.PlacePiece(p, pos); err != nil {
		t.Errorf("PlacePiece(%v, %v) returned err %v, expected nil", p, pos, err)
	}
}

// mustPosition forces a string to parse as a Position, failing on error
func mustPosition(t *testing.T, pos string) Position {
	t.Helper()
	p, err := NewPosition(pos)
	if err != nil {
		t.Fatalf("could not parse position %v, err: %v", pos, err)
	}
	return *p
}
