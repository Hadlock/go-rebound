package status_api

import (
	"log"
	"os"
	"net"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

var dockerSockPath string = os.Getenv("DOCKER_SOCKET")

func newFakeDialer(path string) func(string, string) (net.Conn, error) {
	return func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", path)
	}
}

func newSocketClient(path string) (*http.Client) {
	tr := &http.Transport{
		Dial: newFakeDialer(path),
	}

	return &http.Client{Transport: tr}	
}

func dockerContainerListHandler (w http.ResponseWriter, r *http.Request) {
	
	if dockerSockPath == "" {
		dockerSockPath = "/run/docker.sock"
	}
	
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	dockerClient := newSocketClient(dockerSockPath)

	if resp, err := dockerClient.Get("http:/containers/json"); err != nil {
		log.Fatal(err)
	} else {
		if respBody, err := ioutil.ReadAll(resp.Body); err != nil {
			log.Fatal(err)
		}	else {
			respBuffer := bytes.NewBuffer(respBody)
			respString := respBuffer.String()
			fmt.Println(respString)
			fmt.Fprintln(w, respString)
		}	
	}
}

func Handlers () *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/docker/containers", dockerContainerListHandler)
	r.Handle("/", http.FileServer(http.Dir("static")))

	return r
}
