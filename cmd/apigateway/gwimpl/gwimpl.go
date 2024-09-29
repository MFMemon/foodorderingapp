package gwimpl

import (
	"context"
	"encoding/json"

	// "io"
	"net/http"

	usrsvcprotos "foodordering-svc/internal/gen/protos/usersvc"
	reg "foodordering-svc/internal/svc-discovery/common/registry"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gateway struct {
	svcRegistry *reg.Registry
}

func createHttpResponse(rspCode int, rspConType string, respBody []byte, w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", rspConType)
	(*w).WriteHeader(rspCode)
	(*w).Write(respBody)
}

func dialToService(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}

func NewGateway(regAddr string, regProvider string) (*gateway, error) {

	r, err := reg.NewRegistry(context.Background(), regAddr, regProvider)

	return &gateway{r}, err

}

func (g *gateway) Routes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/user/login/", g.userLoginHandler)
	router.HandleFunc("/user/register/", g.userRegisterHandler)
	router.HandleFunc("/user/review/", g.userReviewHandler)
	router.HandleFunc("/user/createorder/", g.userCreateOrderHandler)
	router.HandleFunc("/user/orderstatus/", g.userOrderStatusHandler)
	router.HandleFunc("/rst/register/", g.rstRegisterHandler)
	router.HandleFunc("/rst/update/", g.rstUpdateHandler)
	router.HandleFunc("/rider/register/", g.riderRegisterHandler)
	router.HandleFunc("/rider/update/", g.riderUpdateHandler)
	router.HandleFunc("/delivery/update/", g.deliveryUpdateHandler)

	return router
}

func (g *gateway) userLoginHandler(w http.ResponseWriter, r *http.Request) {

	svcAddr, err := g.svcRegistry.RegClient.Discover(context.Background(), "services.users")
	if err != nil {
		createHttpResponse(503, "text/plain", []byte(err.Error()), &w)
		return
	}

	conn, err := dialToService(svcAddr)
	if err != nil {
		createHttpResponse(503, "text/plain", []byte(err.Error()), &w)
		return
	}

	useremail := r.URL.Query().Get("email")
	userpwd := r.URL.Query().Get("password")

	authDetails, err := usrsvcprotos.NewUserClient(conn).LoginUser(context.Background(), &usrsvcprotos.UserQueryParams{
		Email:    useremail,
		Password: userpwd,
	})
	if err != nil {
		createHttpResponse(503, "text/plain", []byte(err.Error()), &w)
		return
	}

	encodedAuthDetails, _ := json.Marshal(authDetails)
	createHttpResponse(200, "application/json", encodedAuthDetails, &w)

	// w.WriteHeader(200)
	// io.WriteString(w, "working")

}

func (g *gateway) userRegisterHandler(w http.ResponseWriter, r *http.Request) {

	svcAddr, err := g.svcRegistry.RegClient.Discover(context.Background(), "services.users")
	if err != nil {
		createHttpResponse(503, "text/plain", []byte(err.Error()), &w)
		return
	}

	conn, err := dialToService(svcAddr)
	if err != nil {
		createHttpResponse(503, "text/plain", []byte(err.Error()), &w)
		return
	}

	username := r.URL.Query().Get("name")
	useremail := r.URL.Query().Get("email")
	userpwd := r.URL.Query().Get("password")

	authDetails, err := usrsvcprotos.NewUserClient(conn).RegisterUser(context.Background(), &usrsvcprotos.UserQueryParams{
		Name:     username,
		Email:    useremail,
		Password: userpwd,
	})
	if err != nil {
		createHttpResponse(503, "text/plain", []byte(err.Error()), &w)
		return
	}

	encodedAuthDetails, _ := json.Marshal(authDetails)
	createHttpResponse(200, "application/json", encodedAuthDetails, &w)
}

func (g *gateway) userReviewHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) userCreateOrderHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) userOrderStatusHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) rstRegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) rstUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) riderRegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) riderUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *gateway) deliveryUpdateHandler(w http.ResponseWriter, r *http.Request) {

}
