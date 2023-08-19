package sqlite

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormOpenDB() (*gorm.DB, error) {
	var store = "/.local/share/storydb/"
	homedir, err := os.UserHomeDir()
	os.Mkdir(homedir+store, 0755)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent, // Silent , Error, Warning, Info
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open(homedir+store+"sqlite.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return db, fmt.Errorf("Failed to open database: %w", err)
	}

	return db, err
}
