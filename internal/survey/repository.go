package survey

import (
	"social_bot_backend/pkg/db"
)

type SurveyRepository struct {
	DB *db.DB
}

func NewSurveyRepository(db *db.DB) *SurveyRepository {
	return &SurveyRepository{DB: db}
}

func (repo *SurveyRepository) CreateSurvey(survey *Survey) error {
	return repo.DB.Create(survey).Error
}

func (repo *SurveyRepository) GetAllSurvey() *[]Survey {
	var survey []Survey
	repo.DB.Find(&survey)
	return &survey
}
