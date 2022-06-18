package migrations

import (
	"io/ioutil"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func MigrationUp(db *sqlx.DB) error {
	log.Info("Migration Started")

	query, err := ioutil.ReadFile("internal/migrations/001_driversUp.sql")
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(query)); err != nil {
		return err
	}

	log.Info("End Migration")
	return nil
}

func MigrationDown(db *sqlx.DB) error {
	log.Info("Migration Started")

	query, err := ioutil.ReadFile("internal/migrations/001_driversDown.sql")
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(query)); err != nil {
		return err
	}

	log.Info("End Migration")
	return nil
}
