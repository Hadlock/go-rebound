package main

import (
	"github.com/hadlock/go-rebound/status_api"
	"net/http"
	"log"
	"bufio"
)

func main() {
	// redirect stdout into syslog
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			log.Debug("%s", line)
		}
	}()
	
	log.Debug("Num of CPU's: %d", cpus)
	
	// le web server
	log.Fatal(http.ListenAndServe(":8080", status_api.Handlers()))
}
