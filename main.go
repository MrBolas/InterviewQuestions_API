package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MrBolas/InterviewQuestions_API/api"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadEnvVariables() map[string]string {

	envVariables := make(map[string]string)

	// load en variables from ".env" file
	godotenv.Load()

	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		panic("DB_HOST undefined")
	}

	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		panic("DB_USER undefined")
	}

	dbPw, exists := os.LookupEnv("DB_PW")
	if !exists {
		panic("DB_PW undefined")
	}

	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		panic("DB_NAME undefined")
	}

	dbPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		panic("DB_PORT undefined")
	}

	svPort, exists := os.LookupEnv("PORT")
	if !exists {
		panic("PORT undefined")
	}

	envVariables["DB_HOST"] = dbHost
	envVariables["DB_USER"] = dbUser
	envVariables["DB_PW"] = dbPw
	envVariables["DB_NAME"] = dbName
	envVariables["DB_PORT"] = dbPort
	envVariables["PORT"] = svPort

	return envVariables
}

func main() {

	envVar := loadEnvVariables()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", envVar["DB_HOST"], envVar["DB_USER"], envVar["DB_PW"], envVar["DB_NAME"], envVar["DB_PORT"])

	// initialize db connection
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Fatalf("unale to connect to database: %v", err)
	}

	// Start API
	a := api.NewApi(db)
	err = a.Start(envVar["PORT"])
	if err != nil {
		log.Fatalf("unable to start echo: %v", err)
	}
}
