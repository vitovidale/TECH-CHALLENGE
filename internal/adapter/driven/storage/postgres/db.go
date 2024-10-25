package postgres

import (
  "context"
  "fmt"

  "github.com/Masterminds/squirrel"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driven/config"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
  *pgxpool.Pool
  QueryBuilder *squirrel.StatementBuilderType
  url         string
}

func NewDb(ctx context.Context, config *config.DB) (*DB, error) {
  url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
    config.Connection,
    config.User,
    config.Password,
    config.Host,
    config.Port,
    config.Name,
  )

  db, err := pgxpool.New(ctx, url)
  if err != nil {
    return nil, err
  }

  err = db.Ping(ctx)
  if err != nil {
    return nil, err
  }

  psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

  return &DB{ db, &psql, url}, nil
}

func (db *DB) GetErrorCode(err error) string {
  pgErr := err.(*pgconn.PgError)
  return pgErr.Code
}

func (db *DB) Close() {
  db.Pool.Close()
}
