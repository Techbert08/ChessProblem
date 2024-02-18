package internal

import (
	"testing"
)

var rookMovementTestCases = []struct {
	start          string
	color          Color
	dest           string
	whitePositions []string
	blackPositions []string
	want           bool
}{
	{"b2", WHITE, "b2", []string{}, []string{}, true},
	{"b2", WHITE, "b4", []string{}, []string{}, true},
	{"b2", WHITE, "b1", []string{}, []string{}, true},
	{"b2", WHITE, "c2", []string{}, []string{}, true},
	{"b2", WHITE, "a2", []string{}, []string{}, true},
	{"b2", WHITE, "c3", []string{}, []string{}, false},
	{"b2", WHITE, "a1", []string{}, []string{}, false},
	{"b2", WHITE, "a3", []string{}, []string{}, false},
	{"b2", WHITE, "c1", []string{}, []string{}, false},
	{"b2", WHITE, "c2", []string{"c2"}, []string{}, false},
	{"b2", WHITE, "c2", []string{}, []string{"c2"}, true}, // Capture
	{"b2", BLACK, "c2", []string{}, []string{"c2"}, false},
	{"b2", BLACK, "c2", []string{"c2"}, []string{}, true}, // Capture
	{"b2", WHITE, "h2", []string{}, []string{"c2"}, true},
	{"b2", WHITE, "h2", []string{"a2"}, []string{"c2"}, false},
	{"f5", WHITE, "f8", []string{"f7"}, []string{}, true},
	{"f5", WHITE, "f8", []string{"f7"}, []string{"f2"}, false},
}

func TestRookMovement(t *testing.T) {
	for _, tc := range rookMovementTestCases {
		r := NewRook(tc.color)
		b := NewBoard()
		mustPlace(t, b, r, tc.start)

		for _, pos := range tc.whitePositions {
			mustPlace(t, b, NewRook(WHITE), pos)
		}
		for _, pos := range tc.blackPositions {
			mustPlace(t, b, NewRook(BLACK), pos)
		}
		got := r.IsLegalMove(mustPosition(t, tc.dest))
		if got != tc.want {
			t.Errorf("Case %v failed", tc)
		}
	}
}

var bishopMovementTestCases = []struct {
	start          string
	color          Color
	dest           string
	whitePositions []string
	blackPositions []string
	want           bool
}{
	{"b2", WHITE, "b2", []string{}, []string{}, true},
	{"b2", WHITE, "c3", []string{}, []string{}, true},
	{"b2", WHITE, "c1", []string{}, []string{}, true},
	{"b2", WHITE, "a1", []string{}, []string{}, true},
	{"b2", WHITE, "b1", []string{}, []string{}, false},
	{"b2", WHITE, "b3", []string{}, []string{}, false},
	{"b2", WHITE, "a2", []string{}, []string{}, false},
	{"b2", WHITE, "c2", []string{}, []string{}, false},
	{"b2", WHITE, "h8", []string{}, []string{}, true},
	{"b2", WHITE, "h6", []string{}, []string{}, false},
	{"b2", WHITE, "e3", []string{}, []string{}, false},
	{"b2", WHITE, "d8", []string{}, []string{}, true}, // Wraparound diagonal
	{"b2", WHITE, "e5", []string{"e5"}, []string{}, false},
	{"b2", WHITE, "e5", []string{}, []string{"e5"}, true}, // Capture
	{"b2", BLACK, "e5", []string{}, []string{"e5"}, false},
	{"b2", BLACK, "e5", []string{"e5"}, []string{}, true}, // Capture
	{"b2", WHITE, "d4", []string{"c3"}, []string{}, true},
	{"b2", WHITE, "d4", []string{"c3"}, []string{"a1"}, false},
	{"b2", WHITE, "e7", []string{"d8"}, []string{}, true},
	{"b2", WHITE, "e7", []string{"d8"}, []string{"h4"}, false},
}

func TestBishopMovement(t *testing.T) {
	for _, tc := range bishopMovementTestCases {
		r := NewBishop(tc.color)
		b := NewBoard()
		mustPlace(t, b, r, tc.start)

		for _, pos := range tc.whitePositions {
			mustPlace(t, b, NewBishop(WHITE), pos)
		}
		for _, pos := range tc.blackPositions {
			mustPlace(t, b, NewBishop(BLACK), pos)
		}
		got := r.IsLegalMove(mustPosition(t, tc.dest))
		if got != tc.want {
			t.Errorf("Case %v failed", tc)
		}
	}
}
