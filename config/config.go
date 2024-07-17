package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//Initialize SQLite
	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)
	return logger
}
