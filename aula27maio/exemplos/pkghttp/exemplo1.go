package main

import (
	"fmt"
	"net/http"
)

func helloHandler1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>Olá</h1>\n")
}

func main() {
	http.HandleFunc("/hello1", helloHandler1)
	http.ListenAndServe("localhost:8080", nil)
}
