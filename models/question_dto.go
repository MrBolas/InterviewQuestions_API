package models

import "github.com/gofrs/uuid"

type QuestionRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Category string `json:"category"`
	Level    string `json:"level"`
	Source   string `json:"source"`
}

func (req QuestionRequest) ToQuestion() Question {
	return Question{
		Question: req.Question,
		Answer:   req.Answer,
		Category: req.Category,
		Level:    req.Level,
		Source:   req.Source,
	}
}

func ToListResponse(questions []Question) []QuestionResponse {

	var questionResponseList []QuestionResponse

	for _, q := range questions {
		questionResponseList = append(questionResponseList, q.ToResponse())
	}

	return questionResponseList
}

type QuestionResponse struct {
	ID       uuid.UUID `json:"id"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
	Category string    `json:"category"`
	Level    string    `json:"level"`
	Source   string    `json:"source"`
}
