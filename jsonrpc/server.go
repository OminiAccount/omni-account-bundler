package jsonrpc

import (
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/rpc"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultRpcPath = "/"

	APIEth = "eth"
)

type Server struct {
	cfg RpcsConfig
	srv *rpc.Server
}

type Service struct {
	Name    string
	Service interface{}
}

func NewServer(cfg RpcsConfig, service Service) *Server {
	server := rpc.NewServer()

	if err := server.RegisterName(service.Name, service.Service); err != nil {
		panic(err)
	}

	srv := &Server{
		cfg: cfg,
		srv: server,
	}

	return srv
}

func (s *Server) Start() {
	http.Handle(DefaultRpcPath, enableCORS(s.srv))

	port := strconv.FormatUint(uint64(s.cfg.Port), 10)
	addr := net.JoinHostPort(s.cfg.Host, port)

	errCh := make(chan error, 1)

	go func() {
		errCh <- http.ListenAndServe(addr, nil)
	}()

	<-time.After(10 * time.Millisecond)

	select {
	case err := <-errCh:
		panic(err)
	default:
		log.Infof("Client processor api server listening on: %s", addr)
	}
}
