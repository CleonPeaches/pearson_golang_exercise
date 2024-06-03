package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	quizHandler "quizApp/handlers"
	"quizApp/types"
)

func main() {
	// Define HTTP routes.
	http.HandleFunc("/quiz/list", quizHandler.ListQuizzesHandler)
	http.HandleFunc("/quiz/get", quizHandler.GetQuizByTitleHandler)
	http.HandleFunc("/questions", quizHandler.ListQuestionsByTopicHandler)

	// Start server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("could not start server: %s\n", err)
		}
	}()

	runCLILoop()
}

func runCLILoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter a topic name (or type 'exit' to quit):")
		topic, _ := reader.ReadString('\n')
		topic = strings.TrimSpace(topic)

		if topic == "exit" {
			break
		}

		if topic == "" {
			fmt.Println("Topic cannot be empty. Please enter a valid topic name.")
			continue
		}

		questions, err := fetchQuestionsByTopic(topic)
		if err != nil {
			fmt.Println("Error fetching questions:", err)
			continue
		}

		if len(questions) == 0 {
			fmt.Println("No questions found for the topic:", topic)
			continue
		}

		for _, question := range questions {
			fmt.Println(question.Text)
			for i, option := range question.Options {
				fmt.Printf("%d. %s\n", i+1, option)
			}

			fmt.Print("Enter your choice: ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)

			// Validate input
			if choiceNum, err := strconv.Atoi(choice); err == nil && choiceNum > 0 && choiceNum <= len(question.Options) {
				if question.Options[choiceNum-1] == question.CorrectAnswer {
					fmt.Println("Correct!")
				} else {
					fmt.Println("Incorrect. The correct answer is:", question.CorrectAnswer)
				}
			} else {
				fmt.Println("Invalid choice. Please enter a number between 1 and", len(question.Options))
			}
		}
	}
}

func fetchQuestionsByTopic(topic string) ([]types.Question, error) {
	resp, err := http.Get("http://localhost:8080/questions?topic=" + topic)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("No questions were found with that given keyword, status code: %d", resp.StatusCode)
	}

	var questions []types.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}
