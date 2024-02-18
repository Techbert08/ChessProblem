package internal

import (
	"fmt"
)

// Position represents a place on the board identified by rank and file.
// The zero value is the bottom left corner, a1.
type Position struct {
	// rank and file are 0-7 representing positions on the chess board.
	rank, file int
}

// NewPosition constructs a new Position struct from standard chess position notation.
// The input string must be two characters, the first a file letter ranging from a-h and
// the second a rank number from 1-8.
func NewPosition(p string) (*Position, error) {
	if len(p) != 2 {
		return nil, fmt.Errorf("NewPosition: invalid string. expected len(2), got %v", p)
	}
	f := rune(p[0])
	if f < 'a' || f > 'h' {
		return nil, fmt.Errorf("NewPosition: first file character should be a-h, got %c", p[0])
	}
	r := rune(p[1])
	if r < '1' || r > '8' {
		return nil, fmt.Errorf("NewPosition: second rank character should be 1-8, got %c", p[1])
	}
	return &Position{
		rank: int(r - '1'),
		file: int(f - 'a'),
	}, nil
}

// Move computes a new position offset from the current one.
// positive f moves to a higher lettered file (wrapping around if beyond h)
// positive r moves to a higher numbered rank (wrapping around if beyond 8)
func (p Position) Move(f int, r int) Position {
	rank := (r + p.rank) % 8
	if rank < 0 {
		rank = rank + 8
	}
	file := (f + p.file) % 8
	if file < 0 {
		file = file + 8
	}
	return Position{
		rank: rank,
		file: file,
	}
}

// String converts the position to standard chess notation (i.e. a1 for rank 1 file a)
func (p Position) String() string {
	return fmt.Sprintf("%c%v", rune(p.file+'a'), p.rank+1)
}
