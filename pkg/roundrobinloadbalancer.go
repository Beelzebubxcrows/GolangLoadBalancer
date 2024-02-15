package pkg

import (
	"fmt"
	"net/http"
)

type RoundRobinLoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []Server
}

func CreateNewRoundRobinLoadBalancer(port string, servers []Server) *RoundRobinLoadBalancer {

	return &RoundRobinLoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers}
}

func (loadbalancer *RoundRobinLoadBalancer) getAvailableServer() Server {

	var serversLeftToTraverse = len(loadbalancer.Servers)
	var currentServer = (loadbalancer.RoundRobinCount + 1) % len(loadbalancer.Servers)

	for serversLeftToTraverse > 0 && !loadbalancer.Servers[currentServer].IsServerAlive() {
		fmt.Println("Skipping ", loadbalancer.Servers[currentServer].GetAddress(), " as it is not available.")

		serversLeftToTraverse--
		currentServer = (currentServer + 1) % len(loadbalancer.Servers)
	}

	fmt.Println("The next server available is ", loadbalancer.Servers[currentServer].GetAddress())

	loadbalancer.RoundRobinCount = currentServer
	loadbalancer.Servers[currentServer].IncrementConnectionsCount()
	return loadbalancer.Servers[currentServer]
}

func (loadbalancer *RoundRobinLoadBalancer) ServeProxyRequest(rw http.ResponseWriter, r *http.Request) {
	targetServer := loadbalancer.getAvailableServer()
	targetServer.ServeRequest(rw, r)
}
