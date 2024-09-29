package impl

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	protos "foodordering-svc/internal/gen/protos/usersvc"
	"foodordering-svc/internal/svcconn"
	"foodordering-svc/utils/handlers"

	jwt "github.com/golang-jwt/jwt/v5"
	rand "go.step.sm/crypto/randutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type usersvcconn struct {
// 	svcConns  *svcconn.SvcConn
// }

type usersvcimpl struct {
}

type usersvchandler struct {
	svcConns   *svcconn.SvcConn
	impl       *usersvcimpl
	GrpcServer *grpc.Server
	jwtKey     string
	protos.UnimplementedUserServer
}

func dialToService(addr string) (protos.UserClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return protos.NewUserClient(conn), err
}

func genJwt(key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		})

	ss, _ := token.SignedString(key)
	return ss
}

func NewUserSvcHandler(svcdisaddr string, svcdisProvider string, svcAddr string, dbDriver string, dbConnParams string) (*usersvchandler, error) {

	svcConns, err := svcconn.NewSvcConn(svcdisaddr, svcdisProvider, dbDriver, dbConnParams)
	handlers.HandleErr(err, slog.LevelError)

	err = svcConns.SvcDisRegistry.RegClient.Register(context.Background(), "services.users", svcAddr)
	handlers.HandleErr(err, slog.LevelError)

	svcImpl := newSvcImpl()
	server := grpc.NewServer()
	randKey, _ := rand.Alphanumeric(10)

	svcHandler := usersvchandler{
		svcConns:   svcConns,
		impl:       svcImpl,
		GrpcServer: server,
		jwtKey:     randKey,
	}

	protos.RegisterUserServer(server, &svcHandler)

	return &svcHandler, err

}

func newSvcImpl() *usersvcimpl {
	return &usersvcimpl{}
}

func (s *usersvchandler) RegisterUser(ctx context.Context, usrInfo *protos.UserQueryParams) (*protos.UserAuthInfo, error) {

	type dbRowData struct {
		id int
	}

	qryExec := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := s.svcConns.SqlDb.DBClient.Exec(ctx, qryExec, usrInfo.Name, usrInfo.Email, usrInfo.Password)

	if err != nil {
		return nil, fmt.Errorf("User already exists. Please login to continue.")
	}

	qryResult := dbRowData{}

	qrySelect := "SELECT id FROM USERS WHERE email=$1"
	s.svcConns.SqlDb.DBClient.Query(ctx, &qryResult, qrySelect, usrInfo.Email)

	return &protos.UserAuthInfo{Id: int32(qryResult.id), Token: genJwt(s.jwtKey)}, nil
}

func (s *usersvchandler) LoginUser(ctx context.Context, usrInfo *protos.UserQueryParams) (*protos.UserAuthInfo, error) {

	type dbRowData struct {
		id       int
		password string
	}

	qryResult := dbRowData{}

	qry := "SELECT id,password FROM USERS WHERE email=$1"
	err := s.svcConns.SqlDb.DBClient.Query(ctx, &qryResult, qry, usrInfo.Email)

	if qryResult.password == "" {
		return nil, fmt.Errorf("User does not exist. Please register and then login to continue")
	}

	if err != nil {
		return nil, err
	}

	if qryResult.password != usrInfo.Password {
		return nil, fmt.Errorf("Incorrect email or password.")
	}

	return &protos.UserAuthInfo{Id: int32(qryResult.id), Token: genJwt(s.jwtKey)}, nil
}

func (s *usersvchandler) AuthenticateUser(ctx context.Context, in *protos.UserAuthInfo) (*protos.UserAuthRes, error) {

	token, _ := jwt.ParseWithClaims(
		in.Token,
		jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(s.jwtKey), nil
		})

	if !token.Valid {
		return &protos.UserAuthRes{Status: "Failed"}, nil
	}

	return &protos.UserAuthRes{Status: "Ok"}, nil

}
