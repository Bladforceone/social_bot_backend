package survey

import (
	"log"
	"net/http"
	"social_bot_backend/pkg/request"
	"social_bot_backend/pkg/response"
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
	router.HandleFunc("GET /survey/{id}", handler.GetSurveyById())
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

// GetSurveyById TODO: Implement
func (h *SurveyHandler) GetSurveyById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
