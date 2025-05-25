package config

import (
	"os"

	"github.com/victorramos887/golang_oportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./db/main.db"

	// Garante que o diretório existe
	dir := "./db"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Garante que o arquivo existe
	_, statErr := os.Stat(dbPath)
	if os.IsNotExist(statErr) {
		logger.Info("database file not found, creating ...")
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
