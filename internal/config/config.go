package config

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

type DbPostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port string
}

type AuthConfig struct {
	TokenTTL   string
	SigningKey string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: os.Getenv("SERVPORT"),
	}
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		SigningKey: os.Getenv("SIGNINGKEY"),
		TokenTTL:   os.Getenv("TOKENTTL")}
}

func NewDbConfig() *DbPostgresConfig {
	return &DbPostgresConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DBPORT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	}

}

func EnvsCheck() error {
	requiredEnvs := []string{"HOST", "DBPORT", "USERNAME", "PASSWORD",
		"DBNAME", "SSLMODE", "SERVPORT", "TOKENTTL", "SIGNINGKEY"}
	var msg []string
	for _, el := range requiredEnvs {
		val, exists := os.LookupEnv(el)
		if !exists || len(val) == 0 {
			msg = append(msg, el)
		}
	}
	if len(msg) > 0 {
		err := errors.New(strings.Join(msg, ", ") + " env(s) not set")
		return err
	}
	return nil
}
