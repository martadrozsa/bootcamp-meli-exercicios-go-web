package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin.gincontext
permite passar variáveis entre middleware, assim, como os cabeçalhos, parãmetros de query string, médoto entre outras variáveis de requisição.
Ele é responsável por validar o JSON de uma solicitação e gerar uma resposta JSON
O framework Gin usa parte do pacote nativo Golang para lidar com requisições, assim como http.Request e ResponseWriter.
*/

func handler(ctx *gin.Context) {
	body := ctx.Request.Body
	header := ctx.Request.Header
	method := ctx.Request.Method

	fmt.Println("Eu recebi algo!")
	fmt.Printf("\tMétodo: %s\n", method)
	fmt.Printf("\tConteúdo do cabeçalho: \n")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}

	fmt.Printf("\t body é um io.ReadCloser: (%s), e para trabalhar com ele teremos que ler depois\n", body)
	fmt.Println("Yes")
	ctx.String(http.StatusOK, "Eu recebi")

}

func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.Run()
}
