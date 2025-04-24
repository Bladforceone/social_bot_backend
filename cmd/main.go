package main

import (
	"fmt"
	"net/http"
	"social_bot_backend/configs"
	"social_bot_backend/internal/survey"
	"social_bot_backend/pkg/db"
	"social_bot_backend/pkg/midleware"
)

func main() {
	config := configs.LoadConfig()
	database := db.NewDB(config)
	router := http.NewServeMux()

	// Repositories
	surveyRepository := survey.NewSurveyRepository(database)

	// Services
	surveyService := survey.NewSurveyService(surveyRepository)

	// Handlers
	survey.NewSurveyHandler(router, survey.SurveyHandlerDeps{SurveyService: surveyService})

	server := http.Server{
		Addr:    ":8080",
		Handler: midleware.Logging(router),
	}

	fmt.Println("Listening on port 8080")
	server.ListenAndServe()
}
