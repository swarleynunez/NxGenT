package utils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {

	// Load .env keys for this process
	err := godotenv.Load(".env")
	CheckError(err, ErrorMode)
}

func GetEnv(key string) (value string) {

	value, found := os.LookupEnv(key)
	if !found || value == "" {
		CheckError(errors.New(fmt.Sprintf("\"%s\" environment variable not set", key)), ErrorMode)
	}

	return
}

func SetEnv(key, value string) {

	// Read .env keys into a map
	env, err := godotenv.Read(".env")
	CheckError(err, ErrorMode)

	// Add or modify a key-value
	env[key] = value

	// Write map into .env file
	err = godotenv.Write(env, ".env")
	CheckError(err, ErrorMode)

	// Reload .env configuration
	overloadEnv()
}

func overloadEnv() {

	// Reload .env keys for this process
	err := godotenv.Overload(".env")
	CheckError(err, ErrorMode)
}
