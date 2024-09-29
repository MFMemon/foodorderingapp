package dbimpl

import (
	"context"
	"fmt"
	dbapi "foodordering-svc/internal/db/sql/common/api"
	pg "foodordering-svc/internal/db/sql/postgres"
)

type DBImpl struct {
	DBClient dbapi.DBClient
}

func NewDBImpl(ctx context.Context, dbDriver string, dbDsn string) (*DBImpl, error) {

	switch dbDriver {
	case "postgres":

		dbClient, err := pg.NewPostgresDb(ctx, dbDsn)

		return &DBImpl{dbClient}, err
	}

	return nil, fmt.Errorf("Given db driver not found")

}
