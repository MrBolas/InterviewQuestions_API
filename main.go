package main

import (
	"log"

	"github.com/MrBolas/InterviewQuestions_API/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	// initialize db connection
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Fatalf("unale to connect to database: %v", err)
	}

	// Start API
	a := api.NewApi(db)
	err = a.Start()
	if err != nil {
		log.Fatalf("unable to start echo: %v", err)
	}
}
