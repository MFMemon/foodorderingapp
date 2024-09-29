package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"foodordering-svc/cmd/usersvc/impl"
	"foodordering-svc/utils/dns"
	"foodordering-svc/utils/handlers"
)

func main() {

	svcAddr := fmt.Sprintf("%v:%v", os.Getenv("host"), os.Getenv("port"))

	// svcClusterAddr, err := dns.Resolve(os.Getenv("host"), os.Getenv("port"))
	// handlers.HandleErr(err, "Gateway host DNS resolve error. ", slog.LevelError)

	svcRegProviderAddr, err := dns.Resolve(os.Getenv("svcdis_addr"), os.Getenv("svcdis_port"))
	handlers.HandleErr(err, slog.LevelError)

	svc, err := impl.NewUserSvcHandler(
		svcRegProviderAddr, os.Getenv("svcdis_vendor"), svcAddr,
		os.Getenv("sqldb_driver"), os.Getenv("sqldsn"),
	)

	handlers.HandleErr(err, slog.LevelError)

	tcpListener, err := net.Listen("tcp", svcAddr)
	handlers.HandleErr(err, slog.LevelError)

	err = svc.GrpcServer.Serve(tcpListener)
	handlers.HandleErr(err, slog.LevelError)

}
