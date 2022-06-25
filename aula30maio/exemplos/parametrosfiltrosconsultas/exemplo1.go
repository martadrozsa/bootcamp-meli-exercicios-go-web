package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//"base de dados" obde iremos consultar a informação
var employee = map[string]string{
	"644": "Employee A",
	"755": "Employee B",
	"777": "Employee C",
}

func RootPage(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Bem-vindo à empresa Gopher!")
}

func SearchEmployee(ctx *gin.Context) {
	employee, ok := employee[ctx.Param("id")]
	if ok {
		ctx.String(http.StatusOK, "Informação do empregado %s, nome: %s", ctx.Param("id"), employee)
	} else {
		ctx.String(404, "Informação do empregado não existe!")
	}
}

func main() {
	server := gin.Default()
	server.GET("/", RootPage)
	server.GET("/employee/id", SearchEmployee)
	server.Run(":8085")
}
