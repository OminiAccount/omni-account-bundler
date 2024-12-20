package pgstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/OAB/utils/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"io"
)

// NewPostgresStorage creates a new Storage DB
func NewPostgresStorage(cfg Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%d",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.MaxConns))
	if err != nil {
		log.Errorf("Unable to parse DB config: %v\n", err)
		return nil, err
	}
	if cfg.EnableLog {
		l, err := NewLogger(cfg)
		if err != nil {
			panic(err)
		}
		config.ConnConfig.Logger = logger{log: l, slowTime: cfg.LogSlowTime}
	}
	config.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
		err := conn.Ping(ctx)
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			log.Errorf("sql Ping EOF, reconnect...")
			return false
		}
		return true
	}
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Errorf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return db, nil
}
