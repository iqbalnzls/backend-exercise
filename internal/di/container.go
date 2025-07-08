package di

import (
	"go.uber.org/zap"

	listingsService "github.com/iqbalnzls/backend-exercise/internal/app/usecase/listings"
	userService "github.com/iqbalnzls/backend-exercise/internal/app/usecase/users"
	listingsRepository "github.com/iqbalnzls/backend-exercise/internal/infrastructure/repository/listings"
	userRepository "github.com/iqbalnzls/backend-exercise/internal/infrastructure/repository/users"
	"github.com/iqbalnzls/backend-exercise/pkg/config"
	"github.com/iqbalnzls/backend-exercise/pkg/database"
	"github.com/iqbalnzls/backend-exercise/pkg/validator"
)

type Container struct {
	Config         *config.Config
	ListingService listingsService.Service
	UserService    userService.Service
	Validator      *validator.Validator
	Logger         *zap.Logger
}

func SetupContainer() *Container {
	cfg := config.NewConfig("resources/config.json")

	//initialize the database connection
	db := database.NewDatabase(cfg.Database)

	//initialize the repository
	listingRepo := listingsRepository.NewRepository(db)
	userRepo := userRepository.NewRepository(db)

	//initialize the service
	listingSvc := listingsService.NewService(listingRepo)
	userSvc := userService.NewService(userRepo)

	//initialize the validator
	v := validator.NewValidator()

	//initialize the logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &Container{
		Config:         cfg,
		ListingService: listingSvc,
		UserService:    userSvc,
		Validator:      v,
		Logger:         logger,
	}
}
