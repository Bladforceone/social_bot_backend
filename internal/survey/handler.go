package survey

import (
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
}

func (h *SurveyHandler) CreateSurvey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		surveyReq, err := request.HandleBody[SurveyCreateRequest](&w, r)
		if err != nil {
			response.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.SurveyService.CreateSurvey(surveyReq)
		if err != nil {
			response.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
