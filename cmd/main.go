package main

import (
	"log"

	"github.com/MrBolas/InterviewQuestions_API/cmd/parsing"
)

var filePaths = []string{"./questions/Interview Questions BE - BE.csv",
	"./questions/Interview Questions BE - CAP Theorem.csv",
	"./questions/Interview Questions BE - Concurrency.csv",
	"./questions/Interview Questions BE - Databases.csv",
}

func main() {

	questions := parsing.ReadQuestionsFromFilePaths(filePaths)
	log.Println("Loaded", len(questions), "questions.")

	// Post Questions to database
}
