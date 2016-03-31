package main

import (
	"github.com/ScriptRock/go-appstatus/status_api"
	"net/http"
	"log"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", status_api.Handlers()))
}
