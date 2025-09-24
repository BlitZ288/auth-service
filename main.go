package main

import (
	"fmt"
	"log"

	"github.com/BlitZ288/auth-service/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	_ = godotenv.Load()
	fmt.Println("Start")
	var dbConfiguration config.DatabaseConfiguration
	var authnConfiguration config.AuthorizationConfiguration
	if err := envconfig.Process("", &dbConfiguration); err != nil {
		log.Fatal("Не удалось загрузить конфигурацию бд error: ", err.Error())
	}

	if err := envconfig.Process("", &authnConfiguration); err != nil {
		log.Fatal("Не удалось загрузить конфигурацию auth error: ", err.Error())
	}
}
