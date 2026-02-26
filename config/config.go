package config

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type (
	ApiConfig struct {
		ApiPort string
	}

	DbConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Drive    string
	}

	Config struct {
		ApiConfig
		DbConfig
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
	if err := c.setENV(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Drive:    os.Getenv("DB_DRIVE"),
	}

	if c.ApiPort == "" {
		return errors.New("enviroment required")
	}
	return nil
}

func (c *Config) setENV() error {
	envFile, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		envVar := strings.SplitN(line, "=", 2)

		if len(envVar) != 2 {
			log.Printf("Invalid line in .env file: %s", line)
			continue
		}

		key := strings.TrimSpace(envVar[0])
		value := strings.TrimSpace(envVar[1])

		if key == "" {
			log.Printf("Empty key in .env file: %s", line)
			continue
		}

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
