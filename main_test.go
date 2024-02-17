package main

import (
	"reflect"
	"testing"
)

// loadedCoin is a fake coin that always tosses outcome.
type loadedCoin struct {
	Outcome bool
}

func (h *loadedCoin) Toss() bool {
	return h.Outcome
}

// loadedDice is a fake set of dice that always rolls outcome.
type loadedDice struct {
	Outcome int
}

func (s *loadedDice) Roll() int {
	return s.Outcome
}

var testCases = []struct {
	coin bool
	dice int
	want []string
}{
	{true, 2, []string{"Rolled 2", "Flipped heads", "Black wins"}},
}

func TestEvaluateProblemSingleIteration(t *testing.T) {
	for _, tc := range testCases {
		got, err := evaluateProblem(&loadedCoin{Outcome: tc.coin}, &loadedDice{Outcome: tc.dice}, 1)
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
	if got[len(got)-1] != "Black wins" {
		t.Errorf("evaluateProblem did not terminate with a winner, messages were %v", got)
	}
}
