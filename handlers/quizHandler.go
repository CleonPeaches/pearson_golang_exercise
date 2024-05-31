package handlers

import (
	"encoding/json"
	"net/http"
	quizService "quizApp/services"
	"quizApp/types"
)

func GetQuizByTitleHandler(w http.ResponseWriter, r *http.Request) {
	// Extract quiz title from request parameters
	quizTitle := r.URL.Query().Get("title")
	if quizTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service method to retrieve the quiz by its title
	quiz, err := quizService.GetQuizByTitle(quizTitle)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}

// ListUsersHandler handles requests to list all users.
func ListQuizzesHandler(w http.ResponseWriter, r *http.Request) []types.Quiz {
	return quizService.ListQuizzes()
}
