package game

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindGameHandler(c *gin.Context) {

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

	gameResponse := GameResponse{
		ID:         game.ID,
		MaxGuesses: game.MaxGuesses,
		Players:    GetPlayerNames(game.Players),
	}

	c.JSON(200, gameResponse)

}
