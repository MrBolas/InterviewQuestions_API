package models

import "github.com/gofrs/uuid"

type QuestionRequest struct {
	Question string
	Answer   string
	Category string
	Level    string
	Source   string
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
	ID       uuid.UUID
	Question string
	Answer   string
	Category string
	Level    string
	Source   string
}
