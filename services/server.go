package services

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/omni-account/client/config"
	"github.com/omni-account/client/services/endpoints"
	"net"
	"net/http"
	"strconv"
	"time"
)

const DefaultRpcPath = "/"

// import (
//
//	"bytes"
//	"encoding/json"
//	log2 "github.com/ethereum/go-ethereum/log"
//	"github.com/omni-account/client/config"
//	"github.com/powerman/rpc-codec/jsonrpc2"
//	"io"
//	"log"
//	"net"
//	"net/http"
//	"net/rpc"
//	"reflect"
//	"strconv"
//	"time"
//
// )
//
// const DefaultRpcPath = "/rpc"
//
//	type rpcServer struct {
//		*rpc.Server
//		funcAlias map[string]string
//	}
//
//	func newRpcServer() *rpcServer {
//		server := rpc.NewServer()
//		return &rpcServer{
//			Server:    server,
//			funcAlias: make(map[string]string),
//		}
//	}
//
//	func (s *rpcServer) RegisterServerName(serverName string, rcvr any) {
//		err := s.Server.RegisterName(serverName, rcvr)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	func (s *rpcServer) RegisterFunctionAlias(serverName string, rcvr any, other ...interface{}) {
//		t := reflect.TypeOf(rcvr)
//		for i := 0; i < t.NumMethod(); i++ {
//			method := t.Method(i)
//			rpcServerName := serverName + "." + method.Name
//			if _, ok := s.funcAlias[method.Name]; ok {
//				return
//			}
//			s.funcAlias[method.Name] = rpcServerName
//		}
//	}
//
//	func (s *rpcServer) ParseFunctionAlias(aliasName string) string {
//		return s.funcAlias[aliasName]
//	}
//
//	type JsonRpcModel struct {
//		JsonRpc string        `json:"jsonrpc"`
//		Method  string        `json:"method"`
//		Params  []interface{} `json:"params"`
//		ID      int           `json:"id"`
//	}
//
//	func NewRpcServer(apiConfig *config.APIConfig) error {
//		s := newRpcServer()
//		//s.RegisterServerName("challenge", new(RpcServer))
//		//s.RegisterFunctionAlias("challenge", new(RpcServer))
//		http.HandleFunc(DefaultRpcPath, func(writer http.ResponseWriter, request *http.Request) {
//			if !corsAndAllowMethod(&writer, request) {
//				return
//			}
//			len := request.ContentLength
//			body := make([]byte, len)
//			request.Body.Read(body)
//			var model JsonRpcModel
//			if err := json.Unmarshal(body, &model); err != nil {
//				log.Fatal(err)
//				return
//			}
//			model.Method = s.ParseFunctionAlias(model.Method)
//			requestBody, err := json.Marshal(model)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			var conn io.ReadWriteCloser = struct {
//				io.Writer
//				io.ReadCloser
//			}{
//				Writer:     writer,
//				ReadCloser: io.NopCloser(bytes.NewReader(requestBody)),
//			}
//			s.ServeRequest(jsonrpc2.NewServerCodec(conn, nil))
//		})
//
//		port := strconv.FormatUint(uint64(apiConfig.Port), 10)
//		addr := net.JoinHostPort(apiConfig.Host, port)
//
//		errCh := make(chan error, 1)
//
//		go func() {
//			errCh <- http.ListenAndServe(addr, nil)
//		}()
//
//		// Capture server startup errors
//		<-time.After(10 * time.Millisecond)
//
//		select {
//		case err := <-errCh:
//			return err
//		default:
//			log2.Info("Spv processor api server listening on", "addr", addr)
//			return nil
//		}
//	}

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
