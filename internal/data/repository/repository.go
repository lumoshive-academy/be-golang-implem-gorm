package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo   UserRepository
	WalletRepo WalletRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		UserRepo: NewUserRepository(db, log),
		// WalletRepo: NewWalletRepository(db, log),
	}
}
