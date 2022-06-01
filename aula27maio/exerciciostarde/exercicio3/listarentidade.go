package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

/*
Exercício 3 - Listar Entidade
Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista do tema escolhido.
1. Dentro do "main.go", crie uma estrutura, conforme o tema, com os campos correspondentes.
2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
3. Crie uma handler para o endpoint chamada "GetAll".
4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.
*/

type produto struct {
	Id          int64   `json:"id"`
	Nome        string  `json:"nome"`
	Cor         string  `json:"cor"`
	Preco       float64 `json:"preco"`
	Estoque     bool    `json:"estoque"`
	Codigo      string  `json:"codigo"`
	Publicacao  bool    `json:"publicacao"`
	DataCriacao string  `json:"dataCriacao"`
}

func produtoHandler(c *gin.Context) {

	jsonText, err := os.ReadFile("./products.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var produtos []produto
	err = json.Unmarshal(jsonText, &produtos)
	if err != nil {
		return
	}
	c.JSON(200, produtos)
}

func main() {

	r := gin.Default()
	r.GET("/produtos", produtoHandler)
	r.Run()
}
