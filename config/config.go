package config

import (
	//"errors"

	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initilizing sqlite: %v", err)
	}
	//return errors.New("faker error")
	return nil
}



func GetSQLite() *gorm.DB{
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger

	logger = NewLogger(p)
	return logger
}
