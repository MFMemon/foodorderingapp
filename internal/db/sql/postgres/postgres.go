package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
	db *sqlx.DB
}

func NewPostgresDb(ctx context.Context, dsn string) (*PostgresDb, error) {

	time.Sleep(5 * time.Second) // db start-up wait time

	dbConn, err := sqlx.Connect("postgres", dsn)
	return &PostgresDb{dbConn}, err
}

func (pg *PostgresDb) Query(ctx context.Context, scanner interface{}, query string, qargs ...any) error {
	fmt.Print(pg.db, query, qargs)
	rows, err := pg.db.QueryxContext(ctx, query, qargs...)
	rows.StructScan(scanner)
	return err
}

func (pg *PostgresDb) Exec(ctx context.Context, query string, qargs ...any) (sql.Result, error) {
	result, err := pg.db.ExecContext(ctx, query, qargs...)
	return result, err
}
