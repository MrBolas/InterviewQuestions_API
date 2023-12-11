package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/MrBolas/InterviewQuestions_API/cmd/models"
)

type QuestionClient struct {
	Client http.Client
	Url    url.URL
}

func NewQuestionClient() QuestionClient {
	return QuestionClient{
		Client: *http.DefaultClient,
		Url: url.URL{
			Scheme: "http",
			Host:   "localhost:8080",
		},
	}
}

func (qc *QuestionClient) CreateQuestion(question models.Question) error {

	json_data, err := json.Marshal(question)
	if err != nil {
		return err
	}

	resp, err := qc.Client.Post(qc.Url.String()+"/question", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		return errors.New("Post failed with code: " + resp.Status)
	}

	log.Println("Question from " + question.Category + " in level " + question.Level + " added.")
	return nil
}
