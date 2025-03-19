package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/ting-tinggg", notify)
	router.Run("localhost:8080")
}

func notify(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
