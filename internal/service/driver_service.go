package service

import (
	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain/apperrors"

	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/repository"
)

type Driver interface {
	CreateDriver(user *domain.Driver) (*domain.DriverResponse, error)
	GetDriverByPhone(phone string) (*domain.Driver, error)
	GetDriverByEmail(email string) (*domain.Driver, error)
}

type DriverService struct {
	repo repository.Driver
}

func (s *DriverService) CreateDriver(driver *domain.Driver) (*domain.DriverResponse, error) {
	driver.PasswordHash = HashPassword(driver.PasswordHash)

	if driverCheck, err := s.repo.GetDriverByPhone(driver.Phone); driverCheck.Phone != "" {
		return nil, apperrors.Wrapper(apperrors.ErrPhoneIsAlreadyTaken, err)

	} else if driverCheck, err = s.repo.GetDriverByEmail(driver.Email); driverCheck.Email != "" {
		return nil, apperrors.Wrapper(apperrors.ErrEmailIsAlreadyTaken, err)
	}

	return s.repo.CreateDriver(driver)
}

func (s *DriverService) GeDriverByEmail(email string) (*domain.Driver, error) {
	user, err := s.repo.GetDriverByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *DriverService) GetDriverByPhone(phone string) (*domain.Driver, error) {
	driver, err := s.repo.GetDriverByPhone(phone)
	if err != nil {
		return nil, err
	}

	return driver, err
}

func NewDriverService(repo repository.Driver) *DriverService {
	return &DriverService{repo: repo}
}
