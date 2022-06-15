package service

import (
	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/repository"
)

type Authorization interface {
	CreateDriver(driver *domain.Driver) (*domain.Driver, error)
	GenerateToken(phone, password string) (string, error)
	GetDriverByPhone(phone string) (*domain.Driver, error)
	ParseToken(token string) (int, error)
}

type Driver interface {
	RateClient()
	GetRating()
	GetAllDrives()
}

type Service struct {
	Authorization
	Driver
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
	}
}
