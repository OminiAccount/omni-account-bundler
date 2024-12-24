package pgstorage

import (
	"github.com/OAB/utils/log"
	"github.com/gobuffalo/packr/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"
	"os"
	"strconv"
)

// RunMigrationsUp migrate up.
func RunMigrationsUp(cfg Config) error {
	return runMigrations(cfg, migrate.Up)
}

// RunMigrationsDown migrate down.
func RunMigrationsDown(cfg Config) error {
	return runMigrations(cfg, migrate.Down)
}

// runMigrations will execute pending migrations if needed to keep
// the database updated with the latest changes in either direction of up or down.
func runMigrations(cfg Config, direction migrate.MigrationDirection) error {
	c, err := pgx.ParseConfig("postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.Name)
	if err != nil {
		return err
	}
	db := stdlib.OpenDB(*c)

	var migrations = &migrate.PackrMigrationSource{Box: packr.New("db-migrations", "./migrations")}
	nMigrations, err := migrate.Exec(db, "postgres", migrations, direction)
	if err != nil {
		return err
	}

	log.Info("successfully ran ", nMigrations, " migrations")
	return nil
}

// InitOrReset will initializes the db running the migrations or
// will reset all the known data and rerun the migrations
func InitOrReset(cfg Config) error {
	// connect to database
	_, err := NewPostgresStorage(cfg)
	if err != nil {
		return err
	}

	// run migrations
	if err := RunMigrationsDown(cfg); err != nil {
		return err
	}
	return RunMigrationsUp(cfg)
}

// NewConfigFromEnv creates config from standard postgres environment variables,
func NewConfigFromEnv() Config {
	maxConns, _ := strconv.Atoi(getEnv("OMNI_BUNDLER_DATABASE_MAXCONNS", "500"))
	return Config{
		User:     getEnv("OMNI_BUNDLER_DATABASE_USER", "test_user"),
		Password: getEnv("OMNI_BUNDLER_DATABASE_PASSWORD", "test_password"),
		Name:     getEnv("OMNI_BUNDLER_DATABASE_NAME", "test_db"),
		Host:     getEnv("OMNI_BUNDLER_DATABASE_HOST", "localhost"),
		Port:     getEnv("OMNI_BUNDLER_DATABASE_PORT", "5435"),
		MaxConns: maxConns,
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
