package main

import (
	authhandler "shop/internal/api/handler/auth"
	"shop/internal/api/handler/health"
	userhandler "shop/internal/api/handler/user"
	authmiddleware "shop/internal/api/middleware/auth"
	"shop/internal/api/server"
	"shop/internal/config"
	"shop/internal/migrator"
	"shop/internal/pkg/imageprocessor"
	"shop/internal/repository/mysql"
	authrepository "shop/internal/repository/mysql/auth"
	userrepository "shop/internal/repository/mysql/user"
	authservice "shop/internal/service/auth"
	userservice "shop/internal/service/user"
	authvalidator "shop/internal/validator/auth"
	uservalidator "shop/internal/validator/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load project configuration
	cfg := LoadConfig()
	// migration
	m := migrator.New(cfg.MySQL.GetDSN())
	if mErr := m.Up(); mErr != nil {
		panic(mErr)
	}
	// mysql connection
	mysqlRepo := mysql.New(cfg.MySQL)
	authRepository := authrepository.New(mysqlRepo)
	userRepository := userrepository.New(mysqlRepo)
	// setup middlewares
	authMiddleware := authmiddleware.New(cfg.AuthService.AccessTokenSecret, authRepository)
	// setup image processor
	imageProcessor := imageprocessor.New(cfg.Upload)
	// setup project handlers
	healthHandler := health.New()
	authHandler := SetupAuthModule(authRepository, cfg.AuthService)
	userHandler := SetupUserModule(userRepository, imageProcessor)
	// create new http server and run it
	httpServer := server.New(cfg.Server, healthHandler, authHandler, userHandler, authMiddleware)
	httpServer.Run()
}

func LoadConfig() config.Config {
	appConfig := config.New()
	appConfig.LoadFromDotEnv(".env")
	appConfig.LoadFromYml("config.yml")
	return appConfig.GetConfig()
}

func SetupAuthModule(authRepository authrepository.Repository, cfg authservice.Config) authhandler.Handler {
	authService := authservice.New(authRepository, cfg)
	authValidator := authvalidator.New()
	authHandler := authhandler.New(authRepository, authService, authValidator)
	return authHandler
}

func SetupUserModule(userRepository userrepository.Repository, imageProcessor imageprocessor.Processor) userhandler.Handler {
	userService := userservice.New(userRepository, imageProcessor)
	userValidator := uservalidator.New()
	userHandler := userhandler.New(userRepository, userService, userValidator)
	return userHandler
}
