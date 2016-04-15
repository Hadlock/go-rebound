package main

import (
	"github.com/hadlock/go-rebound/status_api"
	"net/http"
    "log"
	)


func main() {
	log.SetFlags(0)
	log.Println("go-rebound v0.0.1 Listening on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", status_api.Handlers()))
	
}
