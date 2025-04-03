package game

import (
	"fmt"
	"math/rand/v2"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateGameRequest struct {
	MaxGuesses int `json:"max_guesses"`
}

func CreateGameHandler(c *gin.Context) {

	body := CreateGameRequest{}

	// c.ShouldBindJSON(&body) is a method that binds the JSON body of the request to the struct passed as an argument.
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "message": "Unable to parse JSON body"})
		return
	}

	game := Game{
		ID:         uuid.New(),
		Goal:       rand.IntN(100) + 1, // Random number between 1 and 100
		Guesses:    []int{},
		MaxGuesses: body.MaxGuesses,
	}

	games = append(games, game)

	fmt.Println("Game with ID", game.ID, "created with goal", game.Goal)

	gameResponse := GameResponse{
		ID:         game.ID,
		Guesses:    game.Guesses,
		MaxGuesses: game.MaxGuesses,
	}

	c.JSON(200, gameResponse)
}
