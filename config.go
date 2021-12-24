package main

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBDriver   string
	DBName     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBPort     string
}

func LoadConfig() *Config {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("[config]: error loading env file %s", err)
	}

	config := &Config{
		Port:       env["PORT"],
		DBDriver:   env["DB_DRIVER"],
		DBName:     env["DB_NAME"],
		DBHost:     env["DB_HOST"],
		DBUser:     env["DB_USER"],
		DBPassword: env["DB_PASSWORD"],
		DBPort:     env["DB_PORT"],
	}

	return config
}
