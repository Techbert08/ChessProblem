package main

import (
	"fmt"
	"math/rand"

	"github.com/Techbert08/ChessProblem/internal"
)

// coin is an interface for a random (or not) coin
type coin interface {
	// Toss returns true for heads and false for tails.
	Toss() bool
}

// realCoin is a coin that is actually random, returning true or false
// with equal probability.
type realCoin struct{}

func (r realCoin) Toss() bool {
	return rand.Intn(2) == 1
}

// twoDice is an interface for a random (or not) set of dice.
type twoDice interface {
	// Roll returns the sum of two six sided die rolls
	Roll() int
}

// realDice is an implementation that returns random results.
type realDice struct{}

func (d realDice) Roll() int {
	// Intn returns numbers from zero to 5, so add one per die
	return rand.Intn(6) + rand.Intn(6) + 2
}

// evaluateProblem runs the stated problem, emitting log statements to the
// returned slice of strings.  On error the program terminates, but logs emitted
// so far are in the output slice.
func evaluateProblem(c coin, d twoDice, numMoves int) ([]string, error) {
	out := make([]string, 0)
	board := internal.NewBoard()
	rook := internal.NewRook(internal.BLACK)
	if err := board.PlacePiece(rook, "f6"); err != nil {
		return out, err
	}
	bishop := internal.NewBishop(internal.WHITE)
	if err := board.PlacePiece(bishop, "c3"); err != nil {
		return out, err
	}
	for i := 0; i < numMoves; i++ {
		roll := d.Roll()
		if c.Toss() {
			out = append(out, fmt.Sprintf("Heads, rolled %v", roll))
			if err := board.MovePiece(rook, rook.GetPosition().Move(0, roll)); err != nil {
				return out, err
			}
		} else {
			out = append(out, fmt.Sprintf("Tails, rolled %v", roll))
			if err := board.MovePiece(rook, rook.GetPosition().Move(roll, 0)); err != nil {
				return out, err
			}
		}
		out = append(out, fmt.Sprint(rook))
		if bishop.GetPosition() == nil {
			return append(out, "Rook takes bishop, Black wins"), nil
		}
		// Rook should never be nil, panic is fine if it is.
		if bishop.IsLegalMove(*rook.GetPosition()) {
			return append(out, "Bishop can take rook, White wins"), nil
		}
	}
	return append(out, "Rook escapes, Black wins"), nil
}

func main() {
	logs, err := evaluateProblem(&realCoin{}, &realDice{}, 15)
	for _, l := range logs {
		fmt.Println(l)
	}
	if err != nil {
		fmt.Println("Terminated with error: ", err)
	}
}
