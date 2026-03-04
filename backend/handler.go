package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handleTopology(w http.ResponseWriter, _ *http.Request) {
	topo := getTopology()
	writeJSON(w, http.StatusOK, topo)
}

func handleNetworks(w http.ResponseWriter, _ *http.Request) {
	topo := getTopology()
	writeJSON(w, http.StatusOK, topo.Networks)
}

func handleContainers(w http.ResponseWriter, _ *http.Request) {
	topo := getTopology()
	writeJSON(w, http.StatusOK, topo.Containers)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error encoding response: %v", err)
	}
}
