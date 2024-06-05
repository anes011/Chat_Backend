package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anes011/chat/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/uploadFile", handlers.HandleUpload).Methods("POST")
	router.HandleFunc("/createUser", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	fmt.Println("Server running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
