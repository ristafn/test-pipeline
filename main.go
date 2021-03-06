package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("hello from test")
	})

	http.HandleFunc("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("hello from test")
	})

	http.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("hello from pong")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
