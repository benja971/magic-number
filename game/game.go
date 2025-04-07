// game/game.go
package game

import (
	"example/guess_number/player"

	"github.com/google/uuid"
)

type Game struct {
	ID         uuid.UUID       `json:"id"`
	Goal       int             `json:"goal"`
	MaxGuesses int             `json:"max_guesses"`
	Players    []player.Player `json:"players"`
	StartedAt  int64           `json:"started_at"`
}

func (g Game) IsCorrectGuess(guess int) bool {
	return guess == g.Goal
}

func (g Game) HasPlayerLost(player *player.Player) bool {
	return len(player.Guesses) >= g.MaxGuesses
}

func (g Game) HasPlayerWon(player *player.Player) bool {
	return player.IsGuessAlreadyMade(g.Goal)
}

func (g Game) IsValidGuess(guess int) bool {
	return guess >= 1 && guess <= 100
}

func (g Game) GetRemainingGuesses(player *player.Player) int {
	return g.MaxGuesses - len(player.Guesses)
}

func (g Game) IsGuessTooHigh(guess int) bool {
	return guess > g.Goal
}

func (g Game) IsStarted() bool {
	return g.StartedAt != 0
}

func (g Game) IsNameTaken(name string) bool {
	for _, p := range g.Players {
		if p.Name == name {
			return true
		}
	}
	return false
}

type GameResponse struct {
	ID         uuid.UUID `json:"id"`
	MaxGuesses int       `json:"max_guesses"`
	Players    []string  `json:"players"`
}

var games []Game
