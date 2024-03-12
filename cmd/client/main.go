package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("./assets")))
	log.Fatalln(err)
}
