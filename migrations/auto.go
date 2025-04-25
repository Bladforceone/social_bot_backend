package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"social_bot_backend/internal/survey"
	"social_bot_backend/internal/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN is not set in .env file")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	// Сносим всё к чертям
	/*db.Migrator().DropTable(
		&survey.Survey{},
		&survey.Question{},
		&user.User{},
		&user.CurrentUserQuestion{},
		&user.UserAnswer{},
	)*/
	db.AutoMigrate(
		&survey.Survey{},
		&survey.Question{},
		&user.User{},
		&user.CurrentUserQuestion{},
		&user.UserAnswer{},
	)

	fmt.Println("Migrations completed")
}
