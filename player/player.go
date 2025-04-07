package player

import "slices"

type Player struct {
	Name    string `json:"name"`
	Guesses []int  `json:"guesses"`
}

type PlayerResponse struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type PlayerRequest struct {
	Name string `json:"name"`
}

func (p *Player) AppendGuess(guess int) {
	p.Guesses = append(p.Guesses, guess)
}

func (p Player) GetScore() int {
	return len(p.Guesses)
}

func (p Player) IsGuessAlreadyMade(guess int) bool {
	return slices.Contains(p.Guesses, guess)
}

