package main

						import (
							"net/http"
							"log"
							"encoding/json"
						)

						func main() {
							http.HandleFunc("/hello", GreetingHandler)
							log.Fatal(http.ListenAndServe(":8080", nil))
						}
						

						func GreetingHandler(w http.ResponseWriter, r *http.Request)  {
							w.Header().Set("Content-Type", "application/json")
							data := map[string]string{
								"response": "hello world",
							}
							json.NewEncoder(w).Encode(data)
							
						}

					