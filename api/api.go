package api

import (
	"github.com/MrBolas/InterviewQuestions_API/handlers"
	"github.com/MrBolas/InterviewQuestions_API/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Api struct {
	echo *echo.Echo
}

func NewApi(db *gorm.DB) *Api {

	// Echo instance
	e := echo.New()

	// Define repositories
	QuestionRepository := repositories.NewQuestionRepository(nil)

	// Define handlers
	questionsHandler := handlers.NewQuestionsHandler(QuestionRepository)

	// Routes
	e.GET("/question/:id", questionsHandler.GetQuestion)       // GET question by ID
	e.GET("/questions", questionsHandler.GetQuestions)         // GET question by query params
	e.POST("/question", questionsHandler.PostQuestion)         // POST new question
	e.DELETE("/question/:id", questionsHandler.DeleteQuestion) // DELETE question

	return &Api{
		echo: e,
	}
}

func (api *Api) Start() error {
	// Start server
	return api.echo.Start(":8080")
}
