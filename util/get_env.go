package util

import (
	"os"

	"github.com/joho/godotenv"
)

func Getenv(variable string) string {
	// check the local .env file first
	var localEnv map[string]string
	localEnv, err := godotenv.Read()

	var value string

	if err == nil {
		value = localEnv[variable]
	}

	// if value is not found in .env fallback to OS environment
	if value == "" {
		value = os.Getenv(variable)
	}

	return value
}
