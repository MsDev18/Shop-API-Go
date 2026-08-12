package main

import (
	"shop/internal/api/handler/health"
	"shop/internal/api/server"
	"shop/internal/migrator"
	"shop/internal/pkg/imageprocessor"
	"shop/internal/repository/mysql"

	// middlewares
	authmiddleware "shop/internal/api/middleware/auth"
	// repository
	authrepository "shop/internal/repository/mysql/auth"
	categoryrepository "shop/internal/repository/mysql/category"
	userrepository "shop/internal/repository/mysql/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load project configuration
	config := LoadConfig()
	// migration
	m := migrator.New(config.MySQL.GetDSN())
	if mErr := m.Up(); mErr != nil {
		panic(mErr)
	}
	// mysql connection
	mysqlRepo := mysql.New(config.MySQL)
	authRepository := authrepository.New(mysqlRepo)
	userRepository := userrepository.New(mysqlRepo)
	categoryRepository := categoryrepository.New(mysqlRepo)
	// setup middlewares
	authMiddleware := authmiddleware.New(config.AuthService.AccessTokenSecret, authRepository)
	// setup image processor
	imageProcessor := imageprocessor.New(config.Upload)
	// setup project handlers
	healthHandler := health.New()
	authHandler := SetupAuthModule(authRepository, config.AuthService)
	userHandler := SetupUserModule(userRepository, imageProcessor)
	categoryHandler := SetupCategoryModule(categoryRepository)
	// create new http server and run it
	httpServer := server.New(config.Server, healthHandler, authHandler, userHandler ,categoryHandler, authMiddleware)
	httpServer.Run()
}
