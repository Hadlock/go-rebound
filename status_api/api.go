package status_api

import (
	"net/http"
	"fmt"
)

func dockerContainerListHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello, world")
}

func Handlers () *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/docker/containers", dockerContainerListHandler)
	r.Handle("/", http.FileServer(http.Dir("static")))

	return r
}
