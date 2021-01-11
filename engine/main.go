package main

import (
	"log"
	"os"

	"github.com/quoeamaster/go_cicd_sample_project/service"
)

// main - entry point of the binary executable.
func main() {
	if len(os.Args) != 2 {
		log.Printf("expect to provide a message as argument~\n")
		os.Exit(1)
	}
	srv := new(service.EchoServer)
	jDoc, err := srv.Serve(os.Args[1])
	if err != nil {
		panic(err)
	}
	log.Printf("cool~ the returned response from echo-server:\n  %v\n", jDoc)
}
