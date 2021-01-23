package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "d")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "c")
}

func main() {
	http.HandleFunc("/dog", http.HandlerFunc(d))
	http.HandleFunc("/cat", http.HandlerFunc(c))
	http.ListenAndServe(":8081", nil)
}
