package handlers

import (
	"net/http"

	"github.com/MrBolas/InterviewQuestions_API/repositories"
	"github.com/labstack/echo/v4"
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

	return c.NoContent(http.StatusOK)
}

func (h *QuestionsHandler) GetQuestions(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}

func (h *QuestionsHandler) PostQuestion(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}

func (h *QuestionsHandler) DeleteQuestion(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}
