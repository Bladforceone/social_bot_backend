package user

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type User struct {
	TelegramID           int64                 `gorm:"column:tg_id;type:bigint;not null;unique;primarykey"`
	Age                  int                   `gorm:"type:integer;not null"`
	Gender               string                `gorm:"type:varchar(10);not null"`
	CurrentUserQuestions []CurrentUserQuestion `gorm:"constraint: OnUpdate:CASCADE, OnDelete: CASCADE;"`
	UserAnswer           []UserAnswer          `gorm:"constraint: OnUpdate:CASCADE, OnDelete: CASCADE;"`
	DeletedAt            gorm.DeletedAt        `gorm:"index"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CurrentUserQuestion struct {
	gorm.Model
	IsCompletedSurvey bool `gorm:"default:false"`
	UserID            uint `gorm:"not null;index"`
	QuestID           uint `gorm:"not null;index"`
	SurveyID          uint `gorm:"not null;index"`
}

func (CurrentUserQuestion) TableName() string {
	return "current_users_questions"
}

type UserAnswer struct {
	gorm.Model
	UserID   uint           `gorm:"not null;index"`
	QuestID  uint           `gorm:"not null;index"`
	SurveyID uint           `gorm:"not null;index"`
	Answer   datatypes.JSON `gorm:"type:jsonb;not null"`
}

func (UserAnswer) TableName() string {
	return "users_answers"
}
