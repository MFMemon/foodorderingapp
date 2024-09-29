package api

import (
	"context"
	"database/sql"
)

type DBClient interface {
	// Connect(ctx context.Context, driver string, dsn string) (*sql.DB, error)
	Query(ctx context.Context, scanner interface{}, query string, qargs ...any) error
	Exec(ctx context.Context, query string, qargs ...any) (sql.Result, error)
}
