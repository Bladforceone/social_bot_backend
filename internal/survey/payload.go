package survey

import (
	"gorm.io/datatypes"
)

type SurveyCreateRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	IsPublic    bool             `json:"is_public"`
	Questions   []datatypes.JSON `json:"questions"`
}

type AllSurveyResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type SurveyResponse struct {
	Question string   `json:"question"`
	Type     string   `json:"type"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Answer string `json:"answer"`
	Count  int    `json:"count"`
}
