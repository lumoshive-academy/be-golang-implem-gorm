package repository

import (
	"database/sql"
	"go-33/internal/data/entity"

	"go.uber.org/zap"
)

type WalletRepository interface {
	Create(wallet *entity.Wallet) error
}

type walletRepository struct {
	DB  *sql.DB
	Log *zap.Logger
}

func NewWalletRepository(db *sql.DB, logger *zap.Logger) WalletRepository {
	return &walletRepository{db, logger}
}

func (r *walletRepository) Create(wallet *entity.Wallet) error {
	query := `
		INSERT INTO wallets (user_id, balance, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(
		query,
		wallet.UserID,
		wallet.Balance,
	).Scan(&wallet.ID, &wallet.CreatedAt, &wallet.UpdatedAt)
}
