package main

import (
	"reflect"
	"strings"
	"testing"
)

// loadedCoin is a fake coin
type loadedCoin struct {
	// Outcome is the list of flips that will be returned.
	// Flipping more times than length of Outcome slice will panic.
	Outcome []bool
}

func (l *loadedCoin) Toss() bool {
	out, remainder := l.Outcome[0], l.Outcome[1:]
	l.Outcome = remainder
	return out
}

// loadedDice is a fake set of dice
type loadedDice struct {
	// Outcome is the list of rolls that will be returned.
	// Rolling more times than length of Outcome slice will panic.
	Outcome []int
}

func (l *loadedDice) Roll() int {
	out, remainder := l.Outcome[0], l.Outcome[1:]
	l.Outcome = remainder
	return out
}

var testCases = []struct {
	coin     []bool
	dice     []int
	numTurns int
	want     []string
}{
	{[]bool{true}, []int{2}, 1, []string{
		"Heads, rolled 2",
		"Black Rook at f8",
		"Bishop can take rook, White wins",
	}},
	{[]bool{false}, []int{3}, 1, []string{
		"Tails, rolled 3",
		"Black Rook at a6",
		"Rook escapes, Black wins",
	}},
	{[]bool{false}, []int{8}, 1, []string{
		"Tails, rolled 8",
		"Black Rook at f6",
		"Bishop can take rook, White wins",
	}},
	{[]bool{false, true}, []int{5, 5}, 2, []string{
		"Tails, rolled 5",
		"Black Rook at c6",
		"Heads, rolled 5",
		"Black Rook at c3",
		"Rook takes bishop, Black wins",
	}},
}

func TestEvaluateProblemDeterministic(t *testing.T) {
	for _, tc := range testCases {
		got, err := evaluateProblem(&loadedCoin{Outcome: tc.coin}, &loadedDice{Outcome: tc.dice}, tc.numTurns)
		if err != nil {
			t.Errorf("evaluateProblem(%v,%v,1) returned err: %v", tc.coin, tc.dice, err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("evaluateProblem(%v,%v,1) = %v, wanted %v", tc.coin, tc.dice, got, tc.want)
		}
	}
}

func TestEvaluateProblemCompletes(t *testing.T) {
	got, err := evaluateProblem(&realCoin{}, &realDice{}, 15)
	if err != nil {
		t.Errorf("evaluateProblem with random inputs returned error %v", err)
	}
	if !strings.Contains(got[len(got)-1], "wins") {
		t.Errorf("evaluateProblem did not terminate with a winner, messages were %v", got)
	}
}
