package main

import (
	"math/rand/v2"
	"strconv"

	"github.com/gin-gonic/gin"
)

var goal = -1

func main() {

	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.POST("/goal", handleSetGoal)
	router.GET("/goal", handleGetGoal)
	router.GET("/goal/:guess", handleMakeGuess)

	router.Run("localhost:8080")

}

func handleSetGoal(c *gin.Context) {
	goal = rand.IntN(100)

	c.Status(200)
}

func handleGetGoal(c *gin.Context) {

	if goal == -1 {
		c.JSON(400, gin.H{
			"error": "goal not set",
		})
		return
	}

	c.JSON(200, gin.H{
		"goal": goal,
	})
}

func handleMakeGuess(c *gin.Context) {
	if goal == -1 {
		c.JSON(400, gin.H{
			"error": "goal not set",
		})
		return
	}

	guess := c.Param("guess")
	if guess == "" {
		c.JSON(400, gin.H{
			"error": "guess not set",
		})
		return
	}

	guessInt, err := strconv.Atoi(guess)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "guess not a number",
		})
		return
	}

	if guessInt < 0 || guessInt > 100 {
		c.JSON(400, gin.H{
			"error": "guess out of range",
		})
		return
	}

	if guessInt < goal {
		c.JSON(200, gin.H{
			"message": "guess too low",
		})
		return
	}

	if guessInt > goal {
		c.JSON(200, gin.H{
			"message": "guess too high",
		})
		return
	}

	if guessInt == goal {

		// reset goal
		goal = -1

		c.JSON(200, gin.H{
			"message": "guess correct",
		})
		return
	}
}
