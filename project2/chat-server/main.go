package main

import (
	"encoding/json"
	//"time"
	"log"
	"net/http"
	"os"
)

type Message struct {
	//username string
	//timestamp time.Time
	Text string `json:"text"`
	ImagePath string 
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/messages", MessagesHandler)
	http.HandleFunc("/message", MessageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadMessages() []Message {
	var messages []Message
	b, err := os.ReadFile("messages.json")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	} else {
		err = json.Unmarshal(b, &messages)
		if err != nil {
			log.Fatal(err)
		}
	}
	return messages
}

func saveMessages(messages []Message) {
	b, err := json.MarshalIndent(messages, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("messages.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	messages := loadMessages()
	json.NewEncoder(w).Encode(messages)

}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req Message

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	messages := loadMessages()
	messages = append(messages, req)
	saveMessages(messages)

}
