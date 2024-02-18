package internal

import (
	"errors"
	"reflect"
	"testing"
)

var newPositionTestCases = []struct {
	p         string
	wantError error
}{
	{"a1", nil},
	{"a1b", errors.New("NewPosition: invalid string. expected len(2), got a1b")},
	{"h-1", errors.New("NewPosition: invalid string. expected len(2), got h-1")},
	{"A1", errors.New("NewPosition: first file character should be a-h, got A")},
	{"i7", errors.New("NewPosition: first file character should be a-h, got i")},
	{"h7", nil},
	{"c8", nil},
	{"d9", errors.New("NewPosition: second rank character should be 1-8, got 9")},
}

func TestNewPosition(t *testing.T) {
	for _, tc := range newPositionTestCases {
		got, err := NewPosition(tc.p)
		// Although generally not recommended, assert on the actual error strings to confirm
		// structure matches what's expected.
		if !reflect.DeepEqual(err, tc.wantError) {
			t.Errorf("NewPosition(%v) returned err %v, wanted %v", tc.p, err, tc.wantError)
		}
		// Also check that printing the Position matches the input on non-error cases.
		if tc.wantError != nil && got != nil {
			toString := got.String()
			if toString != tc.p {
				t.Errorf("NewPosition(%v).ToString() returned %v, wanted %v", tc.p, toString, tc.p)
			}
		}
	}
}

var moveTestCases = []struct {
	p    string
	f    int
	r    int
	want string
}{
	{"a1", 0, 0, "a1"},
	{"a1", 1, 0, "b1"},
	{"a1", -1, 0, "h1"},
	{"a1", -2, 0, "g1"},
	{"a1", 0, 1, "a2"},
	{"a1", 0, -1, "a8"},
	{"a1", 0, -2, "a7"},
	{"h8", 0, 0, "h8"},
	{"h8", 1, 0, "a8"},
	{"h8", -1, 0, "g8"},
	{"h8", -2, 0, "f8"},
	{"h8", 0, 1, "h1"},
	{"h8", 0, -1, "h7"},
	{"h8", 0, -2, "h6"},
	{"c4", 2, -1, "e3"},
	{"a1", 8, 8, "a1"},
	{"a1", -8, -8, "a1"},
}

func TestMove(t *testing.T) {
	for _, tc := range moveTestCases {
		pos, err := NewPosition(tc.p)
		if err != nil {
			t.Errorf("Bad test case setup, start position %v returned err %v", tc.p, err)
		}
		got := pos.Move(tc.f, tc.r).String()
		if got != tc.want {
			t.Errorf("NewPosition(%v).Move(%v,%v).String() = %v, wanted %v", tc.p, tc.f, tc.r, got, tc.want)
		}
	}
}
