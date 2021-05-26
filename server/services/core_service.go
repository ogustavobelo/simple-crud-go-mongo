package services

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvCheck() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env - ", err)
	}

	envItems := []string{"SERVER_PORT", "DATABASE_PORT", "DATABASE_NAME", "JWT_SECRET_KEY", "JWT_ISSURE"}

	for _, item := range envItems {
		env, envExist := os.LookupEnv(item)
		if !envExist || env == "" {
			log.Fatalf("%v must be set in .env and not be empty", item)
		}
	}
}
