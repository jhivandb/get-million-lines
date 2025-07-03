package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	r := http.DefaultServeMux
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Write(getMillionLineJson())
	})

	log.Println("Starting HTTP server on 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", r)
}

func getMillionLineJson() []byte {

	fileName := "response.json"
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	return data
}
