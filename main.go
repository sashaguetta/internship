package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", GreetingHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"response": "hello world",
	}
	json.NewEncoder(w).Encode(data)

}
