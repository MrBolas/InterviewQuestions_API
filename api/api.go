package api

import (
	"github.com/MrBolas/InterviewQuestions_API/handlers"
	"github.com/MrBolas/InterviewQuestions_API/models"
	"github.com/MrBolas/InterviewQuestions_API/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Api struct {
	echo *echo.Echo
}

func NewApi(db *gorm.DB) *Api {

	err := db.AutoMigrate(models.Question{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	_, err = sqlDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Define repositories
	QuestionRepository := repositories.NewQuestionRepository(db)

	// Define handlers
	questionsHandler := handlers.NewQuestionsHandler(QuestionRepository)

	// Routes
	e.GET("/question/:id", questionsHandler.GetQuestion)       // GET question by ID
	e.GET("/questions", questionsHandler.ListQuestions)        // GET question by query params
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
