package repository

import (
	"context"
	"errors"
	"go-33/internal/data/entity"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	List(ctx context.Context) (*entity.User, error)
}

type userRepositoryImpl struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	query := `
		INSERT INTO users (name, email, password, photo, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`
	result := r.DB.Raw(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Photo,
	).Scan(&user)
	if result.Error != nil {
		return errors.New("error insert data")
	}

	return nil
}

func (r *userRepositoryImpl) List(ctx context.Context) (*entity.User, error) {
	// query := `SELECT pg_sleep(5)` // PostgreSQL delay 10 detik

	// userID := ctx.Value("userID")
	// fmt.Println("user id on repository", userID)

	// _, err := r.DB.ExecContext(ctx, query)
	// if err != nil {
	// 	fmt.Println("Query canceled:", err)
	// 	return nil, err
	// }

	// user := entity.User{
	// 	Name:  "lumoshive",
	// 	Email: "lumoshive@email.com",
	// 	Photo: "https://lumoshive.com/image.jpg",
	// }
	// return &user, nil
	return nil, nil

}
