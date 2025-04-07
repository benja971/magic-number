package game

import (
	"example/guess_number/player"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GuessRequest struct {
	Guess  int    `json:"guess"`
	Player string `json:"player"`
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

	currentPlayer, found := player.FindPlayer(&game.Players, guessRequest.Player)

	if !found {
		c.JSON(404, gin.H{"error": "Player not found"})
		return
	}

	if !game.IsValidGuess(guessRequest.Guess) {
		c.JSON(400, gin.H{"error": "Guess must be between 1 and 100"})
		return
	}

	remainingGuesses := game.GetRemainingGuesses(currentPlayer)

	if game.HasPlayerWon(currentPlayer) {
		c.JSON(200, gin.H{"message": "You have already won!", "remaining_guesses": remainingGuesses})
		return
	}

	if currentPlayer.IsGuessAlreadyMade(guessRequest.Guess) {
		c.JSON(400, gin.H{"error": "Guess already made"})
		return
	}

	currentPlayer.AppendGuess(guessRequest.Guess)
	remainingGuesses = game.GetRemainingGuesses(currentPlayer)

	if game.IsCorrectGuess(guessRequest.Guess) {
		c.JSON(200, gin.H{"message": "Congratulations! You guessed the number!", "remaining_guesses": remainingGuesses})
		return
	}

	if game.HasPlayerLost(currentPlayer) {
		c.JSON(200, gin.H{"message": "Game over! You lost!", "remaining_guesses": remainingGuesses})
		return
	}

	if game.IsGuessTooHigh(guessRequest.Guess) {
		c.JSON(200, gin.H{"message": "Too high!", "remaining_guesses": remainingGuesses})
		return
	}

	c.JSON(200, gin.H{"message": "Too low!", "remaining_guesses": remainingGuesses})
}
