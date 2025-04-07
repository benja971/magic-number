package game

import (
	"example/guess_number/player"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func JoinGameHandler(c *gin.Context) {
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

	if game.IsStarted() {
		c.JSON(400, gin.H{"error": "Game already started"})
		return
	}

	playerName := c.Query("player")

	if playerName == "" {
		c.JSON(400, gin.H{"error": "Player name is required"})
		return
	}

	if game.IsNameTaken(playerName) {
		c.JSON(400, gin.H{"error": "Player name already taken"})
		return
	}

	game.Players = append(game.Players, player.Player{Name: playerName, Guesses: []int{}})

	c.JSON(200, gin.H{"message": "Joined game successfully"})
}
