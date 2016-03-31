package status_api

import "net/http"

func Handlers () *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.Dir("static")))

	return r
}
