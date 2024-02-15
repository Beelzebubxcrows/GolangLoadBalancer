package pkg

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	GetAddress() string
	IsServerAlive() bool
	ServeRequest(rw http.ResponseWriter, r *http.Request)
	GetActiveConnections() int
	IncrementConnectionsCount()
}

type SimpleServer struct {
	address           string
	proxy             *httputil.ReverseProxy
	alive             bool
	activeconnections int
}

func CreateNewSimpleServer(add string, alive bool) *SimpleServer {

	serverUrl, err := url.Parse(add)
	if err != nil {
		fmt.Print("url parsing failed with error : " + err.Error())
		os.Exit(1)
	}

	return &SimpleServer{
		address: add,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
		alive:   alive,
	}
}

func (simpleServer *SimpleServer) GetAddress() string {
	return simpleServer.address
}

func (simpleserver *SimpleServer) IsServerAlive() bool {
	return simpleserver.alive
}

func (simpleserver *SimpleServer) ServeRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Request being served by ", simpleserver.GetAddress())
	simpleserver.proxy.ServeHTTP(rw, r)
}

func (simpleserver *SimpleServer) GetActiveConnections() int {
	return simpleserver.activeconnections
}

func (simpleServer *SimpleServer) IncrementConnectionsCount() {
	simpleServer.activeconnections++
}
