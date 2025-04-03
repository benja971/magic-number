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
	router.GET("/game/:id", game.FindGameHandler)
	router.POST("/game/:id/guess", game.MakeGuessHandler)

	router.Run("localhost:8080")

}
