package database

import (
	"fmt"
	"log"

	"flutter-chat/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	log.Println("Connecting to database...")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	log.Println("Database connection established.")
	DB = db
	return nil
}

func RunMigrations(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	log.Println("Running database migrations...")
	m, err := migrate.New(
		"file://database/migrations",
		dsn,
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Database migrations applied successfully.")
	return nil
}

func RollbackMigration(cfg *config.Config, steps uint) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	log.Printf("Rolling back %d migration step(s)...", steps)
	m, err := migrate.New(
		"file://database/migrations",
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Steps(-int(steps)); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migration rollback completed.")
	return nil
}
