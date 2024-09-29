package svcconn

import (
	"context"
	sqldb "foodordering-svc/internal/db/sql/common/dbimpl"
	reg "foodordering-svc/internal/svc-discovery/common/registry"
	// "foodordering-svc/utils/handlers"
	// "log/slog"
)

type SvcConn struct {
	SvcDisRegistry *reg.Registry
	SqlDb          *sqldb.DBImpl
}

func NewSvcConn(svcdisaddr string, svcdisProvider string, dbDriver string, dbConnParams string) (*SvcConn, error) {

	ctx := context.Background()

	r, err := reg.NewRegistry(ctx, svcdisaddr, svcdisProvider)

	if err != nil {
		return nil, err
	}

	d, err := sqldb.NewDBImpl(ctx, dbDriver, dbConnParams)

	return &SvcConn{r, d}, err

}

// func (s *SvcConn) RegisterServiceToDiscovery(ctx context.Context, svcId string, svcAddr string) error {
// 	return s.svcdisRegistry.RegClient.Register(ctx, svcId, svcAddr)
// }

// func (s *SvcConn) DiscoverServiceFromDiscovery(ctx context.Context, svcId string) (string, error) {
// 	return s.svcdisRegistry.RegClient.Discover(ctx, svcId)
// }

// func (s *SvcConn) QueryDb(ctx context.Context, query string, qargs ...any) () {
// 	s.sqlDb.DBClient.Query(ctx, query, qargs...)
// }

// func (s *SvcConn) ExecuteIntoDb(ctx context.Context, query string, qargs ...any) {
// 	s.sqlDb.DBClient.Exec(ctx, query, qargs...)
// }
