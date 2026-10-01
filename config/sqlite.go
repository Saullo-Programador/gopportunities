package config

import (
	"os"

	"github.com/Saullo-Programador/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("SQLite database file does not exist. Creating a new one at %s", dbPath)

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("failed to create directory for SQLite database: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("failed to create SQLite database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("failed to migrate SQLite database: %v", err)
		return nil, err
	}

	return db, nil
}
