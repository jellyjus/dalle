package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"dalle/config"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn    *pgxpool.Pool
	builder goqu.DialectWrapper
}

func New(cfg config.Postgres) (*Repository, error) {
	connString := strings.Join([]string{
		fmt.Sprintf("user=%s", cfg.User),
		fmt.Sprintf("password=%s", cfg.Password),
		fmt.Sprintf("dbname=%s", cfg.Database),
		fmt.Sprintf("host=%s", cfg.Host),
		fmt.Sprintf("port=%d", cfg.Port),
		fmt.Sprintf("sslmode=%s", "disable"),
	}, " ")
	conf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	conf.MaxConns = 10
	conf.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, fmt.Errorf("pgsql: open connection: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	goqu.SetDefaultPrepared(true)

	return &Repository{
		conn:    pool,
		builder: goqu.Dialect("postgres"),
	}, nil
}
