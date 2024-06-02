package quizService

import (
	"encoding/json"
	"fmt"
	"os"

	"quizApp/types"
)

var quizzes []types.Quiz

func init() {
	data, err := os.ReadFile("./demoData.json")
	if err != nil {
		fmt.Println("failed to read JSON file: %v", err)
	}

	if err := json.Unmarshal(data, &quizzes); err != nil {
		fmt.Println("failed to unmarshal JSON data: %v", err)
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
