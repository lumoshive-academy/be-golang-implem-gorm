package main

import (
	"go-33/cmd"
	"go-33/internal/data"
	"go-33/internal/data/repository"
	"go-33/internal/wire"
	"go-33/pkg/database"
	"go-33/pkg/middleware"
	"go-33/pkg/utils"
	"log"

	"go.uber.org/zap"
)

func main() {

	// read config
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	logger, err := utils.InitLogger(config.PathLogger, config)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}

	//Init db
	db, err := database.InitDB(config)
	if err != nil {
		logger.Fatal("can't connect to database ", zap.Error(err))
	}

	// migration
	if err := data.AutoMigrate(db); err != nil {
		logger.Fatal("failed to run migrations", zap.Error(err))
	}

	// seeder
	if err := data.SeedAll(db); err != nil {
		logger.Fatal("failed to seed initial data", zap.Error(err))
	}

	repo := repository.NewRepository(db, logger)
	mLogger := middleware.NewLoggerMiddleware(logger)
	mAuth := middleware.NewAuthMiddleware(logger)
	router := wire.Wiring(repo, mLogger, mAuth, logger, config)

	cmd.ApiServer(config, logger, router)
}
