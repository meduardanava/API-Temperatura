package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetEnv() Env {
	godotenv.Load()

	err := setEnv(".env")
	if err != nil {
		panic(err)
	}
	return Env{
		Host:     os.Getenv("api-host"),
		Port:     os.Getenv("api-port"),
		User:     os.Getenv("api-user"),
		Password: os.Getenv("api-password"),
		DBName:   os.Getenv("api-dbname"),
	}
}
