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

func (repo *SurveyRepository) GetQuestionWithAnswers(id uint) (*Survey, error) {
	var survey Survey
	err := repo.DB.Preload("Questions.UserAnswer").First(&survey, id).Error
	if err != nil {
		return nil, err
	}

	return &survey, nil
}
