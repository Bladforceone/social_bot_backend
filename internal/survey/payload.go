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
