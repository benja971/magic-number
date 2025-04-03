package game

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GuessRequest struct {
	Guess int `json:"guess"`
}

type GuessResponse struct {
	Message          string `json:"message"`
	RemainingGuesses int    `json:"remaining_guesses"`
}

func MakeGuessHandler(c *gin.Context) {
	gameID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid game ID format"})
		return
	}

	game, found := FindGameById(games, gameID)

	if !found {
		c.JSON(404, gin.H{"error": "Game not found"})
		return
	}

	var guessRequest GuessRequest
	if err := c.ShouldBindJSON(&guessRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid guess format"})
		return
	}

	if game.IsOver() {
		c.JSON(400, gin.H{"error": "Game is over"})
		return
	}

	if !game.IsValidGuess(guessRequest.Guess) {
		c.JSON(400, gin.H{"error": "Guess must be between 1 and 100"})
		return
	}

	if game.IsGuessAlreadyMade(guessRequest.Guess) {
		c.JSON(400, gin.H{"error": "Guess already made"})
		return
	}

	game.AppendGuess(guessRequest.Guess)
	fmt.Println("Guesses so far:", game.Guesses)
	fmt.Println("Guess made:", guessRequest.Guess)
	remainingGuesses := game.MaxGuesses - len(game.Guesses)

	if game.IsWon() {
		c.JSON(200, gin.H{"message": "Congratulations! You guessed the number!", "remaining_guesses": remainingGuesses})
		games, _ = DeleteGame(games, game.ID)
		return
	}

	if game.IsLost() {
		c.JSON(200, gin.H{"message": "Game over! You lost!", "remaining_guesses": 0})
		games, _ = DeleteGame(games, game.ID)
		return
	}

	if game.IsGuessTooHigh(guessRequest.Guess) {
		c.JSON(200, gin.H{"message": "Too high!", "remaining_guesses": remainingGuesses})
		return
	}

	c.JSON(200, gin.H{"message": "Too low!", "remaining_guesses": remainingGuesses})
}
