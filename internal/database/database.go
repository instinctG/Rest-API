package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	DBName   string `mapstructure:"DB_NAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	SSLMode  string `mapstructure:"SSL_MODE"`
}

func NewDataBase(cfg Config) (*pgxpool.Pool, error) {
	log.Info("Setting up new database connection")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode,
	)

	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return db, fmt.Errorf("couldnt connect to database: %w", err)
	}

	conn, err := db.Acquire(context.Background())
	if err != nil {
		log.Errorf("Unable to acquire a database connection: %v\n", err)
	}

	conn.Release()

	log.Info("Connected to database")
	return db, nil
}
