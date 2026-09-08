package configuration

import (
	"os"
	"strings"
)

// Config centraliza as configurações da aplicação (equivalente ao application.properties do Spring Boot)
type Config struct {
	Environment string
	Port        string
	DatabaseURL string
}

func Load() *Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// Fallback para desenvolvimento local conectando no Docker
		databaseURL = "server=localhost;port=1433;user id=sa;password=YourStrong@Passw0rd!;database=storage;encrypt=disable;trustservercertificate=true"
	}

	return &Config{
		Environment: env,
		Port:        port,
		DatabaseURL: databaseURL,
	}
}
