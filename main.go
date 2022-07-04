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
	r.GET("/data", func(ctx *gin.Context) {
		result := GetDb()
		ctx.JSON(http.StatusCreated, gin.H{"message": result})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{"message": "returning no data from the database \n"})
	})

	// POST example
	r.POST("/data", func(ctx *gin.Context) {
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

	r.DELETE("/data", func(ctx *gin.Context) {

		message := DeleteDb()

		ctx.JSON(http.StatusCreated, gin.H{"message": message})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
