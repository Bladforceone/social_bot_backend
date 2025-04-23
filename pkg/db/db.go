package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"social_bot_backend/configs"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(conf.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{db}
}
