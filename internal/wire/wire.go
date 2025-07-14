package wire

import (
	"go-33/internal/adaptor"
	"go-33/internal/data/repository"
	"go-33/internal/usecase"
	"go-33/pkg/middleware"
	"go-33/pkg/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Wiring(repo repository.Repository, mLogger middleware.LoggerMiddleware, middlwareAuth middleware.AuthMiddleware, logger *zap.Logger, config utils.Configuration) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mLogger.LoggingMiddleware)
	router.Use(middlwareAuth.Auth)
	rV1 := chi.NewRouter()
	wireUser(rV1, middlwareAuth, repo, logger, config)
	router.Mount("/api/v1", rV1)

	return router
}

func wireUser(router *chi.Mux, middlwareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseUser := usecase.NewUserService(repo, logger, config)
	adaptorUser := adaptor.NewHandlerUser(usecaseUser, logger)
	router.Post("/register", adaptorUser.Register)
	router.Get("/users", adaptorUser.ListUser)
}
