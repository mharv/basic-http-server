package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		result := GetDb()
		c.JSON(http.StatusCreated, gin.H{"message": result["test"]})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
