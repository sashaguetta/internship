package main

import (
	"encoding/json"
	"log"
	"net/http"
)
	type MultiplyRequest struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	type MultiplyResponse struct {
		Result int `json:"result"`
	}

func main() {
	http.HandleFunc("/multiply", OperationsHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8080", nil))

	
}

func OperationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
    	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req MultiplyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	response := MultiplyResponse{
		Result: req.A * req.B,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

