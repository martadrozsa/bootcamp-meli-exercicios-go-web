package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

/*
Gin
Recebe um request -> analisa a rota -> se tiver um middleware especificado para a rota, ele é passado -> func route hanlder  -> outro middleware -> response (ok ou erro)
Request -> Router -> Middleware -> Router Handler -> Middleware -> Response
*/
// gin.Context temos acesso ao json (marshal e unmarshal)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	gin.SetMode("release")

	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}

	gin.DefaultWriter = io.MultiWriter(f)

	//cria um router com o gin
	router := gin.Default()

	router.GET("/hello", helloHandler)

	router.Run()
}