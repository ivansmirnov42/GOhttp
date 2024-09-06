package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
}

var message string

func send_message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var reqBody requestBody
	json.NewDecoder(r.Body).Decode(&reqBody)
	message = reqBody.Message
}

func receive_message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v", message)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/post", send_message).Methods("POST")
	router.HandleFunc("/get", receive_message).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
	fmt.Println(message)
}
