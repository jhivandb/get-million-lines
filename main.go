package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting to write million-line JSON file...")
	err := writeMillionLineJson()
	if err != nil {
		log.Fatalf("Error writing million-line JSON: %v", err)
	}
	log.Println("Finished writing million-line JSON file.")

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

	fileName := "example.json"
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	return data
}

func writeMillionLineJson() error {
	fileName := "example.json"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("[\n")
	if err != nil {
		return err
	}

	for i := 0; i < 1_000_000; i++ {
		line := fmt.Sprintf(`  {"id": %d, "value": "line %d"}`, i, i)
		if i < 999_999 {
			line += ",\n"
		} else {
			line += "\n"
		}
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString("]\n")
	return err
}
