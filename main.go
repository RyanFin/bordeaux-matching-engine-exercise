package main

import (
	"bordeaux-matching-engine-exercise/pkg/api"
	"log"
)

const (
	serverAddr = "0.0.0.0:8080"
)

func main() {

	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("cannot initialise server: %s", err.Error())
	}

	// main server
	err = server.Start(serverAddr)
	if err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
