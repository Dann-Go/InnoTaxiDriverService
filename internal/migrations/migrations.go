package migrations

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func MigrationUp(db *sqlx.DB) error {
	log.Info("Migration Started")

	query, err := ioutil.ReadFile("internal/migrations/001_driversUp.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := db.Exec(string(query)); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("End Migration")
	return nil
}

func MigrationDown(db *sqlx.DB) error {
	log.Info("Migration Started")

	query, err := ioutil.ReadFile("internal/migrations/001_driversDown.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := db.Exec(string(query)); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("End Migration")
	return nil
}
