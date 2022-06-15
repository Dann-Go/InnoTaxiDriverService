package config

import (
	"github.com/pkg/errors"
	"os"
	"strings"
)

type DbPostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDbConfig() *DbPostgresConfig {
	cfg := &DbPostgresConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DBPORT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	}

	return cfg
}

func EnvsCheck() error {
	requiredEnvs := []string{"HOST", "DBPORT", "USERNAME", "PASSWORD",
		"DBNAME", "SSLMODE", "SERVPORT"}
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
