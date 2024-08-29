package server

import (
	"github.com/OAAC/config"
	"github.com/OAAC/server/endpoints"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"net"
	"net/http"
	"strconv"
	"time"
)

const DefaultRpcPath = "/"

func NewService(apiConfig *config.APIConfig) {
	server := rpc.NewServer()

	myService := new(endpoints.MyBusinessService)
	svcName := "test"
	if err := server.RegisterName(svcName, myService); err != nil {
		panic(err)
	}

	http.Handle(DefaultRpcPath, enableCORS(server))

	port := strconv.FormatUint(uint64(apiConfig.Port), 10)
	addr := net.JoinHostPort(apiConfig.Host, port)

	errCh := make(chan error, 1)

	go func() {
		errCh <- http.ListenAndServe(addr, nil)
	}()

	<-time.After(10 * time.Millisecond)

	select {
	case err := <-errCh:
		panic(err)
	default:
		log.Info("Client processor api server listening on", "addr", addr)
	}
}
