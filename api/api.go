package api

import (
	"net/http"

	"github.com/MrBolas/InterviewQuestions_API/handlers"
	"github.com/MrBolas/InterviewQuestions_API/models"
	"github.com/MrBolas/InterviewQuestions_API/repositories"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	//middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	e.GET("/question/:id", questionsHandler.GetQuestion)       // GET question by ID
	e.GET("/questions", questionsHandler.ListQuestions)        // GET question by query params
	e.POST("/question", questionsHandler.PostQuestion)         // POST new question
	e.DELETE("/question/:id", questionsHandler.DeleteQuestion) // DELETE question

	e.File("/", "./frontend/interview_questions/out/index.html")
	e.Static("/_next/static/", "./frontend/interview_questions/out/_next/static/")

	return &Api{
		echo: e,
	}
}

func (api *Api) Start() error {
	// Start server
	return api.echo.Start(":8080")
}
