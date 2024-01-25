package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", "0.0.0.0:80", "listen address")
	dir    = flag.String("dir", "./static", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on http://%s", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
