package main

import (
	"log"

	"github.com/MrBolas/InterviewQuestions_API/cmd/client"
	"github.com/MrBolas/InterviewQuestions_API/cmd/parsing"
)

var filePaths = []string{"./questions/Interview Questions BE - BE.csv",
	"./questions/Interview Questions BE - CAP Theorem.csv",
	"./questions/Interview Questions BE - Concurrency.csv",
	"./questions/Interview Questions BE - Databases.csv",
}

func main() {

	client := client.NewQuestionClient()

	questions := parsing.ReadQuestionsFromFilePaths(filePaths)
	log.Println("Loaded", len(questions), "questions.")

	for _, q := range questions {
		err := client.CreateQuestion(q)
		if err != nil {
			log.Panic(err)
		}
	}
}
