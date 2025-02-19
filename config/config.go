package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type DB struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	Database string `env:"DB_NAME"`
}

type Node struct {
	ID      uint32 `env:"NODE_ID"`
	Address string `env:"NODE_ADDRESS"`
}

type Config struct {
	DB         DB
	Node       Node
	NodeNumber int    `env:"NODE_NUMBER"`
	WebhookURL string `env:"WEBHOOK_URL"`
	EncryptKey string `env:"ENCRYPT_KEY"`
	RedisAddr  string `env:"REDIS_ADDR"`
}

func Load() (*Config, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
