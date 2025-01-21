package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

var Variables *Config

type Config struct {
	BaseURL   string
	TargetURL string
}

func Init() error {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}

	Variables = &Config{
		BaseURL:   os.Getenv("BASE_URL"),
		TargetURL: os.Getenv("TARGET_URL"),
	}

	return nil
}
