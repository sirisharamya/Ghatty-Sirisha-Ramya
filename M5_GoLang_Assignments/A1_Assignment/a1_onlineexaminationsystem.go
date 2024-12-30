package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Question struct
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

// Function to take the quiz
func takeQuiz(questions []Question) (int, error) {
	var score int
	for i, q := range questions {
		// Display the question and options
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		// Timer (Optional bonus)
		timer := time.NewTimer(10 * time.Second)

		// Channel to get the user input
		answerCh := make(chan int)

		// Start a goroutine to read user input
		go func() {
			var userInput string
			fmt.Print("Enter your answer (1-4 or 'exit' to quit): ")
			fmt.Scanln(&userInput)

			// Handle exit command
			if userInput == "exit" {
				answerCh <- -2 // special code for exit
				return
			}

			// Parse the answer input
			userAnswer, err := strconv.Atoi(userInput)
			if err != nil || userAnswer < 1 || userAnswer > 4 {
				// Invalid input
				fmt.Println("Invalid input, please enter a number between 1 and 4.")
				answerCh <- -1 // code for invalid input
			} else {
				answerCh <- userAnswer
			}
		}()

		// Wait for user input or timeout
		select {
		case userAnswer := <-answerCh:
			if userAnswer == -1 {
				// Retry the current question on invalid input
				i-- // Repeat the current question
				continue
			} else if userAnswer == -2 {
				// Exit the quiz early
				fmt.Println("You exited the quiz early.")
				return score, errors.New("quiz exited")
			} else if userAnswer == q.Answer {
				score++
			}
		case <-timer.C:
			fmt.Println("Time's up for this question!")
		}
	}
	return score, nil
}

// Main function
func main() {
	// Create a question bank
	questions := []Question{
		{"What is the capital of France?", [4]string{"Berlin", "Madrid", "Paris", "Rome"}, 3},
		{"Which language is used for Android development?", [4]string{"C", "Python", "Java", "Go"}, 3},
		{"Who is the founder of Microsoft?", [4]string{"Steve Jobs", "Elon Musk", "Bill Gates", "Mark Zuckerberg"}, 3},
	}

	// Call the takeQuiz function
	score, err := takeQuiz(questions)
	if err != nil && err.Error() == "quiz exited" {
		return // Early exit
	}

	// Calculate and display the score
	fmt.Printf("\nYour final score is: %d/%d\n", score, len(questions))
	if score == len(questions) {
		fmt.Println("Excellent!")
	} else if score >= len(questions)/2 {
		fmt.Println("Good job!")
	} else {
		fmt.Println("Needs Improvement.")
	}
}
