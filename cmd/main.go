package main

import (
	"context"
	"fmt"
	"shop/internal/config"
	authdto "shop/internal/dto/auth"
	authrepository "shop/internal/repository/mysql/auth"
	authservice "shop/internal/service/auth"
	authvalidator "shop/internal/validator/auth"
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
	// migration 
	m := migrator.New(cfg.MySQL.GetDSN())
	if mErr := m.Up(); mErr != nil {
		panic(mErr)
	}
	// mysql connection
	mySqlRepo := mysql.New(cfg.MySQL)
	// auth dependencies
	authRepo := authrepository.New(mySqlRepo)
	authValidator := authvalidator.New(authRepo)
	authService := authservice.New(authRepo)
	// validation body of request 
	isValid , err := authValidator.SendOtp(context.Background(),authdto.SendOtpRequest{PhoneNumber: "09351721415"})
	if err != nil {
		fmt.Println(err)
		return
	}
	if !isValid {
		fmt.Println("not valid")
		return 
	}
	// call the auth service 
	res, err := authService.SendOtp(authdto.SendOtpRequest{
		PhoneNumber: "09126359202",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
