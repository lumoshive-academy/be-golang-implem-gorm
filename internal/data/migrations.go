package data

import (
	"go-33/internal/data/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
	)
}
