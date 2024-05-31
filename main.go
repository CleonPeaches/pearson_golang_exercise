package main

import (
	"log"
	"net/http"

	quizHandler "quizApp/handlers"
)

func main() {

	// Define HTTP routes.
	http.HandleFunc("/quiz/list", quizHandler.ListQuizzesHandler)
	http.HandleFunc("/quiz/get", quizHandler.GetQuizByTitleHandler)

	// Start the server.
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
