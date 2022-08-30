package config

import (
	"os"
	"strconv"
)

type Config struct {
	DbHost     string
	DbUser     string
	DbPort     int
	DbPassword string
	DbName     string
	ApiPort    int
}

func GetConfig() *Config {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))
	return &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPort:     port,
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		ApiPort:    apiPort,
	}
}
