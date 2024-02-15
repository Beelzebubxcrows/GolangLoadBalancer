package pkg

import (
	"fmt"
	"net/http"
)

type LeastConnectionLoadBalancer struct {
	Port              string
	ActiveConnections int
	Servers           []Server
}

func CreateNewLeastConnectionLoadBalancer(port string, servers []Server) *LeastConnectionLoadBalancer {
	return &LeastConnectionLoadBalancer{
		Port:              port,
		ActiveConnections: 0,
		Servers:           servers}
}

func (loadbalancer *LeastConnectionLoadBalancer) getAvailableServer() Server {

	var leastConnection = loadbalancer.Servers[0].GetActiveConnections()
	var serverSelected = 0
	for serverIndex := 0; serverIndex < len(loadbalancer.Servers); serverIndex++ {

		fmt.Println(loadbalancer.Servers[serverIndex].GetAddress(), " has ", loadbalancer.Servers[serverIndex].GetActiveConnections(), "connections.")

		if leastConnection > loadbalancer.Servers[serverIndex].GetActiveConnections() {
			leastConnection = loadbalancer.Servers[serverIndex].GetActiveConnections()
			serverSelected = serverIndex
		}
	}

	fmt.Println(loadbalancer.Servers[serverSelected].GetAddress(), " has least connections.")

	loadbalancer.Servers[serverSelected].IncrementConnectionsCount()
	return loadbalancer.Servers[serverSelected]
}

func (loadbalancer *LeastConnectionLoadBalancer) ServeProxyRequest(rw http.ResponseWriter, r *http.Request) {
	targetServer := loadbalancer.getAvailableServer()
	targetServer.ServeRequest(rw, r)
}
