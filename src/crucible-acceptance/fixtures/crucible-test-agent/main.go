package main

import (
	"crucible-acceptance/fixtures/crucible-test-agent/handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.Int("port", -1, "port the server listens on")

func main() {
	flag.Parse()
	if *port == -1 {
		log.Fatal("no explicit port specified")
	}

	crucibleVar := os.Getenv("CRUCIBLE")
	if crucibleVar == "" {
		log.Fatal("Crucible environment variable not set")
	}

	http.HandleFunc("/", handlers.Hello)
	http.HandleFunc("/whoami", handlers.Whoami)
	http.HandleFunc("/hostname", handlers.Hostname)
	http.HandleFunc("/mounts", handlers.Mounts)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}