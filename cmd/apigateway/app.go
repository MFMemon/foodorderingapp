package main

import (
	// "fmt"

	"fmt"
	"log/slog"
	"net/http"
	"os"

	gw "foodordering-svc/cmd/apigateway/gwimpl"
	dns "foodordering-svc/utils/dns"
	"foodordering-svc/utils/handlers"
)

func main() {

	gwAddr := fmt.Sprintf("%v:%v", os.Getenv("host"), os.Getenv("port"))

	svcRegProviderAddr, err := dns.Resolve(os.Getenv("svcdis_addr"), os.Getenv("svcdis_port"))
	handlers.HandleErr(err, slog.LevelError)

	g, err := gw.NewGateway(svcRegProviderAddr, os.Getenv("svcdis_vendor"))
	handlers.HandleErr(err, slog.LevelError)

	s := http.Server{
		Addr:    gwAddr,
		Handler: g.Routes(),
	}

	err = s.ListenAndServe()
	handlers.HandleErr(err, slog.LevelError)

	// fmt.Print(os.Getenv("PATH"))
}
