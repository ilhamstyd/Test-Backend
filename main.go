package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ID int `json:"id"`
}

type Response struct {
	Position string `json:"position"`
}

func main() {
	r := gin.Default()

	r.POST("/position", func(c *gin.Context) {
		// Get container ID
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		position := getPosition(request.ID)

		c.JSON(http.StatusOK, Response{Position: position})
	})

	r.Run("localhost:5000")
}

func getPosition(containerID int) string {

	position := "REJECT"

	containerIDStr := strconv.Itoa(containerID)

	if containerIDStr[:3] == "313" {
		position = "CENTER"
	}

	if containerIDStr[:3] == "233" && containerIDStr[4:] == containerIDStr[4:] {
		position = "RIGHT"
	}

	if containerIDStr[:3] == "136" && isPrime(containerID) {
		if isPrime(containerID) {
			position = "LEFT"
		}
	}

	return position
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
