package main

import (
	"github.com/hadlock/go-rebound/status_api"
	"net/http"
    "log"
	)


func main() {
	log.Println("Hello World")
	log.Fatal(http.ListenAndServe(":8080", status_api.Handlers()))
	
}
