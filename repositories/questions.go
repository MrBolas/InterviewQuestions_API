package repositories

import (
	"github.com/MrBolas/InterviewQuestions_API/cmd/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetQuestionById(id uuid.UUID) (models.Question, error)
	ListQuestions(filters map[string]interface{}) ([]models.Question, error)
	CreateQuestion(q models.Question) (models.Question, error)
	DeleteQuestionById(id uuid.UUID) error
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r QuestionRepository) GetQuestionById(id uuid.UUID) (models.Question, error) {
	return models.Question{}, nil
}

func (r QuestionRepository) ListQuestions(filters map[string]interface{}) ([]models.Question, error) {
	return []models.Question{}, nil
}

func (r QuestionRepository) CreateQuestion(q models.Question) (models.Question, error) {
	return models.Question{}, nil
}

func (r QuestionRepository) DeleteQuestionById(id uuid.UUID) error {
	return nil
}
