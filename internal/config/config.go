package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB  string
	Port     string
}

func Load() (*Config, error) {

	// ? Load the env by godotenv into process :
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("Failed to load env.")
	}

	MongoUri, err := extractEnv("MONGO_URI")

	if err != nil {
		return nil, err
	}

	MongoDB, err := extractEnv("MONGO_DB")

	if err != nil {
		return nil, err
	}

	Port, err := extractEnv("PORT")

	if err != nil {
		return nil, err
	}

	return &Config{
		MongoURI: MongoUri,
		MongoDB:  MongoDB,
		Port:     Port,
	}, nil

}

func extractEnv(key string) (string, error) {

	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("%s is required.", key)
	}

	return val, nil

}
