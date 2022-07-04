package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("hello from test")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
