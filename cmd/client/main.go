package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":80", http.FileServer(http.Dir("./assets")))
	log.Fatalln(err)
}
