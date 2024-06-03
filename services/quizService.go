package quizService

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"quizApp/types"
)

var quizzes []types.Quiz

func init() {
	data, err := os.ReadFile("./demoData.json")
	if err != nil {
		fmt.Println("failed to read JSON file: ", err)
	}

	if err := json.Unmarshal(data, &quizzes); err != nil {
		fmt.Println("failed to unmarshal JSON data: ", err)
	} else {
		fmt.Println("Successfully loaded JSON data.")
	}
}

func ListQuizzes() []types.Quiz {
	return quizzes
}

func GetQuizByTitle(quizTitle string) (*types.Quiz, error) {
	for _, quiz := range quizzes {
		if quiz.Title == quizTitle {
			return &quiz, nil
		}
	}
	return nil, fmt.Errorf("quiz not found with title %s", quizTitle)
}

func ListQuestionsByTopic(topic string) ([]types.Question, error) {
	var filteredQuestions []types.Question
	for _, quiz := range quizzes {
		for _, question := range quiz.Questions {
			if strings.Contains(strings.ToLower(question.Text), strings.ToLower(topic)) {
				filteredQuestions = append(filteredQuestions, question)
			}
		}
	}
	if len(filteredQuestions) > 0 {
		return filteredQuestions, nil
	}
	return nil, fmt.Errorf("no questions found containing %s", topic)
}
