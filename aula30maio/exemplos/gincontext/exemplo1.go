package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Context is the most important part of gin.
It allows us to pass variables between middleware, manage the flow, validate the JSON of a request and render a JSON response for example.
*/

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}