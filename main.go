package main

import (
	"github.com/hadlock/go-rebound/status_api"
	"github.com/hadlock/go-rebound/log"
	"net/http"
	"bufio"			//for logs to container
	"os"			//for logs to container
	
	
	
	
)

func main() {
	// redirect stdout into syslog
	//oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w
	
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			log.Info("%s", line)
		}
	}()
	
	log.Info("Hello World")
	
	// le web server
	log.Fatal(http.ListenAndServe(":8080", status_api.Handlers()))
}
