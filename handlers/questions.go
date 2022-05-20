package handlers

import (
	"net/http"

	"github.com/MrBolas/InterviewQuestions_API/models"
	"github.com/MrBolas/InterviewQuestions_API/repositories"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type QuestionsHandler struct {
	repo repositories.Repository
}

func NewQuestionsHandler(repo repositories.Repository) *QuestionsHandler {
	return &QuestionsHandler{
		repo: repo,
	}
}

func (h *QuestionsHandler) GetQuestion(c echo.Context) error {

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Question not found")
	}

	question, err := h.repo.GetQuestionById(id)
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, "Question "+id.String()+" not found")
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, question.ToResponse())
}

func (h *QuestionsHandler) ListQuestions(c echo.Context) error {

	// pagination

	// sanitized filters for questions
	filters, err := repositories.NewFilters(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Query Parameters")
	}

	// Sort

	questions, err := h.repo.ListQuestions(filters)
	if err != nil {
		return err
	}

	questionCount, err := h.repo.CountListQuestions(filters)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if questionCount <= 0 {
		return c.JSON(http.StatusNotFound, "No Questions Found")
	}

	return c.JSON(http.StatusOK, models.ToListResponse(questions))
}

func (h *QuestionsHandler) PostQuestion(c echo.Context) error {

	req := new(models.QuestionRequest)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Mal-formed JSON")
	}

	question, err := h.repo.CreateQuestion(req.ToQuestion())
	if err == gorm.ErrRegistered {
		return c.JSON(http.StatusConflict, "Conflict with new question.")
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, question.ToResponse())
}

func (h *QuestionsHandler) DeleteQuestion(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}
