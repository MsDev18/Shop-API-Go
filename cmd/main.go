package main

import (
	authhandler "shop/internal/api/handler/auth"
	"shop/internal/api/handler/health"
	"shop/internal/api/server"
	"shop/internal/config"
	"shop/internal/migrator"
	"shop/internal/repository/mysql"
	authrepository "shop/internal/repository/mysql/auth"
	authservice "shop/internal/service/auth"
	authvalidator "shop/internal/validator/auth"

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
	// setup project handlers
	healthHandler := health.New()
	authHandler := SetupAuthModule(mysqlRepo)
	// create new http server and run it
	httpServer := server.New(cfg.Server, healthHandler, authHandler)
	httpServer.Run()
}

func LoadConfig() config.Config {
	appConfig := config.New()
	appConfig.LoadFromDotEnv(".env")
	appConfig.LoadFromYml("config.yml")
	return appConfig.GetConfig()
}

func SetupAuthModule(mysqlRepo mysql.Connection) authhandler.Handler {
	authRepository := authrepository.New(mysqlRepo)
	authService := authservice.New(authRepository)
	authValidator := authvalidator.New(authRepository)
	authHandler := authhandler.New(authRepository, authService, authValidator)
	return authHandler
}
