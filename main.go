package main

import (
	"github.com/ScriptRock/go-appstatus/status_api"
	"net/http"
)

func main() {
	panic(http.ListenAndServe(":8080", status_api.Handlers()))
}
