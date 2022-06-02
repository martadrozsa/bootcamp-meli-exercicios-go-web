package main

import (
	"net/http"
)
// w http.ResponseWriter -->
// req *http.Request --> ponteiro

func helloHandler1(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`<h1>Olá</h1>`))
}

func main() {
	http.HandleFunc("/hello", helloHandler1)
	http.ListenAndServe("localhost:8080", nil)
}
