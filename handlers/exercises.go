package handlers

import (
	"net/http"

	"github.com/MrBolas/InterviewQuestions_API/network"
	"github.com/labstack/echo/v4"
)

type ExerciseHandler struct {
	OllamaClient *network.OllamaClient
}

type ExerciseResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func NewExercisesHandler(OllamaClient *network.OllamaClient) *ExerciseHandler {
	return &ExerciseHandler{
		OllamaClient: OllamaClient,
	}
}

func (h *ExerciseHandler) GenerateExercise(c echo.Context) error {

	question, err := h.OllamaClient.GenerateQuestion("golang", "backend", "easy")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	answer, err := h.OllamaClient.GenerateAnswer(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	exercise := ExerciseResponse{
		Question: question,
		Answer:   answer,
	}
	return c.JSON(http.StatusOK, exercise)
}
