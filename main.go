package main

import (
	"fmt"
	"math/rand"

	"github.com/Techbert08/ChessProblem/internal"
)

// coin is broken out for easier mocking.
// Toss returns true for heads and false for tails.
type coin interface {
	Toss() bool
}

// realCoin is a coin that is actually random, returning true or false
// with equal probability.
type realCoin struct{}

// Use a value receiver, as the struct is stateless
func (r realCoin) Toss() bool {
	return rand.Intn(2) == 1
}

// twoDice is broken out for easier mocking.
// Roll returns the sum of two six sided die rolls
type twoDice interface {
	Roll() int
}

// realDice is an implementation that returns random results.
type realDice struct{}

// Use a value receiver, as the struct is stateless
func (d realDice) Roll() int {
	// Intn returns numbers from zero to 5, so add one per die
	return rand.Intn(6) + rand.Intn(6) + 2
}

// evaluateProblem runs the stated problem, emitting log statements to the
// returned slice of strings.  On error the program terminates, but logs emitted
// so far are in the output slice.
func evaluateProblem(c coin, d twoDice, numMoves int) ([]string, error) {
	out := make([]string, 0)
	board := internal.Board{}
	rook := internal.Rook{C: internal.BLACK}
	if err := board.PlacePiece(&rook, "f6"); err != nil {
		return out, err
	}
	bishop := internal.Bishop{C: internal.WHITE}
	if err := board.PlacePiece(&bishop, "c3"); err != nil {
		return out, err
	}
	for i := 0; i < numMoves; i++ {
		roll := d.Roll()
		out = append(out, fmt.Sprintf("Rolled %v", roll))
		if c.Toss() {
			out = append(out, "Flipped heads")
		} else {
			out = append(out, "Flipped tails")
		}
	}
	out = append(out, "Black wins")
	return out, nil
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
