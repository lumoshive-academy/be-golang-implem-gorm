package usecase

import (
	"context"
	"fmt"
	"go-33/internal/data/entity"
	"go-33/internal/data/repository"
	"go-33/pkg/codes"
	"go-33/pkg/utils"
	"mime/multipart"

	"go.uber.org/zap"
)

type UserService interface {
	Create(user *entity.User, file multipart.File) error
	List(ctx context.Context) (*entity.User, error)
}

type userService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewUserService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) UserService {
	return &userService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *userService) Create(user *entity.User, file multipart.File) error {
	err := codes.UploadFile(file, user.Photo, s.Logger, s.Config)
	if err != nil {
		s.Logger.Error("failed to upload file:", zap.Error(err))
		return err
	}

	// Create user to DB
	err = s.Repo.UserRepo.Create(user)
	if err != nil {
		s.Logger.Error("failed to create user:", zap.Error(err))
		return err
	}

	// Create wallet with default balance 100
	wallet := &entity.Wallet{
		UserID:  user.ID,
		Balance: 100,
	}

	err = s.Repo.WalletRepo.Create(wallet)
	if err != nil {
		s.Logger.Error("failed to create wallet:", zap.Error(err))
		return err
	}

	return nil
}

func (s *userService) List(ctx context.Context) (*entity.User, error) {
	// get user id
	userID := ctx.Value("userID")
	fmt.Println("user id ", userID)
	return s.Repo.UserRepo.List(ctx)
}
