package status_api

import (
	"os"
	"net"
	"net/http"
	"log"
	"fmt"
	"bufio"
)

var dockerSockPath string = os.Getenv("DOCKER_SOCKET")

func dockerContainerListHandler (w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if dockerSockPath == "" {
		dockerSockPath = "/run/docker.sock"
	}

	conn, err := net.Dial("unix", dockerSockPath)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Fprintf(conn, "GET /containers/json HTTP/1.0\r\n\r\n")
		if dockerResponse, err := bufio.NewReader(conn).ReadSlice('\n'); err != nil {
			log.Print(err)
		} else {
			w.Write(dockerResponse)
		}
	}
}

func Handlers () *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/docker/containers", dockerContainerListHandler)
	r.Handle("/", http.FileServer(http.Dir("static")))

	return r
}
