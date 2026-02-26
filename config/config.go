package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type (
	APIConfig struct {
		APIPort string
	}

	DBConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Driver   string
	}
	TokenConfig struct {
		IssuerName      string
		JwtSignatureKey []byte
		JwtLifeTime     time.Duration
	}
	Config struct {
		APIConfig
		DBConfig
		TokenConfig
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.APIConfig = APIConfig{
		APIPort: os.Getenv("API_PORT"),
	}
	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
	lifetimeStr := os.Getenv("TOKEN_LIFE_TIME")

	lifetimeInt, err := strconv.Atoi(lifetimeStr)
	if err != nil {
		return err
	}

	jwtDuration := time.Duration(lifetimeInt) * time.Hour
	c.TokenConfig = TokenConfig{
		IssuerName:      os.Getenv("TOKEN_ISSUE_NAME"),
		JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
		JwtLifeTime:     jwtDuration,
	}
	if c.APIPort == "" || c.Host == "" || c.Port == "" || c.Name == "" || c.User == "" || c.IssuerName == "" ||
		len(c.JwtSignatureKey) == 0 || c.JwtLifeTime == 0 {
		return errors.New("environment required")
	}
	return nil
}
