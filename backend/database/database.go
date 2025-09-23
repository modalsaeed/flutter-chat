package database

import (
	"fmt"

	"flutter-chat/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func RunMigrations(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	m, err := migrate.New(
		"file://backend/database/migrations",
		dsn,
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func RollbackMigration(cfg *config.Config, steps uint) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	m, err := migrate.New(
		"file://backend/database/migrations",
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Steps(-int(steps)); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
