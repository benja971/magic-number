package main

import (
	"example/guess_number/game"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.POST("/game", game.CreateGameHandler)
	router.PUT("/game/:id/join", game.JoinGameHandler) // /game/:id/join?player=John_Doe
	router.GET("/game/:id", game.FindGameHandler)
	router.POST("/game/:id/guess", game.MakeGuessHandler)

	router.Run("localhost:8080")

}
