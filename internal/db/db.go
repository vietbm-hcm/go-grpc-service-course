package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/vietbm-hcm/go-grpc-service-course/internal/rocket"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_Name")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s  sslmode=%s",
		dbHost,
		dbPort,
		dbUserName,
		dbName,
		dbPassword,
		dbSSLMode,
	)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil

}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
