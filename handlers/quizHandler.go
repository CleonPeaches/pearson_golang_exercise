package handlers

import (
	"encoding/json"
	"net/http"
	quizService "quizApp/services"
)

func GetQuizByTitleHandler(w http.ResponseWriter, r *http.Request) {
	quizTitle := r.URL.Query().Get("title")
	if quizTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	quiz, err := quizService.GetQuizByTitle(quizTitle)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}

func ListQuizzesHandler(w http.ResponseWriter, r *http.Request) {
	quizzes := quizService.ListQuizzes()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quizzes)
}

func ListQuestionsByTopicHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	questions, err := quizService.ListQuestionsByTopic(topic)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
