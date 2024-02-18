package internal

import (
	"errors"
	"reflect"
	"testing"
)

// assertPieceConsistent confirms that piece is at position pos and that
// the board's piece at position pos is this piece.
func assertPieceConsistent(t *testing.T, b *Board, piece ChessPiece, want Position) {
	t.Helper()
	if got := piece.GetPosition(); got != nil && *got != want {
		t.Errorf("assertPieceConsistent invalid, got %v, wanted %v", got, want)
	}
	if got := b.GetPieceAtPosition(want); got != piece {
		t.Errorf("GetPieceAtPosition invalid, got %v, wanted %v", got, piece)
	}
}

func TestMovePiece(t *testing.T) {
	b := NewBoard()
	src := NewRook(WHITE)
	mustPlace(t, b, src, "d3")

	if err := b.MovePiece(src, mustPosition(t, "d5")); err != nil {
		t.Errorf("MovePiece returned err %v, wanted nil", err)
	}

	assertPieceConsistent(t, b, src, mustPosition(t, "d5"))
}

func TestMovePieceStationary(t *testing.T) {
	b := NewBoard()
	src := NewRook(WHITE)
	mustPlace(t, b, src, "d3")

	if err := b.MovePiece(src, mustPosition(t, "d3")); err != nil {
		t.Errorf("MovePiece returned err %v, wanted nil", err)
	}

	assertPieceConsistent(t, b, src, mustPosition(t, "d3"))
}

func TestMovePieceCapturing(t *testing.T) {
	b := NewBoard()
	src := NewRook(WHITE)
	mustPlace(t, b, src, "d3")
	dest := NewRook(BLACK)
	mustPlace(t, b, dest, "d5")

	if err := b.MovePiece(src, mustPosition(t, "d5")); err != nil {
		t.Errorf("MovePiece returned err %v, wanted nil", err)
	}

	assertPieceConsistent(t, b, src, mustPosition(t, "d5"))
	if captured := dest.GetPosition(); captured != nil {
		t.Errorf("Captured piece should have nil position, was %v", captured)
	}
}

func TestMoveIllegalPiece(t *testing.T) {
	b := NewBoard()
	src := NewRook(WHITE)
	mustPlace(t, b, src, "d3")

	err := b.MovePiece(src, mustPosition(t, "e4"))

	want := errors.New("MovePiece: White Rook at d3 cannot move to e4")
	// Generally undesirable, but want to verify error strings.
	if !reflect.DeepEqual(err, want) {
		t.Errorf("MovePiece returned err %v, wanted %v", err, want)
	}
	assertPieceConsistent(t, b, src, mustPosition(t, "d3"))
	if piece := b.GetPieceAtPosition(mustPosition(t, "e4")); piece != nil {
		t.Errorf("GetPieceAtPosition after failed move should be nil, was %v", piece)
	}
}

func TestMovePieceNotOnBoard(t *testing.T) {
	b := NewBoard()
	src := NewRook(WHITE)

	err := b.MovePiece(src, mustPosition(t, "e4"))

	want := errors.New("MovePiece: Off board White Rook is not on the board")
	// Generally undesirable, but want to verify error strings.
	if !reflect.DeepEqual(err, want) {
		t.Errorf("MovePiece returned err %v, wanted %v", err, want)
	}
	if pos := src.GetPosition(); pos != nil {
		t.Errorf("Moving piece off board should have nil position, was %v", pos)
	}
	if piece := b.GetPieceAtPosition(mustPosition(t, "e4")); piece != nil {
		t.Errorf("GetPieceAtPosition after failed place should be nil, was %v", piece)
	}
}

func TestMovePieceOnTop(t *testing.T) {
	b := NewBoard()
	src := NewRook(WHITE)
	mustPlace(t, b, src, "d3")
	dest := NewRook(BLACK)

	err := b.PlacePiece(dest, "d3")

	want := errors.New("PlacePiece: cannot place Off board Black Rook on top of White Rook at d3")
	// Generally undesirable, but want to verify error strings.
	if !reflect.DeepEqual(err, want) {
		t.Errorf("PlacePiece returned err %v, wanted %v", err, want)
	}
	assertPieceConsistent(t, b, src, mustPosition(t, "d3"))
	if pos := dest.GetPosition(); pos != nil {
		t.Errorf("Failing to place piece should have nil position, was %v", pos)
	}
}
