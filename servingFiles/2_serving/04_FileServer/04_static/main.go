package main

import (
	"log"
	"net/http"
)

func main() {
	// log.Fatal Http.Error
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
