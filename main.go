package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// GET example
	r.GET("/getData", func(c *gin.Context) {
		result := GetDb()
		c.JSON(http.StatusCreated, gin.H{"message": result})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "returning no data from the database"})
	})

	// POST example
	r.POST("/postData", func(ctx *gin.Context) {
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			// Handle error
		}

		var input map[string]interface{}
		err = json.Unmarshal(jsonData, &input)

		if err != nil {
			// Handle error
		}
		fmt.Println(input)
		message := SaveDb(input)

		ctx.JSON(http.StatusCreated, gin.H{"message": message})
	})
	// DELETE example

	r.Run() // listen and serve on 0.0.0.0:8080
}
