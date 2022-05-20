package repositories

import (
	"github.com/MrBolas/InterviewQuestions_API/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetQuestionById(id uuid.UUID) (models.Question, error)
	ListQuestions(filters QuestionFilters) ([]models.Question, error)
	CreateQuestion(q models.Question) (models.Question, error)
	CountListQuestions(filters QuestionFilters) (int64, error)
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

	var question models.Question

	if err := r.db.First(&question, id).Error; err != nil {
		return models.Question{}, err
	}

	return question, nil
}

func (r QuestionRepository) ListQuestions(filters QuestionFilters) ([]models.Question, error) {

	var questions []models.Question

	q := r.db.Where(filters.Filters)

	if err := q.Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r QuestionRepository) CreateQuestion(q models.Question) (models.Question, error) {

	if err := r.db.Create(&q).Error; err != nil {
		return models.Question{}, err
	}

	return q, nil
}

func (r QuestionRepository) CountListQuestions(filters QuestionFilters) (int64, error) {

	q := r.db.Model(&models.Question{}).Where(filters.Filters)

	var count int64

	if err := q.Count(&count).Error; err != nil {
		return -1, err
	}

	return count, nil
}

func (r QuestionRepository) DeleteQuestionById(id uuid.UUID) error {
	return nil
}
