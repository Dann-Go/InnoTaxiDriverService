package repository

import (
	"errors"

	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Driver interface {
	CreateDriver(driver *domain.Driver) (*domain.DriverResponse, error)
	GetDriverByPhone(phone string) (*domain.Driver, error)
	GetDriverByEmail(email string) (*domain.Driver, error)
}

type DriverRepository struct {
	db *sqlx.DB
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

const createDriverQuery = `	INSERT INTO drivers(name, phone, email, password_hash, taxi_type) 
								VALUES ($1, $2, $3, $4, $5)
							RETURNING
    							id,name, phone, email, rating, taxi_type;`

func (r *DriverRepository) CreateDriver(driver *domain.Driver) (*domain.DriverResponse, error) {

	if driver.TaxiType == "Economy" || driver.TaxiType == "Comfort" || driver.TaxiType == "Business" {

		row := r.db.QueryRow(createDriverQuery, driver.Name, driver.Phone, driver.Email, driver.PasswordHash, driver.TaxiType)
		var driverResponse domain.DriverResponse

		err := row.Scan(
			&driverResponse.ID,
			&driverResponse.Name,
			&driverResponse.Email,
			&driverResponse.Phone,
			&driverResponse.Rating,
			&driverResponse.TaxiType)

		if err != nil {
			return nil, err
		}
		return &driverResponse, err
	} else {
		return nil, errors.New("wrong taxi type")
	}
}

func (r *DriverRepository) GetDriverByPhone(phone string) (*domain.Driver, error) {
	driver := domain.Driver{}
	query := `SELECT * FROM drivers WHERE phone=$1 `
	err := r.db.Get(&driver, query, phone)
	return &driver, err
}

func (r *DriverRepository) GetDriverByEmail(email string) (*domain.Driver, error) {
	driver := domain.Driver{}
	query := `SELECT * FROM drivers WHERE email=$1 `
	err := r.db.Get(&driver, query, email)
	return &driver, err
}
