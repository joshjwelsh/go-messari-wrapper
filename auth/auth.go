package auth

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	secret string
}

// Load dotenv file
func NewConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("New config failed due to godotenv.Load() error: %v", err)
	}
	return Config{
		secret: os.Getenv("SECRET"),
	}, nil
}

func (c *Config) Get() string {
	return c.secret
}
