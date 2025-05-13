package survey

import (
	"log"
	"net/http"
	"social_bot_backend/pkg/request"
	"social_bot_backend/pkg/response"
	"strconv"
)

type SurveyHandlerDeps struct {
	SurveyService *SurveyService
}

type SurveyHandler struct {
	SurveyService *SurveyService
}

func NewSurveyHandler(router *http.ServeMux, deps SurveyHandlerDeps) {
	handler := SurveyHandler{SurveyService: deps.SurveyService}
	router.HandleFunc("POST /survey", handler.CreateSurvey())
	router.HandleFunc("GET /survey", handler.GetAllSurvey())
	router.HandleFunc("GET /survey/{id}", handler.GetQuestionWithAnswers())
}

func (h *SurveyHandler) CreateSurvey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		surveyReq, err := request.HandleBody[SurveyCreateRequest](&w, r)
		if err != nil {
			log.Println(err)
			response.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.SurveyService.CreateSurvey(surveyReq)
		if err != nil {
			log.Println(err)
			response.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *SurveyHandler) GetAllSurvey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		surveys, err := h.SurveyService.GetAllSurvey()
		if err != nil {
			log.Println(err)
			response.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.JSON(w, surveys, http.StatusOK)
	}
}

func (h *SurveyHandler) GetQuestionWithAnswers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			response.JSON(w, "id is required", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			response.JSON(w, "invalid id", http.StatusBadRequest)
			return
		}

		survey, err := h.SurveyService.GetQuestionWithAnswers(uint(id))
		if err != nil {
			log.Println(err)
			response.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.JSON(w, survey, http.StatusOK)
		log.Println(survey)
	}
}
