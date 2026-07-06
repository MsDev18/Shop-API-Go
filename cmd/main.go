package main

import (
	"shop/internal/api/server"
	"shop/internal/config"
	"shop/internal/migrator"
	"shop/internal/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load project configuration
	appConfig := config.New()
	appConfig.LoadFromDotEnv(".env")
	appConfig.LoadFromYml("config.yml")
	cfg := appConfig.GetConfig()
	// migration 
	m := migrator.New(cfg.MySQL.GetDSN())
	if mErr := m.Up(); mErr != nil {
		panic(mErr)
	}
	// mysql connection
	_ = mysql.New(cfg.MySQL)
	
	// create new http server and run it
	httpServer := server.New(cfg.Server)
	httpServer.Run()
}
