//go:build client
// +build client

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed assets
var f embed.FS

func main() {
	assets, err := fs.Sub(f, "assets")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(assets)))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
