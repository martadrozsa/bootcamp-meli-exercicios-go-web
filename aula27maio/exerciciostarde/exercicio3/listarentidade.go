package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Exercício 3 - Listar Entidade
Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista do tema escolhido.
1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos correspondentes.
2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
3. Crie uma handler para o endpoint chamada "GetAll".
4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.
*/

type produto struct {
	id          int64   `json:"id"`
	nome        string  `json:"nome"`
	cor         string  `json:"tipo"`
	preco       float64 `json:"quantidade"`
	estoque     bool    `json:"preco"`
	codigo      string  `json:"codigo"`
	publicacao  bool    `json:"publicacao"`
	dataCriacao string  `json:"dataCriacao"`
}

type produtos []produto

var prod produto

func listProdutos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Produtos": &prod})
}

func main() {

	r := gin.Default()

	r.GET("/produtos", listProdutos)

	r.Run()
}
