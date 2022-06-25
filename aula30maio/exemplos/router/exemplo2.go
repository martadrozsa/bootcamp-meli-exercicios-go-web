package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Router
Permite definir várias rotas e depois decidir onde cada uma delas é processada
*/

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Route",
	})
}

func GophersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Hello gophers",
	})
}

func GetGophersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Get gophers",
	})
}

func GetInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Get info gophers",
	})
}

func AboutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "About gophers",
	})
}

func main() {
	//rota padrão
	/*
		router := gin.Default()

		//toda vez que chamamos GET e passamos uma rota, definimos um novo endpoint
		router.GET("/", RootHandler)
		router.GET("/gophers", GophersHandler)
		router.GET("/gophers/get", GetGophersHandler)
		router.GET("/gophers/info", GetInfoHandler)
		router.GET("/about", AboutHandler)
		router.Run(":8080")
	*/

	server := gin.Default()
	server.GET("/", RootHandler)

	//Agrupamento de endpoints

	//agora podemos atender solicitações para /gophers || /gophers/get || /gophers/info || /gophers/info de forma mais elegante
	gopher := server.Group("/gophers")
	{
		gopher.GET("/", GophersHandler)
		gopher.GET("/get", GetGophersHandler)
		gopher.GET("/info", GetInfoHandler)
	}
	server.GET("/about", AboutHandler)
	server.Run(":8080")
}
