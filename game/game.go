// game/game.go
package game

import (
	"slices"

	"github.com/google/uuid"
)

type Game struct {
	ID         uuid.UUID `json:"id"`
	Goal       int       `json:"goal"`
	Guesses    []int     `json:"guesses"`
	MaxGuesses int       `json:"max_guesses"`
}

func (g Game) IsOver() bool {
	return len(g.Guesses) >= g.MaxGuesses
}

func (g Game) IsWon() bool {
	return g.Goal == g.Guesses[len(g.Guesses)-1]
}

func (g Game) IsLost() bool {
	return len(g.Guesses) >= g.MaxGuesses && g.Goal != g.Guesses[len(g.Guesses)-1]
}

func (g Game) IsValidGuess(guess int) bool {
	return guess >= 1 && guess <= 100
}

func (g Game) IsGuessAlreadyMade(guess int) bool {
	return slices.Contains(g.Guesses, guess)
}

func (g *Game) AppendGuess(guess int) {
	g.Guesses = append(g.Guesses, guess)
}

func (g Game) GetRemainingGuesses() int {
	return g.MaxGuesses - len(g.Guesses)
}

func (g Game) IsGuessTooHigh(guess int) bool {
	return guess > g.Goal
}

type GameResponse struct {
	ID         uuid.UUID `json:"id"`
	Guesses    []int     `json:"guesses"`
	MaxGuesses int       `json:"max_guesses"`
}

var games []Game
