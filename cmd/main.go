package main

import (
	"fmt"
	"shop/internal/config"
)

func main() {
	fmt.Println("Hi ...")
	// load project configuration
	config := config.New()
	config.LoadFromDotEnv(".env")
	config.LoadFromYml("config.yml")
	cfg := config.GetConfig()
	
	fmt.Println(cfg)
}
