package repository

import (
	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateDriver(driver *domain.Driver) (*domain.Driver, error) {
	query := `INSERT INTO drivers(name, phone, email, password_hash) VALUES ($1, $2, $3, $4);`
	_, err := r.db.Exec(query, driver.Name, driver.Phone, driver.Email, driver.PasswordHash)
	if err != nil {
		return nil, err
	}
	driver, err = r.GetDriverByPhone(driver.Phone)
	driver.PasswordHash = ""
	return driver, err
}

func (r *AuthPostgres) GetDriverByPhone(phone string) (*domain.Driver, error) {
	driver := domain.Driver{}
	query := `SELECT * FROM drivers WHERE phone=$1 `
	err := r.db.Get(&driver, query, phone)
	return &driver, err
}
