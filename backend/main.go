package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("GET /topology", handleTopology)
	mux.HandleFunc("GET /networks", handleNetworks)
	mux.HandleFunc("GET /containers", handleContainers)

	log.Println("PacketDeck backend listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
