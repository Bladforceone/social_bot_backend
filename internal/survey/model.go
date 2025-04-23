package survey

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"social_bot_backend/internal/user"
)

type Survey struct {
	gorm.Model
	Name                 string                     `gorm:"type:varchar(255);not null"`
	Description          string                     `gorm:"type:text"`
	CountQuestions       int                        `gorm:"type:integer;not null"`
	IsPublic             bool                       `gorm:"not null"`
	Questions            []Question                 `gorm:"constraint: OnUpdate:CASCADE, OnDelete: CASCADE;"`
	CurrentUserQuestions []user.CurrentUserQuestion `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserAnswer           []user.UserAnswer          `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Question struct {
	gorm.Model
	SurveyID             uint                       `gorm:"not null;index"`
	Question             datatypes.JSON             `gorm:"type:jsonb;not null"`
	CurrentUserQuestions []user.CurrentUserQuestion `gorm:"foreignKey:QuestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserAnswer           []user.UserAnswer          `gorm:"foreignKey:QuestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
