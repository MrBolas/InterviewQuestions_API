package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Question struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	Question  string
	Answer    string
	Category  string
	Level     string
	Source    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Question) ToResponse() QuestionResponse {
	return QuestionResponse{
		ID:       q.ID,
		Question: q.Question,
		Answer:   q.Answer,
		Category: q.Category,
		Level:    q.Level,
		Source:   q.Source,
	}
}
