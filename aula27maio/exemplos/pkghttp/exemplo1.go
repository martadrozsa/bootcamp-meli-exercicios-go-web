package main

import (
	"fmt"
	"net/http"
)

func helloHandler1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Olá\n")
}

func main()  {
	http.HandleFunc("/hello", helloHandler1)
	http.ListenAndServe(":8080", nil)
}