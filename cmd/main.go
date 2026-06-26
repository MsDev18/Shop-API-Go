package main

import (
	"fmt"
	"shop/internal/config"
	"shop/internal/migrator"
	"shop/internal/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	// load project configuration
	config := config.New()
	config.LoadFromDotEnv(".env")
	config.LoadFromYml("config.yml")
	cfg := config.GetConfig()
	// print configuration
	fmt.Println(cfg)
	// mysql connection 
	_ = mysql.New(cfg.MySQL)
	m := migrator.New(cfg.MySQL.GetDSN())
	m.Step(10)

}
