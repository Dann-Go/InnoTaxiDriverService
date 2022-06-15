package repository

import (
	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateDriver(driver *domain.Driver) (*domain.Driver, error)
	GetDriverByPhone(phone string) (*domain.Driver, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
